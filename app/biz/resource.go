package biz

import (
	"NewProject/app/data"
	"NewProject/app/scheme"
	"NewProject/models"
	"NewProject/pkg/wapper"
)

type ResourceBiz struct {
	ResourceData *data.ResourceData
}

func NewResourceBiz(resourceData *data.ResourceData) *ResourceBiz {
	return &ResourceBiz{
		ResourceData: resourceData,
	}
}

// 获取资源列表
func (b *ResourceBiz) GetResourceList(resourceListReq scheme.ResourceListReq) (scheme.ResourceListResp, wapper.ErrorCode) {
	resourceList, err := b.ResourceData.GetResourceList(resourceListReq)
	if err != nil {
		return scheme.ResourceListResp{}, wapper.GetResourceListFailed
	}
	return resourceList, wapper.Success
}

// 增加资源
func (b *ResourceBiz) CreateResource(createResourceReq scheme.ResourceCreateReq) (models.Resource, wapper.ErrorCode) {
	createInfo, err := b.ResourceData.CreateResource(createResourceReq)
	if err != nil {
		return models.Resource{}, wapper.AddResourceFailed
	}
	return createInfo, wapper.Success
}

//更新资源

//删除资源
