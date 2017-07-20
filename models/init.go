package models

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	DbInit()
}

func DbInit() {
	read_db := "read_db"
	rdbHost := beego.AppConfig.String(read_db + "::db_host")
	rdbName := beego.AppConfig.String(read_db + "::db_name")
	rdbPort := beego.AppConfig.String(read_db + "::db_port")
	rdbUser := beego.AppConfig.String(read_db + "::db_user")
	rdbPawd := beego.AppConfig.String(read_db + "::db_pawd")
	if rdbPort == "" {
		rdbPort = "3306"
	}

	write_db := "write_db"
	wdbHost := beego.AppConfig.String(write_db + "::db_host")
	wdbName := beego.AppConfig.String(write_db + "::db_name")
	wdbPort := beego.AppConfig.String(write_db + "::db_port")
	wdbUser := beego.AppConfig.String(write_db + "::db_user")
	wdbPawd := beego.AppConfig.String(write_db + "::db_pawd")
	if wdbPort == "" {
		wdbPort = "3306"
	}
	wdbCharset := beego.AppConfig.String(write_db + "::db_charset")
	if wdbCharset == "" {
		wdbCharset = "utf8"
	}

	user_db := "user_db"
	udbHost := beego.AppConfig.String(user_db + "::db_host")
	udbName := beego.AppConfig.String(user_db + "::db_name")
	udbPort := beego.AppConfig.String(user_db + "::db_port")
	udbUser := beego.AppConfig.String(user_db + "::db_user")
	udbPawd := beego.AppConfig.String(user_db + "::db_pawd")
	if wdbPort == "" {
		wdbPort = "3306"
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.DefaultTimeLoc, _ = time.LoadLocation("Asia/Shanghai")
	conn := wdbUser + ":" + wdbPawd + "@tcp(" + wdbHost + ":" + wdbPort + ")/" + wdbName + "?charset=" + wdbCharset + "&loc=Local"
	orm.RegisterDataBase("default", "mysql", conn)

	conn_read := rdbUser + ":" + rdbPawd + "@tcp(" + rdbHost + ":" + rdbPort + ")/" + rdbName + "?charset=" + wdbCharset + "&loc=Local"
	orm.RegisterDataBase("read", "mysql", conn_read)

	conn_user := udbUser + ":" + udbPawd + "@tcp(" + udbHost + ":" + udbPort + ")/" + udbName + "?charset=" + wdbCharset + "&loc=Local"
	orm.RegisterDataBase("user", "mysql", conn_user, 5, 30)

	orm.RegisterModelWithPrefix("login_", new(User), new(Role), new(UserRole), new(Function), new(RoleFunction))
	//orm.RunSyncdb("user", false, false)
	orm.RegisterModelWithPrefix("mg_", new(OperationRecord))
	//orm.RunSyncdb("default", false, false)
}
