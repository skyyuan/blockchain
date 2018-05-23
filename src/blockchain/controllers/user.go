package controllers

import (
	"blockchain/utils"
	"blockchain/models"
	"github.com/astaxie/beego"
	"fmt"
	"blockchain/blocks"
)
type UserController struct {
	beego.Controller
}

// @router / [get]
func (this *UserController) Index() {
	//mdb, mSession := utils.GetMgoDbSession()
	//defer mSession.Close()
	this.Data["json"] = 111
	this.ServeJSON()
}


// @router / [post]
func (this *UserController) Create() {
	mdb, mSession := utils.GetMgoDbSession()
	defer mSession.Close()
	name := this.GetString("name")
	account := this.GetString("account")
	pwd := this.GetString("pwd")
	_, err := models.Newuser(mdb, name, account, pwd)
	var code string = "添加成功"
	if err != nil {
		code = "添加失败"
	}
	this.Data["json"] = code
	this.ServeJSON()
}

// @router /show [get]
func (this *UserController) Edit() {
	mdb, mSession := utils.GetMgoDbSession()
	defer mSession.Close()
	name := this.GetString("name")
	user,_ := models.GetUser(mdb, name)
	this.Data["json"] = user
	this.ServeJSON()
	return
}

// @router /:id/update [post]
func (this *UserController) Update() {

}

// @router /:id/delete [post]
func (this *UserController) Destroy() {

}

// @router /verify [get]
func (this *UserController) VerifyByAccount() {
	account := this.GetString("account")
	fmt.Println(account)
	mdb, mSession := utils.GetMgoDbSession()
	defer mSession.Close()
	var flag string
	user, err := models.GetUserByAccount(mdb, account)
	if err != nil {
		flag = "数据不存在"
		this.Data["json"] = flag
		this.ServeJSON()
		return
	}
	flag = blocks.Verify(user.Id_.Hex())

	this.Data["json"] = flag
	this.ServeJSON()
}
