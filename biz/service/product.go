package service

import (
	"github.com/gin-gonic/gin"
	"go-mall/biz/dal/entity"
	"go-mall/biz/model"
	"go-mall/client/mysql"
	"go-mall/client/redis"
	"go-mall/pkg/status"
	"gorm.io/gorm"
	"mime/multipart"
	"strconv"
	"sync"
)

func CreateProductServiceCheck(c *gin.Context, createProductReq *model.CreateProductRequest) {
	// todo: check 一下创建商品入参有无问题
}

func CreateProductService(c *gin.Context, userId uint, createProductReq *model.CreateProductRequest, files []*multipart.FileHeader) (*model.CreateProductResponse, *status.Status, error) {
	mysqlCli := mysql.Client(c)

	// 获取用户数据
	user := new(entity.User)
	if err := mysqlCli.Model(&entity.User{}).Where("id = ?", userId).Find(user).Error; err != nil {
		return nil, status.Error, err
	}

	// todo: 这里要加锁，从查询有多少行商品，到新增一个商品数据，这块内容需要加锁包裹起来

	// 先查一下 product id 应该排到多少
	lastProduct := new(entity.Product)
	if err := mysqlCli.Model(&entity.Product{}).Last(lastProduct).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, status.Error, err
	}
	nextProductId := lastProduct.ID + 1

	// 以第一张图作为封面图
	coverImg, err := files[0].Open() // todo: 这个操作需要学习
	if err != nil {
		return nil, status.Error, err
	}
	// 获取封面图路径
	coverImgPath := GetUploadProductImgPath(userId, nextProductId, createProductReq.Name)
	// 存储封面图
	if sta, err := UploadProductImgToLocal(coverImg, coverImgPath); err != nil {
		return nil, sta, err
	}

	// 生成商品
	product := &entity.Product{
		Num:           1, // todo: 这里可以后续要调整，因为闲置商品一般都是 1
		Name:          createProductReq.Name,
		CategoryId:    createProductReq.CategoryId,
		Title:         createProductReq.Title,
		Info:          createProductReq.Info,
		ImgPath:       coverImgPath, // 封面图
		Price:         createProductReq.OriginPrice,
		DiscountPrice: createProductReq.DiscountPrice,
		Locate:        createProductReq.Locate,
		Freight:       uint(createProductReq.Freight),
		OnSale:        true,
		BossID:        user.ID,
		BossName:      user.NickName,
		BossAvatar:    user.Avatar,
	}

	// 存储到商品表中
	if err := mysqlCli.Model(&entity.Product{}).Create(product).Error; err != nil {
		return nil, status.Error, err
	}

	// todo: 这里锁结束

	// 存储商品图片到 static，图片路径存储到 product_img 表中，用 WaitGroup todo: 这么做可以吗？
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	var imgSta *status.Status
	var imgErr error
	for i, f := range files {
		go func() { // todo: 多协程怎么返回状态和 err
			if i != 0 {
				imgPath := GetUploadProductImgPath(userId, nextProductId, createProductReq.Name+strconv.Itoa(i)) // 可以随机生成一个名字
				imgFile, _ := f.Open()
				imgSta, imgErr = UploadProductImgToLocal(imgFile, imgPath)

				productImg := &entity.ProductImg{
					ProductID: nextProductId,
					ImgPath:   imgPath,
				}
				imgErr = mysqlCli.Model(&entity.ProductImg{}).Create(productImg).Error // 存储图片
				wg.Done()
			}
		}()
	}
	wg.Wait()
	if imgErr != nil {
		return nil, imgSta, imgErr
	}

	// todo: 输出 resp
	return &model.CreateProductResponse{
		Id:                 product.ID,
		Name:               product.Name,
		CategoryId:         product.CategoryId,
		Tile:               product.Title,
		Info:               product.Info,
		ImgPath:            product.ImgPath, // 封面图
		Locate:             product.Locate,
		OriginPrice:        product.Price,
		DiscountPrice:      product.DiscountPrice,
		Freight:            model.FreightType(product.Freight),
		OnSale:             product.OnSale,
		BossId:             product.BossID,
		BossName:           product.BossName,
		BossAvatar:         product.BossAvatar,
		ExposureTimes:      0,
		ClickTimes:         0,
		CommunicationTimes: 0,
		CreateAt:           product.CreatedAt.Unix(),
		UpdateAt:           product.UpdatedAt.Unix(),
	}, status.Success, nil
}

// GainExposureTimes 获取商品曝光量从 redis 中拿
func GainExposureTimes(productId uint) uint64 {
	exposureTimesStr, _ := redis.RedisCli.Get("exposure_product_" + strconv.Itoa(int(productId))).Result()
	exposureTimes, _ := strconv.ParseUint(exposureTimesStr, 10, 64)
	return exposureTimes
}

// AddExposureTimes 增加商品曝光量，增加 redis 中
func AddExposureTimes(productId uint) uint64 {
	redis.RedisCli.Incr("exposure_product_" + strconv.Itoa(int(productId))) // 将存储的数字值增加 1
	return 0
}

// GainClickTimes 获取商品点击数，从 redis 中拿
func GainClickTimes(productId uint) int64 {
	return 0
}

// AddClickTimes 增加商品点击数，增加 redis 中
func AddClickTimes(productId uint) int64 {
	return 0
}

// GainCommunicationTimes 获取商品沟通数，从 redis 中拿
func GainCommunicationTimes() int64 {
	return 0
}

// AddCommunicationTimes 增加商品沟通数，增加 redis 中
func AddCommunicationTimes() int64 {
	return 0
}
