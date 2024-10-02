package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// BcryptMake 生成密码哈希
func BcryptMake(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// BcryptMakeCheck 检查密码哈希
func BcryptMakeCheck(pwd []byte, hashedPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, pwd)
	return err == nil
}
