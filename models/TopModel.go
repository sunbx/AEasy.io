package models

import (
	"database/sql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Top500 struct {
	Id         int64   `orm:"column(id)" `
	Ak         string  `orm:"column(ak)"`
	Balance    float64 `orm:"column(balance);digits(30);decimals(0)"`
	Per        float64 `orm:"column(per)" `
	LastUpdate string  `orm:"column(lastupdate)" `
}

type Top500Json struct {
	Id         int64   ` json:"id"`
	Ak         string  ` json:"ak"`
	Balance    string  ` json:"balance"`
	Per        float64 ` json:"per"`
	LastUpdate string  ` json:"lastupdate"`
}

func InsertTop(top []Top500) {
	for i := 0; i < len(top); i++ {
		_, _ = orm.NewOrm().Insert(&top[i])
	}

}

func GetTopAll(page int) ([]Top500, error) {
	var tops []Top500
	qs := orm.NewOrm().QueryTable("aea_top500")
	_, err := qs.Limit(20).Offset((page - 1) * 20).OrderBy("-balance").All(&tops)
	return tops, err
}

func DeleteAllTop() {
	host := beego.AppConfig.String("db::host")
	port := beego.AppConfig.String("db::port")
	dbname := beego.AppConfig.String("db::databaseName")
	user := beego.AppConfig.String("db::userName")
	pwd := beego.AppConfig.String("db::password")
	dbconnect := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	db, _ := sql.Open("mysql", dbconnect);
	////root为数据库用户名，后面为密码，tcp代表tcp协议，test处填写自己的数据库名称
	//o := orm.NewOrm()
	//for i := 0; i < 500; i++ {
	//	_, e := o.Delete(&Top500{Id: int64(i)})
	//	if e != nil {
	//		fmt.Println(e)
	//	}
	//}

	db.Exec("Delete from aea_top500")
	db.Exec("ALTER TABLE aea_top500 auto_increment=1")
}

func (email *Top500) TableName() string {
	return "aea_top500"
}

func RegisterTopDB() {
	orm.RegisterModel(new(Top500))
}
