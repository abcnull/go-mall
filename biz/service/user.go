package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-mall/biz/dal/entity"
	"go-mall/biz/model"
	"go-mall/client/mysql"
	"go-mall/conf"
	"go-mall/pkg/status"
	"go-mall/pkg/util"
	"gopkg.in/mail.v2"
	"gorm.io/gorm"
	"io/ioutil"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

// LoginCheck 登录时候后端纯文本校验
func LoginCheck(c *gin.Context, loginReq *model.LoginRequest) (*status.Status, error) {
	// 校验用户名

	// 校验密码 len 属于 [6, 15]
	if len(loginReq.PassWord) < 6 || len(loginReq.PassWord) > 15 {
		return status.InvalidParam, errors.New("密码不符合规范")
	}
	return status.Success, nil
}

// LoginService 登录操作
func LoginService(c *gin.Context, loginReq *model.LoginRequest) (*model.LoginResponse, *status.Status, error) {
	mysqlCli := mysql.Client(c)

	// 看看 user 是否存在
	count := new(int64)
	user := new(entity.User)
	err := mysqlCli.Model(&entity.User{}).Where("user_name = ?", loginReq.UserName).Find(user).Count(count).Error
	if err != nil && err != gorm.ErrRecordNotFound { // 数据库有问题
		return nil, status.Error, err
	} else if *count == 0 { // 搜不到记录
		return nil, status.InvalidParam, errors.New("账号未注册")
	}

	// 如果存在此用户，判定密码校验没有问题
	if flag, _ := user.CheckPwd(loginReq.PassWord); !flag {
		return nil, status.InvalidParam, errors.New("密码错误")
	}

	// http 是无状态协议，所以为了识别请求者，后端需要签发 token 给请求者
	token, err := util.GenerateToken(user.ID, user.UserName, 0)
	if err != nil {
		return nil, status.AccessErr, err // token 签发失败
	}

	loginResp := &model.LoginResponse{
		Token: token,
		UserInfo: model.BasicUserInfo{
			UserName: user.UserName,
			NickName: user.NickName,
			Email:    user.Email,
			Avatar:   conf.Host + conf.HttpPort + conf.AvatarPath + user.Avatar,
			CreateAt: user.CreatedAt.Unix(),
			UpdateAt: user.UpdatedAt.Unix(),
		},
	}
	return loginResp, status.Success, nil
}

// RegisterCheck 注册时候后端纯文本校验
func RegisterCheck(c *gin.Context, registerReq model.RegisterRequest) (*status.Status, error) {
	// 校验昵称 len 属于 [3, 10]
	if len(registerReq.NickName) < 3 || len(registerReq.NickName) > 10 {
		return status.InvalidParam, errors.New("昵称不符合规范")
	}

	// 校验用户名

	// 校验密码 len 属于 [6, 15]
	if len(registerReq.PassWord) < 6 || len(registerReq.PassWord) > 15 {
		return status.InvalidParam, errors.New("密码不符合规范")
	}

	// 校验 req 中密钥是否正确
	if registerReq.Key == "" || len(registerReq.Key) != 16 {
		return status.InvalidParam, errors.New("密钥问题")
	}

	// 对称加密
	aes := new(util.Encryption)
	aes.Key = registerReq.Key
	// todo:

	return status.Success, nil
}

// RegisterService 注册操作
func RegisterService(c *gin.Context, registerReq model.RegisterRequest) (*status.Status, error) {
	mysqlCli := mysql.Client(c)

	// 查看 user_name 是否已经存在
	var count int64
	var user = new(entity.User)
	err := mysqlCli.Model(&entity.User{}).Where("user_name = ?", registerReq.UserName).Find(user).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound { // 如果数据查询出错
		return status.Error, err
	} else if count != 0 { // 如果用户名已存在
		return status.InvalidParam, errors.New("用户名已存在")
	}

	// user_name 不存在
	user = &entity.User{
		UserName: registerReq.UserName,
		NickName: registerReq.NickName,
		Status:   entity.DefaultUserStatus,
		Avatar:   "avatar.JPG",
		Money:    "0", // util.Encrypt.AesEncoding("0")
	}

	// 先给前端传过来的 pwd 做一个解密操作 todo

	// 给密码加密；其实一般来说前端传过来的也是密文，服务端需要先进行解密操作
	if err := user.SetPwd(registerReq.PassWord); err != nil {
		return status.Error, errors.New("加密出错")
	}

	// 数据库创建用户
	if err := mysqlCli.Model(&entity.User{}).Create(user).Error; err != nil {
		return status.Error, errors.New("数据库插入错误")
	}

	return status.Success, nil
}

// UpdateUserCheck 更新后纯文本校验
func UpdateUserCheck(c *gin.Context, updateReq *model.UpdateUserRequest) (*status.Status, error) {
	// 校验昵称 len 属于 [3, 10]
	if len(updateReq.NickName) < 3 || len(updateReq.NickName) > 10 {
		return status.InvalidParam, errors.New("昵称不符合规范")
	}
	return status.Success, nil
}

// UpdateUserService 更新操作
func UpdateUserService(c *gin.Context, updateReq *model.UpdateUserRequest, userId uint) (*status.Status, error) {
	db := mysql.Client(c)
	user := new(entity.User)

	// 依据 userId 获取用户
	if err := db.Model(&entity.User{}).Where("id = ?", userId).First(user).Error; err != nil {
		return status.Error, err
	}

	// 修改 user todo: 目前只提供变更 nick_name
	user.NickName = updateReq.NickName

	// 更新 user
	if err := db.Model(&entity.User{}).Where("id = ?", userId).Updates(user).Error; err != nil {
		return status.Error, err
	}

	return status.Success, nil
}

// UploadAvatarCheck 后端校验一下 Avatar 上传请求入参是否正常 todo
func UploadAvatarCheck(c *gin.Context, uploadAvatarReq *model.UploadAvatarRequest) (*status.Status, error) {
	return status.Success, nil
}

// UploadAvatarService 更新用户头像
func UploadAvatarService(c *gin.Context, userId uint, file multipart.File, size int64) (*model.UploadAvatarResponse, *status.Status, error) {
	mysqlCli := mysql.Client(c)

	// 通过 userId 找到 user
	user := new(entity.User)
	if err := mysqlCli.Model(&entity.User{}).Where("id = ?", userId).Find(user).Error; err != nil {
		return nil, status.Error, err
	}

	// 在 static/img/avatar 下创建指定用户文件夹，不存在则创建
	userPath := "." + conf.AvatarPath + "user" + strconv.Itoa(int(userId)) + "/"
	if !util.IsDirExist(userPath) { // 判定 dir 是否存在，如果不存在则创建
		if err := util.CreateDir(userPath); err != nil {
			return nil, status.Error, err
		}
	}

	// 保存图片到本地 static 中 todo: 为什么谷歌浏览器中能搜到这个文件，但是项目这个 jpg 文件没有出现
	avatarPath := userPath + user.UserName + ".jpg" // todo: 有没有什么方式不是默认 jpg
	img, err := ioutil.ReadAll(file)                // 文件读成 byte slice
	if err != nil {
		return nil, status.Error, err
	}
	if err := ioutil.WriteFile(avatarPath, img, 0666); err != nil { // 把文件流写入到某个路径 todo: 0666
		return nil, status.Error, err
	}

	// 更新用户 avatar todo: 需要把上传到文件名称更新到数据库 avatar 中
	// mysqlCli.Model(&entity.User{}).Where("id = ?", userId).Update("avatar", )

	uploadAvatarResp := &model.UploadAvatarResponse{
		UserInfo: model.BasicUserInfo{
			UserName: user.UserName,
			NickName: user.NickName,
			Email:    user.Email,
			Avatar:   conf.Host + conf.HttpPort + conf.AvatarPath + "user" + strconv.Itoa(int(userId)) + "/" + user.UserName + ".jpg",
			CreateAt: user.CreatedAt.Unix(),
			UpdateAt: user.UpdatedAt.Unix(),
		},
	}
	return uploadAvatarResp, status.Success, nil
}

// SendEmailCheck 发送邮件请求 check
func SendEmailCheck(c *gin.Context, sendEmailReq *model.SendEmailRequest) (*status.Status, error) {
	return status.Success, nil
}

// SendEmailService 发送邮件 service
func SendEmailService(c *gin.Context, sendEmailReq *model.SendEmailRequest, userId uint) (*status.Status, error) {
	mysqlCli := mysql.Client(c)

	// 获取 emailToken
	emailToken, err := util.GenerateEmailToken(userId, sendEmailReq.Email, sendEmailReq.Password, sendEmailReq.OperationType)
	if emailToken == "" || err != nil {
		return status.Error, err
	}

	// 模版
	notice := new(entity.Notice)
	if err := mysqlCli.Model(&entity.Notice{}).Where("id = ?", sendEmailReq.OperationType).Find(notice).Error; err != nil {
		return status.Error, err
	}
	address := conf.ValidEmail + emailToken // todo: 为什么是这个地址？
	noticeText := strings.Replace(notice.Text, "Email", address, -1)

	// 发送邮件
	SendEmail(c, sendEmailReq.Email, noticeText)

	return status.Success, nil
}

// SendEmail 发送邮件
func SendEmail(c *gin.Context, toEmail string, emailText string) {
	// 发送邮件
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", toEmail)
	m.SetBody("text/html", emailText)
	dialer := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass) // todo: smtppass现在没填写
	dialer.StartTLSPolicy = mail.MandatoryStartTLS
	// todo: 没写完
}

// ValidateEmailService 验证邮箱
func ValidateEmailService(c *gin.Context, emailToken string) (*status.Status, error) {
	emailClaims, err := util.ParseEmailToken(emailToken) // email 权限不过
	if err != nil || emailClaims.ExpiresAt < time.Now().Unix() {
		return status.AccessErr, err
	}

	// 通过 userId 获取 user
	mysqlCli := mysql.Client(c)
	user := new(entity.User)
	if err := mysqlCli.Model(&entity.User{}).Where("id = ?", emailClaims.UserId).Find(user).Error; err != nil {
		return status.Error, err
	}

	// 看看 emailClaims 中的 operationType 的类型
	if emailClaims.OperationType == 1 {
		// 绑定邮箱
		user.Email = emailClaims.Email
	} else if emailClaims.OperationType == 2 {
		// 解绑邮箱
		user.Email = ""
	} else if emailClaims.OperationType == 3 {
		// 重新设置密码
		if err := user.SetPwd(emailClaims.Password); err != nil {
			return status.Error, err
		}
	}

	// 更新 userId 指定的 user
	if err := mysqlCli.Model(&entity.User{}).Where("id = ?", emailClaims.UserId).Updates(user).Error; err != nil {
		return status.Error, err
	}

	return status.Success, nil
}

func ShowMoneyService(c *gin.Context, showMoneyReq *model.ShowMoneyRequest, userId uint) (*status.Status, error, *model.ShowMoneyResponse) {
	mysqlCli := mysql.Client(c)

	// 依据 user_id 找到 user
	user := new(entity.User)
	if err := mysqlCli.Model(&entity.User{}).Where("id = ?", userId).Find(user).Error; err != nil {
		return status.Error, err, nil
	}

	util.Encrypt.Key = showMoneyReq.Key
	return status.Success, nil, &model.ShowMoneyResponse{
		UserId:   user.ID,
		UserName: user.UserName,
		Money:    util.Encrypt.AesDecoding(user.Money), // todo: 用户的钱是加密形式存放数据库的，所以展示需要解密
	}
}
