/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-16 15:42:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-17 11:48:17
***********************************************/
package models

import (
	"github.com/astaxie/beego/orm"
)

const (
	ADMIN_DB_NAME = "uc_admin"
)

type Admin struct {
	Id         int    //唯一标识
	LoginName  string //登录名字
	RealName   string //真实名字
	Password   string //密码
	RoleIds    string //角色ID
	Phone      string //手机号码
	Email      string //email
	Salt       string //配置做密码校验
	LastLogin  int64  //最后登录时间
	LastIp     string //最后登录的IP地址
	Status     int    //当前的状态
	CreateId   int    //创建者的ID
	UpdateId   int    //更新者的ID
	CreateTime int64  //创建时间
	UpdateTime int64  //更新时间
}

func (a *Admin) TableName() string {
	return TableName(ADMIN_DB_NAME)
}

func AdminAdd(a *Admin) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func AdminGetByName(loginName string) (*Admin, error) {
	a := new(Admin)
	err := orm.NewOrm().QueryTable(TableName(ADMIN_DB_NAME)).Filter("login_name", loginName).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func AdminGetList(page, pageSize int, filters ...interface{}) ([]*Admin, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Admin, 0)
	query := orm.NewOrm().QueryTable(TableName(ADMIN_DB_NAME))
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

func AdminGetById(id int) (*Admin, error) {
	r := new(Admin)
	err := orm.NewOrm().QueryTable(TableName(ADMIN_DB_NAME)).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *Admin) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
