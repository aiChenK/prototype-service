package dto

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type PrototypeResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	ProjectName string `json:"projectName"`
	Path        string `json:"path"`
	IsDel       uint8  `json:"-"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}

func (r *PrototypeResponse) Parse(row orm.Params) *PrototypeResponse {
	r.Id = uint(row["Id"].(uint64))
	r.Name = row["Name"].(string)
	r.ProjectName = row["ProjectName"].(string)
	r.Path = row["Path"].(string)
	if value, ok := row["CreateTime"].(time.Time); ok {
		r.CreateTime = value.Format("2006-01-02 15:04:05")
	}
	if value, ok := row["UpdateTime"].(time.Time); ok {
		r.UpdateTime = value.Format("2006-01-02 15:04:05")
	}
	if value, ok := row["StartDate"].(time.Time); ok {
		r.StartDate = value.Format("2006-01-02")
	}
	if value, ok := row["EndDate"].(time.Time); ok {
		r.EndDate = value.Format("2006-01-02")
	}
	return r
}

type PrototypeCreate struct {
	Name        string `json:"name" valid:"Required;MaxSize(32)"`
	ProjectName string `json:"projectName" valid:"Required;MaxSize(32)"`
	Path        string `json:"path" valid:"Required;MaxSize(255)"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
}

type FileResponse struct {
	Path string `json:"path"`
}
