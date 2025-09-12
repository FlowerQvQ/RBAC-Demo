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

// 查询资源
func (b *ResourceBiz) GetResource(resourceGetReq scheme.ResourceGetReq) (models.Resource, wapper.ErrorCode) {
	resourceInfo, err := b.ResourceData.GetResource(resourceGetReq)
	if err != nil {
		return models.Resource{}, wapper.DataNotFound
	}
	return resourceInfo, wapper.Success
}

// 增加资源
func (b *ResourceBiz) CreateResource(addResourceInfo models.Resource) (models.Resource, wapper.ErrorCode) {
	createInfo, err := b.ResourceData.CreateResource(addResourceInfo)
	if err != nil {
		return models.Resource{}, wapper.AddResourceFailed
	}
	return createInfo, wapper.Success
}

// 更新资源
func (b *ResourceBiz) UpdateResource(updateResourceInfo models.Resource) (models.Resource, wapper.ErrorCode) {
	updateData, err := b.ResourceData.UpdateResource(updateResourceInfo)
	if err != wapper.Success {
		return models.Resource{}, wapper.UpdateResourceFailed
	}
	return updateData, wapper.Success
}

// 删除资源
func (b *ResourceBiz) DelResource(resourceDelReq scheme.ResourceDelReq) wapper.ErrorCode {
	err := b.ResourceData.DelResource(resourceDelReq)
	if err != wapper.Success {
		return wapper.DelResourceFailed
	}
	return wapper.Success

}
