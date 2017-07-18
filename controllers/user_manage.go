package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ss1917/common/utils"
	"github.com/ss1917/ss_admin/models"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Prepare() {
	Username, User_id, Regionid, _ := utils.LoginAuthentication(beego.AppConfig.String("sso_uri"))

	if Username == "" {
		this.Ctx.Redirect(302, beego.AppConfig.String("login_uri"))
	}
	if Regionid != 0 {

	}

	this.Data["sysmanage"] = true
	this.Data["usermanage"] = true
	this.Data["user_id"] = User_id
	this.Data["username"] = Username
	this.Data["regionid"] = Regionid
}

// @router / [get]
func (this *UserController) Get() {
	userinfo, _ := models.UserGetAll()
	this.Data["userinfo"] = userinfo
	this.TplName = "manage/user.html"
}

// @router / [post]
func (this *UserController) Post() {
	this.Data["json"] = map[string]interface{}{
		"status": -2,
		"msg":    "密码错误",
	}
	this.ServeJSON()
}

// @router / [patch]
func (this *UserController) Patch() {
	this.Data["json"] = map[string]interface{}{
		"status": -2,
		"msg":    "密码错误",
	}
	this.ServeJSON()
}

// @router /:uid [delete]
func (this *UserController) Delete() {
	uid := this.GetString(":uid")
	if uid != "" {
		if Uid, err := strconv.Atoi(uid); err != nil {
			fmt.Println(Uid)
			if err := models.DeleteUser(Uid); err != nil {
				this.Data["json"] = map[string]interface{}{
					"status": 0,
					"msg":    "删除成功",
				}
				this.ServeJSON()
			}
		}
	}

	this.Data["json"] = map[string]interface{}{
		"status": -1,
		"msg":    "",
	}
	this.ServeJSON()
}
