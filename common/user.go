package common

import (
	"crypto/sha256"
	"errors"
)

func GetPwd(password, salt string) string {
	encrypt := sha256.Sum256([]byte(password))
	secEncrypt := sha256.Sum256([]byte(string(encrypt[:]) + salt))
	return string(secEncrypt[:])
}

func CheckPwd(password, enPassword, salt string) error {
	uPassword := GetPwd(password, salt)
	if string(uPassword[:]) != enPassword {
		return errors.New("用户密码错误")
	}
	return nil
}
