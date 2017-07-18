package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

// @router / [get]
func (this *HomeController) Get() {
	this.Data["Email"] = "astaxie@gmail.com"
	//this.Layout = "base/admin_layout.html"
	this.TplName = "home.html"
}
