package controllers

import (
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

// BaseController 基础控制器
type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
}

// @title 返回错误消息
// @description 发生错误时使用
// @auth kay
// @param msg string
// @param code int
func (c *BaseController) sendError(msg string, code int) {
	if code >= 400 && code < 600 {
		c.Ctx.ResponseWriter.WriteHeader(code)
	}
	c.Data["json"] = map[string]interface{}{"code": code, "msg": msg}
	_ = c.ServeJSON()
	c.StopRun()
}

// sendJson
// @title 返回Json消息
// @description 一般用于返回结构体
// @auth kay
// @param data interface{}
func (c *BaseController) sendJson(data interface{}) {
	c.Data["json"] = data
	_ = c.ServeJSON()
	c.StopRun()
}

// @title 返回成功消息
// @description 操作类接口返回成功消息
// @auth kay
// @param msg string
func (c *BaseController) sendSuccess(msg string) {
	c.Data["json"] = map[string]interface{}{"code": 200, "msg": msg}
	_ = c.ServeJSON()
	c.StopRun()
}

// @title 返回文本消息
// @description 一般用于返回文本
// @auth kay
// @param msg string
func (c *BaseController) sendText(msg string) {
	c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	c.Ctx.WriteString(msg)
	c.StopRun()
}

// @title 获取json结构体
// @auth kay
// @return map[string]interface{}
func (c *BaseController) getJson() map[string]interface{} {
	var body map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err != nil {
		c.sendError("json解析失败", 400)
	}
	return body
}
