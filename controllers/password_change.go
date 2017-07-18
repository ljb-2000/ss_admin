package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ss1917/common/utils"
	"github.com/ss1917/ss_admin/models"
	"time"
)

type PasswordController struct {
	beego.Controller
}

func (this *PasswordController) Prepare() {
	Username, User_id, Regionid, _ := utils.LoginAuthentication(beego.AppConfig.String("sso_uri"))
	if Username == "" {
		this.Ctx.Redirect(302, beego.AppConfig.String("login_uri"))
	}
	if Regionid != 0 {

	}

	this.Data["sysmanage"] = true
	this.Data["changepassword"] = true
	this.Data["user_id"] = User_id
	this.Data["username"] = Username
	this.Data["regionid"] = Regionid
}

// @router / [get]
func (this *PasswordController) Get() {
	this.TplName = "manage/password.change.html"
}

// @router / [post]
func (this *PasswordController) Post() {
	type PasswordInfo struct {
		Oldpassword  string `json:"oldpassword"`
		Newpassword1 string `json:"newpassword1"`
		Newpassword2 string `json:"newpassword2"`
	}

	var passwordinfo *PasswordInfo = new(PasswordInfo)
	json.Unmarshal(this.Ctx.Input.RequestBody, passwordinfo)

	user, _ := models.UserGetById(this.Data["user_id"].(int))
	fmt.Println(this.Data["username"], passwordinfo)
	if utils.GenMD5(passwordinfo.Oldpassword) == user.Password {
		if passwordinfo.Newpassword1 == passwordinfo.Newpassword2 {
			user.Last_login = time.Now()
			user.Password = utils.GenMD5(passwordinfo.Newpassword1)
			models.UserUpdate(user)
			this.Data["json"] = map[string]interface{}{
				"status":   0,
				"msg":      "修改成功",
				"next_uri": beego.AppConfig.String("login_uri"),
			}
			this.ServeJSON()

		} else {
			this.Data["json"] = map[string]interface{}{
				"status": -1,
				"msg":    "输入不一致",
			}
			this.ServeJSON()
		}

	} else {
		this.Data["json"] = map[string]interface{}{
			"status": -2,
			"msg":    "密码错误",
		}
		this.ServeJSON()
	}
}
