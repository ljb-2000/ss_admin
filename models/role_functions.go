package models

import (
	"time"
)

type RoleFunction struct {
	Id         int       `orm:"column(role_function_id);auto"`
	RoleId     int       `orm:"column(role_id);null"`
	FunctionId int       `orm:"column(function_id);null"`
	RegionId   int       `orm:"column(region_id);null"`
	Status     int16     `orm:"column(status);null"`
	Uptime     time.Time `orm:"column(uptime);type(datetime);null"`
	Crtime     time.Time `orm:"column(crtime);type(datetime);null"`
}

func (t *RoleFunction) TableName() string {
	return "role_functions"
}
