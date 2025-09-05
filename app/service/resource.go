package service

import (
	"NewProject/app/biz"
	"NewProject/app/scheme"
	"NewProject/models"
	"NewProject/pkg/wapper"
	"github.com/gin-gonic/gin"
)

type ResourceService struct {
	ResourceBiz *biz.ResourceBiz
}

func NewResourceService(resourceBiz *biz.ResourceBiz) *ResourceService {
	return &ResourceService{
		ResourceBiz: resourceBiz,
	}
}

// 获取资源列表
func (s *ResourceService) GetResourceList(c *gin.Context) {
	var (
		resourceListReq scheme.ResourceListReq
		err             error
		resourceList    scheme.ResourceListResp
		errCode         wapper.ErrorCode
	)
	err = c.ShouldBindJSON(&resourceListReq)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}
	resourceList, errCode = s.ResourceBiz.GetResourceList(resourceListReq)
	if errCode != wapper.Success {
		wapper.ResError(c, wapper.GetResourceListFailed)
		return
	}
	wapper.ResSuccess(c, resourceList)
}

//查询资源

// 增加资源
func (s *ResourceService) CreateResource(c *gin.Context) {
	var (
		CreateResource     scheme.ResourceCreateReq
		err                error
		errCode            wapper.ErrorCode
		createResourceInfo models.Resource
	)
	err = c.ShouldBindJSON(&CreateResource)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}
	createResourceInfo, errCode = s.ResourceBiz.CreateResource(CreateResource)
	if errCode != wapper.Success {
		wapper.ResError(c, wapper.AddResourceFailed)
		return
	}
	wapper.ResSuccess(c, createResourceInfo)

}

//更新资源

//删除资源
