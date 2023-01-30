package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	RedisDb     string
	RedisAddr   string
	RedisPwd    string
	RedisDbName string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	Host        string
	ProductPath string
	AvatarPath  string
)

var (
	ReadDSN  string
	WriteDSN string
)

// InitConf 从本地读取环境变量
func InitConf(path string) {
	// 读取环境配置，path 跟执行者路径有关
	file, err := ini.Load(path)
	if err != nil {
		panic(err)
	}
	// 加载配置
	LoadIniData(file)
	// todo: mysql 读写分离，注意主从同步配置。读读操作多，所以作为主，写多操作相对于读少，所以作为从
	// mysql read(主)
	ReadDSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		DbUser,
		DbPassword,
		DbHost,
		DbPort,
		DbName,
	)
	// mysql write(从)(主从复制)
	WriteDSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		DbUser,
		DbPassword,
		DbHost,
		DbPort,
		DbName,
	)
}

func LoadIniData(file *ini.File) {
	// service
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
	// mysql
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
	// redis
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPwd = file.Section("redis").Key("RedisPwd").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
	// email
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtpPass = file.Section("email").Key("SmtpPass").String()
	// path
	Host = file.Section("path").Key("Host").String()
	ProductPath = file.Section("path").Key("ProductPath").String()
	AvatarPath = file.Section("path").Key("AvatarPath").String()
}
