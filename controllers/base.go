package controllers

import (
	"encoding/json"
	"prototype/dto"

	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
}

func (c *BaseController) checkLogin() bool {
	auth := c.GetSession("auth")
	if isLogin, ok := auth.(bool); ok {
		return isLogin
	}
	return false
}

// @title 返回错误消息
func (c *BaseController) sendError(msg string, code int) {
	if code >= 400 && code < 600 {
		c.Ctx.ResponseWriter.WriteHeader(code)
	}
	// c.Data["json"] = map[string]interface{}{"code": code, "msg": msg}
	c.Data["json"] = dto.ErrorResponse{Code: code, Msg: msg}
	_ = c.ServeJSON()
	c.StopRun()
}

// @title 返回Json消息
func (c *BaseController) sendJson(data interface{}) {
	c.Data["json"] = data
	_ = c.ServeJSON()
	c.StopRun()
}

// @title 返回成功消息
func (c *BaseController) sendSuccess(msg string) {
	// c.Data["json"] = map[string]interface{}{"code": 200, "msg": msg}
	c.Data["json"] = dto.ErrorResponse{Code: 200, Msg: msg}
	_ = c.ServeJSON()
	c.StopRun()
}

// @title 返回文本消息
func (c *BaseController) sendText(msg string) {
	c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	c.Ctx.WriteString(msg)
	c.StopRun()
}

// @title 获取json结构体
func (c *BaseController) getJson() map[string]interface{} {
	var body map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err != nil {
		c.sendError("json解析失败", 400)
	}
	return body
}
