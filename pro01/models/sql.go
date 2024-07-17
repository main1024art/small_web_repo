package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"pro01/errors"

	"pro01/utils"
)

type User struct {
	Id       int
	Name     string `orm:"size(100)"`
	Email    string `orm:"size(100)"`
	Password string `orm:"size(100)"`
}

func (u *User) ValiPassword(password string) bool {
	return u.Password == utils.Md5Text(password)
}

func (u *User) GetFormUser(user *User) *errors.Error {
	o := orm.NewOrm()
	errs := errors.New()
	user.Password = utils.Md5Text(user.Password)
	if _, err := o.Insert(user); err != nil {
		errs.Add("01", "注册失败")
		return errs
	}
	return nil
}

func GetUserByName(name string) *User {
	o := orm.NewOrm()
	user := &User{Name: name}
	if err := o.Read(user, "Name"); err == nil {
		return user
	}
	return nil
}

func GetUserById(id int) *User {
	o := orm.NewOrm()
	user := &User{Id: id}
	if err := o.Read(user); err != nil {
		fmt.Println(err)
		return nil
	}
	return user
}

func Queryuser() []*User {
	o := orm.NewOrm()
	var users []*User
	queryset := o.QueryTable(&User{})
	queryset.All(&users)
	return users
}
func DeleteUserById(id int) *errors.Error {
	o := orm.NewOrm()
	errs := errors.New()
	user := &User{Id: id}
	if _, err := o.Delete(user); err != nil {
		errs.Add("02", "删除失败")
		return errs
	}
	fmt.Println("222222")
	return nil
}

func ModifyUser(user *User) *errors.Error {
	o := orm.NewOrm()
	errs := errors.New()
	if _, err := o.Update(user); err != nil {
		errs.Add("03", "删除失败")
		return errs
	}
	return nil
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}
