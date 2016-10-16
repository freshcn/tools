package vaildate

import (
	"regexp"
)

// IsEmail 判断是否为正确的邮箱
func IsEmail(email string) bool {
	reg := regexp.MustCompile(`^[\w-\.]*@[\w-][\.\w-]*\.[\w]{2,4}$`)
	return reg.MatchString(email)
}
