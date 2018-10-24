package models

import (
	"github.com/astaxie/beego/orm"
)

// TableName 设置表名
func (d *DomainManager) TableName() string {
	return DomainManagerTBName()
}

// DomainManagerQueryParam 用于搜索的类
type DomainManagerQueryParam struct {
	BaseQueryParam
	NameLike string
}

// DomainManager 域名管理对应的表
type DomainManager struct {
	Id     int    `form:"Id"`
	Client string `form:"Client"`
	Domain string `form:"Domain"`
}

// DomainManagerPageList 获取分页数据
func DomainManagerPageList(params *DomainManagerQueryParam) ([]*DomainManager, int64) {
	query := orm.NewOrm().QueryTable(DomainManagerTBName())
	data := make([]*DomainManager, 0)
	sortorder := "Id"

	query = query.Filter("client__istartswith", params.NameLike)
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)

	return data, total
}

// DomainManagerDataList 获取数据
func DomainManagerDataList(params *DomainManagerQueryParam) []*DomainManager {
	params.Limit = -1
	params.Sort = "Id"
	params.Order = "asc"
	data, _ := DomainManagerPageList(params)
	return data
}

// DomainManagerBatchDelete 批量删除
func DomainManagerBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(DomainManagerTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}
