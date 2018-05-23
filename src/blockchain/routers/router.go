package routers

import (
	"blockchain/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("api/blocks/",
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
