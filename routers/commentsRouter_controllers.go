package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["prototype/controllers:LoginController"] = append(beego.GlobalControllerRouter["prototype/controllers:LoginController"],
        beego.ControllerComments{
            Method: "LoginIn",
            Router: "/api/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["prototype/controllers:LoginController"] = append(beego.GlobalControllerRouter["prototype/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Info",
            Router: "/api/login/info",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["prototype/controllers:LoginController"] = append(beego.GlobalControllerRouter["prototype/controllers:LoginController"],
        beego.ControllerComments{
            Method: "LoginOut",
            Router: "/api/login/out",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["prototype/controllers:PrototypeController"] = append(beego.GlobalControllerRouter["prototype/controllers:PrototypeController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/api/prototype",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["prototype/controllers:PrototypeController"] = append(beego.GlobalControllerRouter["prototype/controllers:PrototypeController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/api/prototype",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["prototype/controllers:PrototypeController"] = append(beego.GlobalControllerRouter["prototype/controllers:PrototypeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/api/prototype",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["prototype/controllers:PrototypeController"] = append(beego.GlobalControllerRouter["prototype/controllers:PrototypeController"],
        beego.ControllerComments{
            Method: "File",
            Router: "/api/prototype/file",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["prototype/controllers:PrototypeController"] = append(beego.GlobalControllerRouter["prototype/controllers:PrototypeController"],
        beego.ControllerComments{
            Method: "Project",
            Router: "/api/prototype/project",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
