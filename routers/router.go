// @APIVersion 1.0.0
// @Title beego admin API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/ss1917/ss_admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/umg/",
			beego.NSNamespace("/func",
				beego.NSInclude(
					&controllers.FuncController{},
				),
			),

			beego.NSNamespace("/role",
				beego.NSInclude(
					&controllers.RoleController{},
				),
			),

			beego.NSNamespace("/user",
				beego.NSInclude(
					&controllers.UserController{},
				),
			),

			beego.NSNamespace("/change_passwd",
				beego.NSInclude(
					&controllers.PasswordController{},
				),
			),
		),

		beego.NSNamespace("/home",
			beego.NSInclude(
				&controllers.HomeController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
