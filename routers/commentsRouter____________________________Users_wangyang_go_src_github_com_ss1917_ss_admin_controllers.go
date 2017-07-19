package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:HomeController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:HomeController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:PasswordController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:PasswordController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:PasswordController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:PasswordController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Patch",
			Router: `/`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:RoleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:UserController"],
		beego.ControllerComments{
			Method: "Patch",
			Router: `/`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_admin/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
