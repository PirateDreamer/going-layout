package xerrs

import "going-demo/pkg/xerr"

var (
	ErrUserNotFound = xerr.NewBizErr("001001001", "用户不存在")
)
