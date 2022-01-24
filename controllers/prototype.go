package controllers

import (
	"fmt"
	"os"
	"prototype/helper"
	"prototype/models"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
)

type PrototypeController struct {
	BaseController
}

// @description 原型列表
// @router /api/prototype [get]
func (c *PrototypeController) List() {
	status, _ := c.GetInt("status", 1)
	projectName := c.GetString("projectName")
	page, _ := c.GetInt("page", 1)
	size, _ := c.GetInt("size", 20)
	offset := (page - 1) * size

	o := orm.NewOrm()
	qs := o.QueryTable(&models.Prototype{})

	rows := []orm.Params{}
	switch status {
	case 1:
		qs = qs.Filter("IsDel", 0)
	case 2:
		qs = qs.Filter("IsDel", 1)
	}

	if projectName != "" {
		qs = qs.Filter("ProjectName", projectName)
	}

	total, _ := qs.Count()
	qs.
		OrderBy("-Id").
		Limit(size, offset).
		Values(&rows, "Id", "Name", "ProjectName", "Path", "CreateTime", "UpdateTime", "StartDate", "EndDate")

	//数据格式化
	nowTime := time.Now().Unix()
	for _, row := range rows {
		if value, ok := row["CreateTime"].(time.Time); ok {
			row["CreateTime"] = value.Format("2006-01-02 15:04:05")
		} else {
			row["CreateTime"] = ""
		}
		if value, ok := row["UpdateTime"].(time.Time); ok {
			row["UpdateTime"] = value.Format("2006-01-02 15:04:05")
		} else {
			row["UpdateTime"] = ""
		}
		if value, ok := row["StartDate"].(time.Time); ok {
			row["StartDate"] = value.Format("2006-01-02")
		} else {
			row["StartDate"] = ""
		}
		if value, ok := row["EndDate"].(time.Time); ok {
			row["IsExprie"] = value.Unix() < nowTime
			row["EndDate"] = value.Format("2006-01-02")
		} else {
			row["IsExprie"] = false
			row["EndDate"] = ""
		}
	}

	pager := &helper.Pager{}

	c.sendJson(pager.RunPage(page, size, int(total), rows))
}

// @description 文件上传
// @router /api/prototype/file [post]
func (c *PrototypeController) File() {
	nowTime := fmt.Sprintf("%d", time.Now().UnixMicro())
	file, fileHead, err := c.GetFile("file")
	if err != nil {
		c.sendError("文件上传失败："+err.Error(), 400)
	}
	defer file.Close()

	//保存到临时目录(需删除)
	tempPath := "static/tmp/" + nowTime
	c.SaveToFile("file", tempPath)
	defer os.Remove(tempPath)

	//解压到指定目录并删除缓存
	savePath := "static/files/" + time.Now().Format("2006-01") + "/" + nowTime
	err = helper.UnZipWithReplacePath(tempPath, savePath, helper.GetFileBaseNameOnly(fileHead.Filename))
	if err != nil {
		panic(err)
	}

	c.sendJson(map[string]interface{}{"path": savePath})
}

// @description 保存原型
// @router /api/prototype [post]
func (c *PrototypeController) Create() {
	//参数解析
	body := c.getJson()

	prototypeModel := models.Prototype{}
	prototypeModel.Name = body["name"].(string)
	prototypeModel.ProjectName = body["projectName"].(string)
	prototypeModel.Path = body["path"].(string)
	if body["startDate"] != "" && body["startDate"] != nil {
		prototypeModel.StartDate, _ = time.Parse("2006-01-02", body["startDate"].(string))
	}
	if body["endDate"] != "" && body["endDate"] != nil {
		prototypeModel.EndDate, _ = time.Parse("2006-01-02", body["endDate"].(string))
	}

	//参数验证
	valid := validation.Validation{}
	isValid, err := valid.Valid(&prototypeModel)
	if err != nil {
		c.sendError("验证有误："+err.Error(), 400)
	}
	if !isValid {
		for _, err := range valid.Errors {
			c.sendError("参数有误："+err.Error(), 400)
		}
	}

	//保存
	o := orm.NewOrm()
	_, err = o.Insert(&prototypeModel)
	if err != nil {
		c.sendError("保存失败："+err.Error(), 500)
	}
	c.sendSuccess("操作成功")
}

// @description 获取产品枚举
// @router /api/prototype/project [get]
func (c *PrototypeController) Project() {
	var projects orm.ParamsList
	orm.NewOrm().
		Raw("select project_name from (select project_name from prototype order by update_time desc) a GROUP BY project_name").
		ValuesFlat(&projects)
	c.sendJson(projects)

	// o := orm.NewOrm()
	// qs := o.QueryTable(&models.Prototype{})

	// rows := []orm.ParamsList{}
	// qs.
	// 	Distinct().
	// 	ValuesList(&rows, "ProjectName")

	// //转换格式
	// projects := []string{}
	// for _, row := range rows {
	// 	projects = append(projects, row[0].(string))
	// }

	// c.sendJson(projects)
}

// func (c *PrototypeController) DeleteView() {
// 	idParam := c.Ctx.Input.Param(":id")
// 	idInt, err := strconv.ParseInt(idParam, 10, 32)
// 	if err != nil {
// 		panic(err)
// 	}
// 	id := uint32(idInt)

// 	o := orm.NewOrm()
// 	prototypeModel := models.Prototype{Id: id}
// 	err = o.Read(&prototypeModel)
// 	if err != nil {
// 		panic(err)
// 	}

// 	prototypeModel.IsDel = 1
// 	_, err = o.Update(&prototypeModel, "IsDel", "UpdateTime")
// 	if err != nil {
// 		panic(err)
// 	}

// 	c.Redirect("/prototype", 302)
// }

// func (c *PrototypeController) Upload() {

// 	//参数解析
// 	prototypeModel := models.Prototype{}
// 	if err := c.ParseForm(&prototypeModel); err != nil {
// 		c.sendError("表达解析失败"+err.Error(), 400)
// 	}

// 	//参数验证
// 	valid := validation.Validation{}
// 	isValid, err := valid.Valid(&prototypeModel)
// 	if err != nil {
// 		c.sendError("验证有误"+err.Error(), 400)
// 	}
// 	if !isValid {
// 		for _, err := range valid.Errors {
// 			c.sendError("参数有误："+err.Error(), 400)
// 		}
// 	}

// 	//文件处理
// 	nowTime := fmt.Sprintf("%d", time.Now().UnixMicro())
// 	file, fileHead, err := c.GetFile("file")
// 	if err != nil {
// 		c.sendError("文件上传失败："+err.Error(), 400)
// 	}
// 	defer file.Close()
// 	tempPath := "static/tmp/" + nowTime
// 	c.SaveToFile("file", tempPath)

// 	prototypeModel.Path = "static/files/" + nowTime

// 	//解压到指定目录并删除缓存
// 	err = helper.UnZipWithReplacePath(tempPath, prototypeModel.Path, helper.GetFileBaseNameOnly(fileHead.Filename))
// 	if err != nil {
// 		panic(err)
// 	}
// 	os.Remove(tempPath)

// 	//保存
// 	o := orm.NewOrm()
// 	_, err = o.Insert(&prototypeModel)
// 	if err != nil {
// 		c.sendError("保存失败"+err.Error(), 500)
// 	}
// 	// c.sendSuccess("保存成功")

// 	flash := web.NewFlash()
// 	flash.Success("上传成功")
// 	flash.Store(&c.Controller)
// 	c.Redirect("/prototype", 302)
// }
