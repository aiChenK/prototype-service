package controllers

type LoginController struct {
	BaseController
}

// @summary 获取登录信息
// @description 获取登录信息
// @tags login
// @success 200 {object} dto.SuccessResponse
// @failure 401 {object} dto.ErrorResponse
// @router /api/login/info [get]
func (c *LoginController) Info() {
	auth := c.GetSession("auth")
	if isLogin, ok := auth.(bool); ok {
		if isLogin {
			c.sendSuccess("已登录")
		}
	}
	c.sendError("未登录", 401)
}

// @summary 登录
// @description 密码认证
// @tags login
// @accept json
// @param body body dto.LoginForm true "body"
// @success 200 {object} dto.SuccessResponse
// @failure 400 {object} dto.ErrorResponse
// @router /api/login [post]
func (c *LoginController) LoginIn() {
	body := c.getJson()
	if body["pwd"] != "xh1234" {
		c.sendError("密码不正确", 400)
	}
	c.SetSession("auth", true)
	c.sendSuccess("认证成功")
}

// @summary 登出
// @description 退出认证
// @tags login
// @success 200 {object} dto.SuccessResponse
// @router /api/login/out [post]
func (c *LoginController) LoginOut() {
	c.DestroySession()
	c.sendSuccess("登出成功")
}
