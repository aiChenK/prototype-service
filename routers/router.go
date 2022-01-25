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

	// ns := web.NewNamespace("/api",
	// 	web.NSNamespace("/login",
	// 		web.NSInclude(
	// 			&controllers.LoginController{},
	// 		),
	// 	),
	// 	web.NSNamespace("/prototype",
	// 		web.NSInclude(
	// 			&controllers.PrototypeController{},
	// 		),
	// 	),
	// )

	// web.AddNamespace(ns)
}
