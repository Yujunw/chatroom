package model

import "errors"

var (
	ERROR_USER_NOT_EXISTS     = errors.New("该用户不存在")
	ERROR_USER_ALREADY_EXISTS = errors.New("该用户已存在")
	ERROR_USER_PWD_WRONG      = errors.New("用户密码错误")
)
