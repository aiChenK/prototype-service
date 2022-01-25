package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Prototype struct {
	Id          uint      `orm:"pk;auto" json:"id" form:"-"`
	Name        string    `orm:"size(32)" json:"name" form:"name" valid:"Required;MaxSize(32)"`
	ProjectName string    `orm:"size(32);index" json:"projectName" form:"projectName" valid:"Required;MaxSize(32)"`
	Path        string    `orm:"size(255)" json:"path" form:"path" valid:"Required;MaxSize(255)"`
	IsDel       uint8     `orm:"size(1);default(0)" json:"isDel"`
	StartDate   time.Time `orm:"type(date);null" json:"startDate" form:"-"`
	EndDate     time.Time `orm:"type(date);null" json:"endDate" form:"-"`
	CreateTime  time.Time `orm:"type(datetime);auto_now_add" json:"createTime"`
	UpdateTime  time.Time `orm:"type(datetime);auto_now" json:"updateTime"`
}

func init() {
	orm.RegisterModel(new(Prototype))
}
