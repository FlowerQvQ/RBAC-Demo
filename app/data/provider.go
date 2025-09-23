package data

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewRedisData,
	NewUserData,
	NewResourceData,
	NewRoleData,
	NewRoleResourceData,
	NewRoleUserData,
)
