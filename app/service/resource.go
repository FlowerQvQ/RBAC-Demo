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

// 查询资源
func (s *ResourceService) GetResource(c *gin.Context) {
	var (
		getResourceReq scheme.ResourceGetReq
		err            error
	)
	err = c.ShouldBindJSON(&getResourceReq)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}
	resourceInfo, errCode := s.ResourceBiz.GetResource(getResourceReq)
	if errCode != wapper.Success {
		wapper.ResError(c, wapper.GetResourceFailed)
		return
	}
	wapper.ResSuccess(c, resourceInfo)
}

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

// 更新资源
func (s *ResourceService) UpdateResource(c *gin.Context) {
	var (
		updateResourceReq scheme.ResourceUpdateReq
		err               error
	)
	err = c.ShouldBindJSON(&updateResourceReq)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}
	updateResourceData, errCode := s.ResourceBiz.UpdateResource(updateResourceReq)
	if errCode != wapper.Success {
		wapper.ResError(c, wapper.UpdateResourceFailed)
		return
	}
	wapper.ResSuccess(c, updateResourceData)
}

// 删除资源
func (s *ResourceService) DelResource(c *gin.Context) {
	var (
		resourceId scheme.ResourceDelReq
		err        error
	)
	err = c.ShouldBindJSON(&resourceId)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}
	errCode := s.ResourceBiz.DelResource(resourceId)
	if errCode != wapper.Success {
		wapper.ResError(c, wapper.DelResourceFailed)
		return
	}
	wapper.ResSuccess(c, "删除成功")
}
