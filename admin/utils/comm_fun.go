package utils

import (
	"crypto/md5"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"regexp"
	"time"
)

// 生成uuid
func NewUUID() string {
	return uuid.NewV4().String()
}

//MD5生成哈希值
func GetMD5HashCode(messge []byte) string {
	//创建一个使用MD5校验的hash.Hash接口的对象`
	hash := md5.New()
	//输入数据
	hash.Write(messge)
	//计算机出哈希值,返回数据data的MD5校验和
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashcode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashcode
}

// 分页计算一共多少页
func GetPageLimit(num int, page int) int {
	maxpage := 0

	// 除数
	c := num / page
	// 取模
	m := num % page

	if m == 0 {
		maxpage = c
	} else {
		c += 1
		maxpage = c
	}

	return maxpage
}

// 校验邮箱
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 获取当前时间字符串格式 2017-04-11 13:24:04
func GetTimeNowFormatDate() string {
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	return timeNow
}

// 获取随机字符串
func GetRandNumByNumber(n int64) int64 {
	rand.Seed(time.Now().UnixNano()) // 纳秒时间戳
	return rand.Int63n(n)
}
