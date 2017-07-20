package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

var (
	OperationRecords map[string]*OperationRecord
)

type OperationRecord struct {
	Id            int       `orm:"column(operation_id);auto"`
	Username      string    `orm:"column(username);size(60);null"`
	RegionId      int       `orm:"column(region_id);null"`
	OperationType string    `orm:"column(operation_type);size(20);null"`
	Record        string    `orm:"column(record);size(100);null"`
	Crtime        time.Time `orm:"column(crtime);type(datetime);null"`
}

func (t *OperationRecord) TableName() string {
	return "operation_record"
}

func OperationLog(rid int, username, ot, record string) error {
	o := orm.NewOrm()
	var op OperationRecord
	op.Username = username
	op.RegionId = rid
	op.OperationType = ot
	op.Record = record
	op.Crtime = time.Now()
	_, err := o.Insert(&op)
	return err
}
