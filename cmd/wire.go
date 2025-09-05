//go:build wireinject
// +build wireinject

package main

import (
	"NewProject/app/biz"
	"NewProject/app/data"
	"NewProject/app/router"
	"NewProject/app/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitApp() *gin.Engine {
	panic(wire.Build(
		router.ProviderSet,
		router.InitGenEngine,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
	))
}
