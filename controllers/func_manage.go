package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ss1917/common/utils"
	"github.com/ss1917/ss_admin/models"
	"strconv"
)

type FuncController struct {
	beego.Controller
}

//鉴定登录和权限
func (this *FuncController) Prepare() {
	//登录验证，如果没登录自动跳转登录
	Username, User_id, Regionid, _ := utils.LoginAuthentication(beego.AppConfig.String("sso_uri"))
	if Username != "" {
		this.Ctx.Redirect(302, beego.AppConfig.String("login_uri"))
	}

	if Regionid != 0 {
		this.Data["json"] = map[string]interface{}{
			"status": -1,
			"msg":    "没有权限",
		}
		this.ServeJSON()
		this.StopRun()
	}

	this.Data["sysmanage"] = true
	this.Data["funcmanage"] = true
	this.Data["user_id"] = User_id
	this.Data["username"] = Username
	this.Data["regionid"] = Regionid
}

// @router / [get]
func (this *FuncController) Get() {
	//models.OperationUpdate
	models.OperationLog(this.Data["regionid"].(int),this.Data["username"].(string),"访问","权限列表")
	fmt.Println(this.Data["username"].(string))
	this.TplName = "manage/role.html"
}

// @router / [post]
func (this *FuncController) Post() {
	this.Data["json"] = map[string]interface{}{
		"status": -2,
		"msg":    "密码错误",
	}
	this.ServeJSON()
}

// @router / [patch]
func (this *FuncController) Patch() {
	this.Data["json"] = map[string]interface{}{
		"status": -2,
		"msg":    "密码错误",
	}
	this.ServeJSON()
}

// @router /:uid [delete]
func (this *FuncController) Delete() {
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
