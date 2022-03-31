package models

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Tag struct {
	Id         uint      `orm:"pk;auto" json:"id" form:"-"`
	Name       string    `orm:"size(32)" json:"name" form:"name" valid:"Required;MaxSize(32)"`
	Type       uint8     `orm:"default(1)" json:"type" description:"1=team，2=tag" valid:"Required"`
	IsDel      uint8     `orm:"size(1);default(0)" json:"isDel"`
	CreateTime time.Time `orm:"type(datetime);auto_now_add" json:"createTime"`
	UpdateTime time.Time `orm:"type(datetime);auto_now" json:"updateTime"`
}

type TagBind struct {
	Id          uint      `orm:"pk;auto" json:"id" form:"-"`
	TagId       uint      `orm:"index;default(0)"`
	PrototypeId uint      `orm:"index;default(0)"`
	IsDel       uint8     `orm:"size(1);default(0)" json:"isDel"`
	CreateTime  time.Time `orm:"type(datetime);auto_now_add" json:"createTime"`
	UpdateTime  time.Time `orm:"type(datetime);auto_now" json:"updateTime"`
}

func (m *Tag) TableUnique() [][]string {
	return [][]string{
		[]string{"Type", "Name"},
	}
}

func init() {
	orm.RegisterModel(new(Tag), new(TagBind))
}

// 根据原型id删除关联
func (m *TagBind) DeleteByPrototypeIds(ids []int) (num int64, err error) {
	o := orm.NewOrm()
	return o.QueryTable(&TagBind{}).
		Filter("PrototypeId__in", ids).
		Update(orm.Params{
			"is_del":      1,
			"update_time": time.Now().UTC(),
		})
}

// 处理前端数据，string则创建新tag
func DealNewTags(o orm.TxOrmer, tags []interface{}, tagType uint8) (ids []uint, err error) {
	for _, tag := range tags {
		id, ok := tag.(float64)
		if ok {
			ids = append(ids, uint(id))
		} else {
			// 新增
			tagModel := new(Tag)
			tagModel.Type = tagType
			tagModel.Name = fmt.Sprintf("%v", tag)
			_, err = o.Insert(tagModel)
			if err != nil {
				return
			}
			ids = append(ids, tagModel.Id)
		}
	}
	return
}

// 保存原型对应tagId
func SavePrototypeTag(o orm.TxOrmer, prototypeId uint, tagIds []uint) (int64, error) {
	tagBindModels := []TagBind{}
	for _, tagId := range tagIds {
		tagBindModels = append(tagBindModels, TagBind{
			TagId:       tagId,
			PrototypeId: prototypeId,
		})
	}
	return o.InsertMulti(100, tagBindModels)
}
