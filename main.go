package main

import (
	"encoding/json"
	"net/http"
	"prototype/models"
	_ "prototype/routers"

	// _ "prototype/tasks"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

// @title          原型管理api
// @description    学海内部适用
// @version        1.0
// @contact.name   aiChenK
// @contact.email  aichenk@qq.com

// @tag.name login
// @tag.description 登录
// @tag.name prototype
// @tag.description 原型管理

func main() {

	//开启session
	beego.BConfig.WebConfig.Session.SessionOn = true

	//开启文件浏览
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/static"] = "static"

	// 更改404错误提示
	beego.ErrorHandler("404", func(writer http.ResponseWriter, request *http.Request) {
		data := map[string]interface{}{"code": 404, "msg": "404 Not Found"}
		content, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		_, _ = writer.Write(content)
	})

	models.Init()

	// toolbox.StartTask()
	// defer toolbox.StopTask()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		beego.Run("127.0.0.1")
	} else {
		beego.Run()
	}
}
