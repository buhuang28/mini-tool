package etype

import "errors"

var (
	NAME_EXIST = errors.New("该进程名称已存在，请勿重复添加")
)
