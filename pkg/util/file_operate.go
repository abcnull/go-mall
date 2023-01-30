package util

import "os"

// IsDirExist 判断 dir 是否存在
func IsDirExist(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

// CreateDir 创建 dir 路径
func CreateDir(path string) error {
	return os.Mkdir(path, 755) // todo: 755？
}
