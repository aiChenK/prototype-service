package models

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func Init() error {
	// 处理数据库配置
	dbConfig := map[string]string{
		"hostname": "",
		"hostport": "",
		"database": "",
		"username": "",
		"password": "",
		"charset":  "",
	}
	for key, _ := range dbConfig {
		dbConfig[key], _ = beego.AppConfig.String("mysql::" + key)
	}
	if dbConfig["hostport"] == "" {
		dbConfig["hostport"] = "3306"
	}
	if dbConfig["charset"] == "" {
		dbConfig["charset"] = "utf8mb4"
	}

	// 组装dsn
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local",
		dbConfig["username"],
		dbConfig["password"],
		dbConfig["hostname"],
		dbConfig["hostport"],
		dbConfig["database"],
		dbConfig["charset"],
	)

	runMode, _ := beego.AppConfig.String("RunMode")
	isDev := runMode == "dev"

	if isDev {
		orm.Debug = true
	}

	// set default database
	_ = orm.RegisterDataBase(
		"default",
		"mysql",
		dsn,
	)

	orm.RunSyncdb("default", false, true)

	return nil
}
