package router

import (
	"NewProject/app/middleware"
	"NewProject/app/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type App struct {
	UserService         *service.UserService
	ResourceService     *service.ResourceService
	RoleService         *service.RoleService
	RoleResourceService *service.RoleResourceService
	UserRoleService     *service.UserRoleService
}

var NewApp = wire.NewSet(wire.Struct(new(App), "*"))

func (a *App) SetGroupRouter(g *gin.Engine) {
	userRouter := g.Group("/user")
	resourceRouter := g.Group("/resource")
	roleRouter := g.Group("/role")
	roleResourceRouter := g.Group("/roleResource")
	userRoleRouter := g.Group("/userRole")
	a.SetRouters(userRouter)
	a.SetRouters(resourceRouter)
	a.SetRouters(roleRouter)
	a.SetRouters(roleResourceRouter)
	a.SetRouters(userRoleRouter)
}

func (a *App) SetRouters(group *gin.RouterGroup) {
	//用户路由
	userGroup := group.Group("/userOperation")
	userGroup.POST("/register", a.UserService.Register)
	userGroup.POST("/login", a.UserService.Login)
	userGroup.Use(middleware.ParseToken()) //使用中间件验证token
	userGroup.GET("/getUserList", a.UserService.GetUserList)
	userGroup.GET("/getUserInfo", a.UserService.GetUserInfo)
	userGroup.PUT("/updateUser", a.UserService.UpdateUser)
	userGroup.DELETE("/delUser", a.UserService.DelUser)
	//资源路由
	resourceGroup := group.Group("/resourceOperation")
	resourceGroup.GET("/getResourceList", a.ResourceService.GetResourceList)
	resourceGroup.GET("/getResource", a.ResourceService.GetResource)
	resourceGroup.POST("/createResource", a.ResourceService.CreateResource)
	resourceGroup.PUT("/updateResource", a.ResourceService.UpdateResource)
	resourceGroup.DELETE("/deleteResource", a.ResourceService.DelResource)
	//角色路由
	roleGroup := group.Group("/roleOperation")
	roleGroup.POST("addRole", a.RoleService.AddRole)
	roleGroup.GET("getRole", a.RoleService.GetRole)
	roleGroup.GET("getRoleList", a.RoleService.GetRoleList)
	roleGroup.PUT("updateRole", a.RoleService.UpdateRole)
	roleGroup.DELETE("delRole", a.RoleService.DelRole)
	//角色--资源绑定
	roleResourceGroup := group.Group("/roleResourceBindOperation")
	roleResourceGroup.POST("addRoleResourceBind", a.RoleResourceService.RoleResourceBind)
	roleResourceGroup.GET("getRoleOwnedResourceList", a.RoleResourceService.GetRoleOwnedResourceList)
	//用户--角色绑定
	userRoleGroup := group.Group("/userRoleOperation")
	userRoleGroup.POST("addUserRole", a.UserRoleService.AddUserRole)
	userRoleGroup.GET("userOwnedRole", a.UserRoleService.UserOwnedRole)
	userRoleGroup.GET("userOwnedResource", a.UserRoleService.UserOwnedResource)
}

func InitGenEngine(app *App) *gin.Engine {

	engin := gin.Default()
	// gin.Default()默认包含 Logger 和 Recovery 中间件
	//Logger 中间件用于记录每个请求的基本信息，包括请求路径、请求方法、请求状态码、响应时间等。这对于监控应用和调试问题非常有用。
	//Recovery 中间件主要作用是从任何恐慌（panics）中恢复，并在发生恐慌时写入 HTTP 状态码 500。
	//在 Go 语言里，当使用 panic() 时，程序会崩溃退出。
	//而 gin.Recovery 中间件会捕获这种异常，避免程序因未处理的异常而崩溃退出，保证服务的持续运行。
	//不过，对于链接断开的情况，不会有 HTTP 状态码返回

	app.SetGroupRouter(engin)

	return engin
}
