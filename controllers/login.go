package controllers

type LoginController struct {
	BaseController
}

// @description 判断登录
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

// @description 密码认证
// @router /api/login [post]
func (c *LoginController) LoginIn() {
	body := c.getJson()
	if body["pwd"] != "xh1234" {
		c.sendError("密码不正确", 400)
	}
	c.SetSession("auth", true)
	c.sendSuccess("认证成功")
}

// @description 退出登录
// @router /api/login/out [post]
func (c *LoginController) LoginOut() {
	c.DestroySession()
	c.sendSuccess("登出成功")
}
