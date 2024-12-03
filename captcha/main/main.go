package main

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

const (
	// Default number of digits in captcha solution.
	DefaultLen = 6
	// The number of captchas created that triggers garbage collection used
	// by default store.
	CollectNum = 100
	// Expiration time of captchas used by default store.
	Expiration = 10 * time.Minute
)

const (
	// Standard width and height of a captcha image.
	StdWidth  = 240
	StdHeight = 80
)

var (
	ErrNotFound = errors.New("captcha: id not found")
)

func main() {
	email := ValidateEmail("1808479176@qq.com")
	fmt.Println(email)
}

var regEmail = regexp.MustCompile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$")
var regTel = regexp.MustCompile("^1[345789]{1}\\d{9}$") // 假设手机号为11位数字

// ValidateEmail 校验邮箱格式
func ValidateEmail(email string) bool {

	return regEmail.MatchString(email)
}

// ValidatePhone 校验电话号码格式
func ValidatePhone(phone string) bool {

	return regTel.MatchString(phone)
}
