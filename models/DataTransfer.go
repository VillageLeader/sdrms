package models

import (
	"github.com/astaxie/beego/orm"
)

// TableName 设置表名
func (d *DataTransfer) TableName() string {
	return DataTransferTBName()
}

// DataTransferQueryParam 用于搜索的类
type DataTransferQueryParam struct {
	BaseQueryParam
	NameLike string
}

// DataTransfer 数据传输 实体类
type DataTransfer struct {
	Id   int    `form:"Id"`
	Host string `form:"Host"`
	Addr string `form:"Addr"`
    Path string `form:"Path"`
    Protocol string `form;"Protocol"`
	Seq  int
}

// DataTransferPageList 获取分页数据
func DataTransferPageList(params *DataTransferQueryParam) ([]*DataTransfer, int64) {
	query := orm.NewOrm().QueryTable(DataTransferTBName())
	data := make([]*DataTransfer, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	case "Seq":
		sortorder = "Seq"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}

	query = query.Filter("host__istartswith", params.NameLike)
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// DataTransferDataList 获取域名列表
func DataTransferDataList(params *DataTransferQueryParam) []*DataTransfer {
	params.Limit = -1
	params.Sort = "Id"
	params.Order = "asc"
	data, _ := DataTransferPageList(params)
	return data
}

// DataTransferBatchDelete 批量删除
func DataTransferBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(DataTransferTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}
