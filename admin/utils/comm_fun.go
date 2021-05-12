package utils

import (
	"crypto/md5"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
)

// 生成uuid
func NewUUID() string {
	return uuid.NewV4().String()
}


//MD5生成哈希值
func GetMD5HashCode(messge []byte)string{
	//创建一个使用MD5校验的hash.Hash接口的对象`
	hash := md5.New()
	//输入数据
	hash.Write(messge)
	//计算机出哈希值,返回数据data的MD5校验和
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashcode:=hex.EncodeToString(bytes)
	//返回哈希值
	return hashcode
}