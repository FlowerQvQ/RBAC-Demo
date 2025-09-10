package data

import (
	"NewProject/app/scheme"
	"NewProject/models"
	"NewProject/pkg/wapper"
)

type ResourceData struct {
	DB *Data
}

func NewResourceData(data *Data) *ResourceData {
	return &ResourceData{
		DB: data,
	}
}

// 获取资源列表
func (d *ResourceData) GetResourceList(resourceListReq scheme.ResourceListReq) (scheme.ResourceListResp, error) {
	var (
		err              error
		resourceList     []models.Resource
		resourceListInfo = d.DB.DBClient.Model(&models.Resource{})
		totalRecords     int64
		total            int
		offset           int
	)
	//查询条件
	if resourceListReq.Name != "" {
		resourceListInfo = resourceListInfo.Where("name like ?", "%"+resourceListReq.Name+"%")
	}
	if resourceListReq.Path != "" {
		resourceListInfo = resourceListInfo.Where("path like ?", "%"+resourceListReq.Path+"%")
	}
	if resourceListReq.Type != 0 {
		resourceListInfo = resourceListInfo.Where("type = ?", resourceListReq.Type)
	}
	if resourceListReq.IsStatus != 0 {
		resourceListInfo = resourceListInfo.Where("status = ?", resourceListReq.IsStatus)
	}
	//记录总数
	err = resourceListInfo.Count(&totalRecords).Error
	if err != nil {
		return scheme.ResourceListResp{}, err
	}
	//总页数  总记录数+（每页数量-1）/每页数量
	total = int((totalRecords + int64(resourceListReq.Limit) - 1) / int64(resourceListReq.Limit))
	//分页
	if resourceListReq.Limit > 0 && resourceListReq.Page > 0 {
		offset = (resourceListReq.Page - 1) * resourceListReq.Limit
		resourceListInfo = resourceListInfo.Offset(offset).Limit(resourceListReq.Limit)
	}
	err = resourceListInfo.Find(&resourceList).Error
	if err != nil {
		return scheme.ResourceListResp{}, err
	}
	//返回
	resourceListResp := scheme.ResourceListResp{
		ResourceList: resourceList,
		Total:        total,
		Pagination: scheme.Pagination{
			Limit: resourceListReq.Limit,
			Page:  resourceListReq.Page,
		},
	}
	return resourceListResp, err
}

// 查询资源
func (d *ResourceData) GetResource(resourceGetReq scheme.ResourceGetReq) (models.Resource, error) {
	var (
		resourceData models.Resource
		err          error
	)
	err = d.DB.DBClient.Model(&models.Resource{}).Where("id = ?", resourceGetReq.Id).First(&resourceData).Error
	if err != nil {
		return models.Resource{}, err
	}
	return resourceData, nil
}

// 增加资源
func (d *ResourceData) CreateResource(createResourceReq scheme.ResourceCreateReq) (models.Resource, error) {
	createResourceData := models.Resource{
		Pid:         createResourceReq.Pid,
		Name:        createResourceReq.Name,
		Description: createResourceReq.Description,
		Path:        createResourceReq.Path,
		Type:        createResourceReq.Type,
		Status:      createResourceReq.Status,
	}
	err := d.DB.DBClient.Model(&models.Resource{}).Create(&createResourceData).Error
	if err != nil {
		return models.Resource{}, err
	}
	return createResourceData, nil
}

// 更新资源
func (d *ResourceData) UpdateResource(updateResourceReq scheme.ResourceUpdateReq) (models.Resource, wapper.ErrorCode) {
	updateResourceData := models.Resource{
		Id:          updateResourceReq.Id,
		Pid:         updateResourceReq.Pid,
		Name:        updateResourceReq.Name,
		Description: updateResourceReq.Description,
		Path:        updateResourceReq.Path,
		Type:        updateResourceReq.Type,
		Status:      updateResourceReq.Status,
		UpdatedBy:   updateResourceReq.UpdateBy,
	}
	err := d.DB.DBClient.Model(&models.Resource{}).Where("id = ?", updateResourceReq.Id).Updates(&updateResourceData).Error
	if err != nil {
		return models.Resource{}, wapper.UpdateResourceFailed
	}
	var updateDResource models.Resource
	err = d.DB.DBClient.Model(&models.Resource{}).Where("id = ?", updateResourceReq.Id).First(&updateDResource).Error
	if err != nil {
		return models.Resource{}, wapper.DataNotFound
	}
	return updateDResource, wapper.Success

}

// 删除资源
func (d *ResourceData) DelResource(resourceDelReq scheme.ResourceDelReq) wapper.ErrorCode {
	err := d.DB.DBClient.Model(&models.Resource{}).Where("id = ?", resourceDelReq.Id).Update("status", scheme.StatusOff).Error
	if err != nil {
		return wapper.DelResourceFailed
	}
	return wapper.Success

}
