package main

import (
	"encoding/json"
	"net/http"
	"prototype/models"
	_ "prototype/routers"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

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

	runMode, _ := beego.AppConfig.String("RunMode")
	isDev := runMode == "dev"

	if isDev {
		beego.Run("127.0.0.1")
	} else {
		beego.Run()
	}
}
