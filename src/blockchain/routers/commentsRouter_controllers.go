package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["blockchain/controllers:UserController"] = append(beego.GlobalControllerRouter["blockchain/controllers:UserController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blockchain/controllers:UserController"] = append(beego.GlobalControllerRouter["blockchain/controllers:UserController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blockchain/controllers:UserController"] = append(beego.GlobalControllerRouter["blockchain/controllers:UserController"],
		beego.ControllerComments{
			Method: "Destroy",
			Router: `/:id/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blockchain/controllers:UserController"] = append(beego.GlobalControllerRouter["blockchain/controllers:UserController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/:id/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blockchain/controllers:UserController"] = append(beego.GlobalControllerRouter["blockchain/controllers:UserController"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/show`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blockchain/controllers:UserController"] = append(beego.GlobalControllerRouter["blockchain/controllers:UserController"],
		beego.ControllerComments{
			Method: "VerifyByAccount",
			Router: `/verify`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
