/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-16 15:42:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-10-09 11:48:17
***********************************************/
package models

import (
	"github.com/astaxie/beego/orm"
)

type ApiDetail struct {
	Id         int    // 唯一的标识
	SourceId   int    // 隶属于哪块资源
	Method     int    // 采用的是什么方法
	ApiName    string // api的命名
	ApiUrl     string // api的path
	Detail     string // api的详情
	Status     int    // api的当前状态
	CreateId   int    // 创建ID
	AuditId    int    // 审查ID
	UpdateId   int    // 更新ID
	CreateTime int64  // 创建时间
	UpdateTime int64  // 更新时间
	AuditTime  int64  // 审查时间
}
type ApiDetails struct {
	ApiDetail
	CreateName string
	UpdateName string
	AuditName  string
}

const (
	API_SOURCE = "api_source"
	API_DETAIL = "api_detail"
)

type ApiSourceList struct {
	Id         int
	SourceName string
	ApiLists   []*ApiList
}

type ApiList struct {
	Id      int
	Method  int
	ApiName string
}

func ApiTreeData(groupId int) ([]*ApiSourceList, error) {
	list := make([]*ApiSourceList, 0)
	_, err := orm.NewOrm().
		QueryTable(TableName(API_SOURCE)).
		Filter("group_id", groupId).
		All(&list, "id", "source_name")
	if err != nil {
		return nil, err
	}
	apiList := make([]*ApiSourceList, 0)
	for _, apisource := range list {
		detailList := make([]*ApiList, 0)
		orm.NewOrm().
			QueryTable(TableName(API_DETAIL)).
			Filter("status", 3).
			Filter("source_id", apisource.Id).
			All(&detailList, "id", "method", "api_name")
		apisource.ApiLists = detailList
		apiList = append(apiList, apisource)
	}
	return apiList, nil
}

func (a *ApiDetail) TableName() string {
	return TableName(API_DETAIL)
}

func ApiDetailAdd(a *ApiDetail) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func ApiDetailGetById(id int) (*ApiDetail, error) {
	r := new(ApiDetail)
	err := orm.NewOrm().
		QueryTable(TableName(API_DETAIL)).
		Filter("id", id).
		One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func CopyDetail(from *ApiDetail, to *ApiDetails) {
	to.Id = from.Id
	to.SourceId = from.SourceId
	to.Method = from.Method
	to.ApiName = from.ApiName
	to.ApiUrl = from.ApiUrl
	to.Detail = from.Detail
	to.Status = from.Status
	to.CreateId = from.CreateId
	to.AuditId = from.AuditId
	to.UpdateId = from.UpdateId
	to.CreateTime = from.CreateTime
	to.UpdateTime = from.UpdateTime
	to.AuditTime = from.AuditTime
}

func ApiFullDetailById(id int) (detail *ApiDetails, err error) {
	detail_tmp, err := ApiDetailGetById(id)
	if err != nil {
		return
	}
	ids := []int{detail_tmp.CreateId, detail_tmp.AuditId, detail_tmp.UpdateId}
	var admins []*Admin
	_, err = orm.NewOrm().QueryTable(TableName(ADMIN_DB_NAME)).Filter("id__in", ids).All(&admins, "id", "real_name")
	if err != nil {
		return
	}
	var idx = map[int]string{}
	for _, admin := range admins {
		idx[admin.Id] = admin.RealName
	}
	detail = &ApiDetails{
		AuditName:  idx[detail_tmp.AuditId],
		CreateName: idx[detail_tmp.CreateId],
		UpdateName: idx[detail_tmp.UpdateId],
	}
	CopyDetail(detail_tmp, detail)
	return
}

func ApiDetailGetList(page, pageSize int, filters ...interface{}) ([]*ApiDetail, int64) {
	offset := (page - 1) * pageSize
	list := make([]*ApiDetail, 0)
	query := orm.NewOrm().QueryTable(TableName(API_DETAIL))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

func (a *ApiDetail) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}

func ApiChangeStatus(ids string, status int) (num int64, err error) {
	var sql = QueryBuilder().
		Update(TableName(API_DETAIL)).
		Set("status=?").
		Where("id in (?)").
		String()
	res, err := orm.NewOrm().Raw(sql, status, ids).Exec()
	num = 0
	if err == nil {
		num, _ = res.RowsAffected()
	}
	return num, err
}
