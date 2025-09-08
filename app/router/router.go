package router

import (
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
	userGroup.GET("/getUserList", a.UserService.GetUserList)
	userGroup.GET("/getUserInfo", a.UserService.GetUserInfo)
	userGroup.PUT("/updateUser", a.UserService.UpdateUser)
	//资源路由
	resourceGroup := group.Group("/resourceOperation")
	resourceGroup.GET("/getResourceList", a.ResourceService.GetResourceList)
	resourceGroup.POST("/createResource", a.ResourceService.CreateResource)
	//角色路由
	roleGroup := group.Group("/roleOperation")
	roleGroup.POST("addRole", a.RoleService.AddRole)
	//角色--资源绑定
	roleResourceGroup := group.Group("/roleResourceBindOperation")
	roleResourceGroup.POST("addRoleResourceBind", a.RoleResourceService.RoleResourceBind)
	//用户--角色绑定
	userRoleGroup := group.Group("/userRoleOperation")
	userRoleGroup.POST("addUserRole", a.UserRoleService.AddUserRole)
	userRoleGroup.GET("userOwnedResource", a.UserRoleService.UserOwnedResource)
}

func InitGenEngine(app *App) *gin.Engine {

	engin := gin.Default()

	app.SetGroupRouter(engin)

	return engin
}
