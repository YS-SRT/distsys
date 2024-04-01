package utils

import (
	"distsys/global"
	"encoding/base64"
	"os"

	"github.com/duke-git/lancet/v2/cryptor"
	"go.uber.org/zap"
)

func PathExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {

	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.GLogger.Debug("create directory" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				global.GLogger.Error("create directory"+v, zap.Any(" error:", err))
				return err
			}
		}
	}
	return err
}

func AESEncode4Conf(value string) string {

	return AESEncode(value, global.GCONFIG.Postgres.EncryCode)
}

func AESEncode(value string, secretKey string) string {

	if len(value) > 0 {

		return base64.StdEncoding.EncodeToString(cryptor.AesEcbEncrypt([]byte(value), []byte(secretKey)))
	}

	return value
}

func AESDecode(value string, secretKey string) string {

	if len(value) > 0 {

		if b, err := base64.StdEncoding.DecodeString(value); err == nil {

			return string(cryptor.AesEcbDecrypt(b, []byte(secretKey)))
		}

	}

	return value
}
