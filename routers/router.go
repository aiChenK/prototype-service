// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"prototype/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	// web.Get("/", func(ctx *context.Context) {
	// 	ctx.Redirect(302, "/prototype")
	// })

	web.Include(&controllers.LoginController{})
	web.Include(&controllers.PrototypeController{})
}
