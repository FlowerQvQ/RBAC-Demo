package service

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewUserService,
	NewResourceService,
	NewRoleService,
	NewRoleResourceService,
	NewUserRoleService,
)
