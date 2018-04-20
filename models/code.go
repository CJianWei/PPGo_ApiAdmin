/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-16 15:42:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-25 11:48:17
***********************************************/
package models

import (
	"github.com/astaxie/beego/orm"
)

type Code struct {
	Id         int
	Code       string
	Desc       string
	Detail     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

const SET_CODE = "set_code"

func (a *Code) TableName() string {
	return TableName(SET_CODE)
}

func CodeAdd(a *Code) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func CodeGetByName(CodeName string) (*Code, error) {
	a := new(Code)
	err := orm.NewOrm().QueryTable(TableName(SET_CODE)).Filter("code", CodeName).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func CodeGetList(page, pageSize int, filters ...interface{}) ([]*Code, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Code, 0)
	query := orm.NewOrm().QueryTable(TableName(SET_CODE))
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

func CodeGetByIds(ids string) ([]*Code, error) {
	list := make([]*Code, 0)
	sql := QueryBuilder().
		Set("*").
		From(TableName(SET_CODE)).
		Where("id in (?)").
		String()
	orm.NewOrm().Raw(sql, ids).QueryRows(&list)

	return list, nil
}

func CodeGetById(id int) (*Code, error) {
	r := new(Code)
	err := orm.NewOrm().QueryTable(TableName(SET_CODE)).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *Code) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
