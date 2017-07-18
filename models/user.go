package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

var (
	Users map[string]*User
)

type User struct {
	UserId     int       `orm:"column(user_id);auto"`
	RegionId   int       `orm:"column(region_id);null"`
	Username   string    `orm:"column(username);size(120)" json:"name`
	Password   string    `orm:"column(password);size(200)"`
	Email      string    `orm:"column(email);size(200)"`
	Tel        string    `orm:"column(tel);size(11)"`
	Status     int16     `orm:"column(status);null"`
	Superuser  string    `orm:"column(superuser);size(2);"`
	Nickname   string    `orm:"column(nickname);size(64);"`
	Sex        string    `orm:"column(sex);size(10);null"`
	LastIp     string    `orm:"column(lastip);size(20);null"`
	Last_login time.Time `orm:"column(last_login);type(datetime);null"`
	Wtime      time.Time `orm:"column(wtime);type(datetime);null"`
}

func (t *User) TableName() string {
	return "user"
}

func UserGetById(id int) (*User, error) {
	u := new(User)
	o := orm.NewOrm()
	o.Using("user")

	err := o.QueryTable(u).Filter("UserId", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UserGetByName(username string) (*User, error) {
	u := new(User)
	o := orm.NewOrm()
	o.Using("user")

	err := o.QueryTable(u).Filter("Username", username).One(u)

	if err != nil {
		return nil, err
	}
	return u, nil
}

func UserUpdate(user *User, fields ...string) error {
	o := orm.NewOrm()
	o.Using("user")
	_, err := o.Update(user, fields...)
	return err
}

func UserGetAll() (maps []orm.Params, err error) {
	o := orm.NewOrm()
	o.Using("user")
	u := new(User)
	if _, err = o.QueryTable(u).Exclude("Status", -20).Values(&maps, "UserId", "Username", "Nickname", "Email", "Tel", "Status", "Last_login"); err != nil {
		return
	}
	return
}

//通过ID删除
func DeleteUser(Id int) (err error) {
	o := orm.NewOrm()
	o.Using("user")
	u := new(User)
	_, err = o.QueryTable(u).Filter("UserId", Id).Delete()
	return
}
