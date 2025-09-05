package biz

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewUserBiz,
	NewResourceBiz,
	NewRoleBiz,
	NewRoleResourceBiz,
	NewUserRoleBiz,
)
