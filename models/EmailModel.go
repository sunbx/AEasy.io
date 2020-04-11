package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)

type Email struct {
	ID         int    `orm:"column(id)" json:"id"`
	Email      string `orm:"column(email)" json:"email"`
	TemplateId int    `orm:"column(template_id)" json:"template_id"`
	Status     int    `orm:"column(status)" json:"status"`
	Data       string `orm:"column(data)" json:"data"`
	Type       int    `orm:"column(type)" json:"send_type"`
	IP         string `orm:"column(ip)" json:"ip"`
	CTime      int64  `orm:"column(create_time)" json:"create_time"`
	UTime      int64  `orm:"column(update_time)" json:"update_time"`
	DTime      int64  `orm:"column(delete_time)" json:"delete_time"`
}

func (email *Email) TableName() string {
	return "aea_email"
}

func InstEmail(mail string, ip string, t int, captcha string) (int64, error) {
	unix := time.Now().UnixNano() / 1e6
	email := Email{
		Email:      mail,
		TemplateId: 1,
		Status:     0,
		Data:       captcha,
		Type:       t,
		IP:         ip,
		CTime:      unix,
		UTime:      unix}
	id, err := orm.NewOrm().Insert(&email)
	return id, err
}

//验证ip 邮箱
func VerifyIpEmail(mail string, ip string) (Email, error) {
	unix := time.Now().UnixNano() / 1e6
	var email Email
	qs := orm.NewOrm().QueryTable("aea_email")
	err := qs.
		Filter("ip", ip).
		Filter("create_time__gt", unix-1000*60).
		One(&email)
	return email, err
}

//验证注册 邮箱
func VerifyEmail(mail string, captcha string,t int) (Email, error) {
	unix := time.Now().UnixNano() / 1e6
	var email Email
	qs := orm.NewOrm().QueryTable("aea_email")
	err := qs.
		Filter("email", mail).
		Filter("data", captcha).
		Filter("type", t).
		Filter("status", 0).
		Filter("create_time__gt", unix-1000*60).
		One(&email)
	if err == nil {
		//设置已使用
		email.Status = 1
		i, err := orm.NewOrm().Update(&email)
		if i > 0 && err == nil {
			return email, err
		} else {
			return email, err
		}
	} else {
		return email, err
	}
}

func RegisterEmailDB() {
	orm.RegisterModel(new(Email))
}
