package data

import (
	"NewProject/app/scheme"
	"NewProject/models"
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
		resourceListInfo = resourceListInfo.Where("type like ?", resourceListReq.Type)
	}
	if resourceListReq.IsStatus != 0 {
		resourceListInfo = resourceListInfo.Where("status like ?", resourceListReq.IsStatus)
	}
	//记录总数
	err = resourceListInfo.Count(&totalRecords).Error
	if err != nil {
		return scheme.ResourceListResp{}, err
	}
	//总页数  总记录数+（每页数量-1）/每页数量
	total = int(totalRecords-int64(resourceListReq.Limit-1)) / resourceListReq.Limit
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

//更新资源
//删除资源
