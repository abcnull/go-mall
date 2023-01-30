package service

import (
	"go-mall/conf"
	"go-mall/pkg/status"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
)

// GetUploadProductImgPath 图片存放路径
func GetUploadProductImgPath(userId uint, productId uint, productName string) string {
	// 获取路径
	bossPrePath := "." + conf.ProductPath + "boss" + strconv.Itoa(int(userId))
	productPrePath := bossPrePath + "/product" + strconv.Itoa(int(productId))

	// 文件夹不存在就需要创建
	if _, err := os.Stat(productPrePath); os.IsNotExist(err) {
		os.MkdirAll(productPrePath, os.ModePerm)
	}

	return productPrePath + "/" + productName + ".jpg"
}

// UploadProductImgToLocal 保存图片资源存储本地
func UploadProductImgToLocal(file multipart.File, imgPath string) (*status.Status, error) {
	// 读取文件内容
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return status.Error, err
	}

	// 将读取的文件内容写入流
	if err := ioutil.WriteFile(imgPath, content, 0666); err != nil {
		return status.Error, err
	}

	return status.Success, nil
}
