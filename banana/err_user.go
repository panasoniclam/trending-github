package banana

import (
	"errors"
)
var (
	UserConfList  = errors.New("Người dùng đã tồn tại")
	UserConfFound = errors.New("Người dùng không tồn tại")
	SignUpFailed = errors.New("Đăng kí thất bại")
)