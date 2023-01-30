package util

var Encrypt *Encryption

// Encryption AES 对称加密 // todo
type Encryption struct {
	Key string
}

// AesEncoding 加密
func (k *Encryption) AesEncoding(src string) string {
	return ""
}

// AesDecoding 解密
func (k *Encryption) AesDecoding(pwd string) string {
	return ""
}

// PadPwd 填充密码长度
func PadPwd(srcByte []byte, blockSize int) []byte {
	return nil
}

// UnPadPwd 去掉填充部分
func UnPadPwd(dst []byte) ([]byte, error) {
	return nil, nil
}
