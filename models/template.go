/**********************************************
** @Des: markdown模板
** @Author: haodaquan
** @Date:   2018-01-16 15:42:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-01-16 11:48:17
***********************************************/
package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	TEMPLATE_DB_NAME = "set_template"
)

type Template struct {
	Id           int    //	唯一标识
	TemplateName string //模板名字
	Detail       string //模板的详情
	Status       int    //模板的状态
	CreateId     int    //创建者
	UpdateId     int    //更新者
	CreateTime   int64  //创建时间
	UpdateTime   int64  //更新时间
}

func (a *Template) TableName() string {
	return TableName(TEMPLATE_DB_NAME)
}

func TemplateGetList(page, pageSize int, filters ...interface{}) ([]*Template, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Template, 0)
	query := orm.NewOrm().QueryTable(TableName(TEMPLATE_DB_NAME))
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

func TemplateAdd(a *Template) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func TemplateGetById(id int) (Template, error) {
	var list Template
	query := orm.NewOrm().QueryTable(TableName(TEMPLATE_DB_NAME))
	query.Filter("id", id).Filter("status", 1).One(&list)
	return list, nil
}

func (a *Template) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}

func (a *Template) Delete(id int64, update_id int) (int64, error) {
	sql := QueryBuilder().
		Update(TableName(TEMPLATE_DB_NAME)).
		Set("status=0").
		Set("update_id=?").
		Update("update_time=?").
		Where("id=?").
		String()
	res, err := orm.NewOrm().Raw(sql, update_id, time.Now().Unix(), id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		return num, nil
	}
	return 0, err
}
