package controllers

import (
	"fmt"
	"os"
	"prototype/dto"
	"prototype/helper"
	"prototype/models"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
)

type PrototypeController struct {
	BaseController
}

func (c *PrototypeController) Prepare() {
	_, action := c.GetControllerAndAction()
	if helper.InArray(action, []string{"File", "Create", "Edit", "Delete"}) && !c.checkLogin() {
		c.sendError("未登录不可操作", 401)
	}
	c.BaseController.Prepare()
}

// @summary 获取原型列表
// @tags prototype
// @param status query int false "状态，1=正常（默认），2=已删除" enums(1, 2) default(1)
// @param page query int false "页码" default(1)
// @param size query int false "每页数量" default(10)
// @param projectName query string false "产品名(精确）" minLength(30)
// @success 200 {object} helper.Pager{data=[]dto.PrototypeResponse}
// @router /api/prototype [get]
func (c *PrototypeController) List() {
	status, _ := c.GetInt("status", 1)
	projectName := c.GetString("projectName")
	page, _ := c.GetInt("page", 1)
	size, _ := c.GetInt("size", 10)
	team, _ := c.GetInt("team", 0)
	tag, _ := c.GetInt("tag", 0)
	offset := (page - 1) * size

	o := orm.NewOrm()
	qs := o.QueryTable(&models.Prototype{})

	switch status {
	case 1:
		qs = qs.Filter("IsDel", 0)
	case 2:
		qs = qs.Filter("IsDel", 1)
	}

	if projectName != "" {
		qs = qs.Filter("ProjectName", projectName)
	}
	if team != 0 {
		qs = qs.FilterRaw("id", "IN (select prototype_id from tag_bind where is_del = 0 and tag_id = "+strconv.Itoa(team)+")")
	}
	if tag != 0 {
		qs = qs.FilterRaw("id", "IN (select prototype_id from tag_bind where is_del = 0 and tag_id = "+strconv.Itoa(tag)+")")
	}

	total, _ := qs.Count()

	fields := []string{"Id", "Name", "ProjectName", "Path", "CreateTime", "UpdateTime", "StartDate", "EndDate"}
	rows := []orm.Params{}
	qs.
		OrderBy("-Id").
		Limit(size, offset).
		Values(&rows, fields...)

	// 数据格式化
	data := []*dto.PrototypeResponse{}
	pIds := []uint{}
	for _, row := range rows {
		prototype := &dto.PrototypeResponse{}
		data = append(data, prototype.Parse(row))
		pIds = append(pIds, uint(prototype.Id))
	}

	//获取标签
	if len(pIds) > 0 {
		pIdStr := strings.Replace(strings.Trim(fmt.Sprint(pIds), "[]"), " ", ",", -1)
		var maps []struct {
			PrototypeId uint
			Type        uint8
			Id          uint
			Name        string
		}
		orm.NewOrm().
			Raw("select b.prototype_id,t.type,t.id,t.name from tag_bind as b left join tag t on t.id = b.tag_id where b.is_del = 0 and t.is_del = 0 and b.prototype_id in (" + pIdStr + ")").
			QueryRows(&maps)
		for _, row := range data {
			for _, tag := range maps {
				if tag.PrototypeId != row.Id {
					continue
				}
				if tag.Type == 1 {
					row.Team = append(row.Team, dto.TagResponse{Id: tag.Id, Name: tag.Name})
				} else {
					row.Tag = append(row.Tag, dto.TagResponse{Id: tag.Id, Name: tag.Name})
				}
			}
		}
	}

	pager := &helper.Pager{}
	c.sendJson(pager.RunPage(page, size, int(total), data))
}

// @summary 文件上传
// @tags prototype
// @accept mpfd
// @param file formData file true "zip压缩包"
// @success 200 {object} dto.FileResponse
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

// @summary 创建
// @tags prototype
// @param body body dto.PrototypeCreate true "原型"
// @success 200 {object} dto.SuccessResponse
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

	//事务处理
	o := orm.NewOrm()
	// o.Begin()
	txOrm, err := o.Begin()
	if err != nil {
		c.sendError("事务开启失败："+err.Error(), 500)
	}

	//保存
	_, err = txOrm.Insert(&prototypeModel)
	if err != nil {
		c.sendError("保存失败："+err.Error(), 500)
	}

	//获取team tagid
	teamIds, _ := models.DealNewTags(txOrm, body["team"].([]interface{}), 1)
	tagIds, _ := models.DealNewTags(txOrm, body["tag"].([]interface{}), 2)

	//保存tag关联信息
	_, err = models.SavePrototypeTag(txOrm, prototypeModel.Id, append(teamIds, tagIds...))

	if err != nil {
		txOrm.Rollback()
	} else {
		txOrm.Commit()
	}

	c.sendSuccess("操作成功")
}

// @summary 修改
// @tags prototype
// @param id path int true "原型id"
// @param body body dto.PrototypeCreate true "原型"
// @success 200 {object} dto.SuccessResponse
// @router /api/prototype/:id [patch]
func (c *PrototypeController) Edit() {
	//参数解析
	idParam := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idParam)
	body := c.getJson()

	o := orm.NewOrm()

	prototypeModel := models.Prototype{}
	prototypeModel.Id = uint(id)

	//路径变更且非外链则删除
	o.Read(&prototypeModel)
	if prototypeModel.Path != body["path"].(string) && !strings.HasPrefix(prototypeModel.Path, "http") {
		os.RemoveAll(prototypeModel.Path)
	}

	//填充数据
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
	_, err = o.Update(&prototypeModel)
	if err != nil {
		c.sendError("保存失败："+err.Error(), 500)
	}
	c.sendSuccess("操作成功")
}

// @summary 获取产品枚举
// @tags prototype
// @success 200 {array} string
// @router /api/prototype/project [get]
func (c *PrototypeController) Project() {
	var projects orm.ParamsList
	orm.NewOrm().
		Raw("select project_name from (select project_name from prototype where is_del = 0 order by update_time desc) a GROUP BY project_name").
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

// @summary 获取团队枚举
// @tags prototype
// @success 200 {array} []dto.TagResponse
// @router /api/prototype/team [get]
func (c *PrototypeController) Team() {
	var maps []dto.TagResponse
	orm.NewOrm().
		Raw("select id, name from tag where is_del = 0 and type = 1 limit 100").
		QueryRows(&maps)
	c.sendJson(maps)
}

// @summary 获取标签枚举
// @tags prototype
// @success 200 {array} []dto.TagResponse
// @router /api/prototype/tag [get]
func (c *PrototypeController) Tag() {
	var maps []dto.TagResponse
	orm.NewOrm().
		Raw("select id, name from tag where is_del = 0 and type = 2 limit 100").
		QueryRows(&maps)
	c.sendJson(maps)
}

// @summary 删除原型
// @tags prototype
// @param body body []int true "ids"
// @success 200 {object} dto.SuccessResponse
// @router /api/prototype [delete]
func (c *PrototypeController) Delete() {
	//处理数组
	idParam := strings.Split(c.GetString("ids"), ",")
	ids := []int{}
	for _, id := range idParam {
		if id == "" {
			continue
		}
		tmp, _ := strconv.Atoi(id)
		if tmp <= 0 {
			continue
		}
		ids = append(ids, tmp)
	}
	if len(ids) == 0 {
		c.sendError("缺少有效id", 400)
	}

	prototypeModel := &models.Prototype{}
	prototypeModel.DeleteByIds(ids)

	//删除关联
	tagBindModel := &models.TagBind{}
	tagBindModel.DeleteByPrototypeIds(ids)

	//文件删除
	o := orm.NewOrm()
	rows := []orm.Params{}
	o.QueryTable(&models.Prototype{}).
		Filter("id__in", ids).
		Values(&rows, "Id", "Path")
	for _, row := range rows {
		path := row["Path"].(string)
		if strings.HasPrefix(path, "http") {
			continue
		}
		os.RemoveAll(path)
	}

	c.sendSuccess("删除成功")
}
