package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/ss1917/ss_admin/routers"
)

func main() {
	beego.Run()
}
