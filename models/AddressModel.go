package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type AeaMiddleAddress struct {
	Id         int64   `orm:"column(id)"`
	Address    string  `orm:"column(address)" json:"address"`
	Balance    float64 `orm:"column(balance)"`
	BalanceStr string  `orm:"-" json:"balance"`
	UpdateTime int64   `orm:"column(update_time)" json:"time"`
}

// TableName sets the insert table name for this struct type
func (a *AeaMiddleAddress) TableName() string {
	return "aea_middle_address"
}

func InsertAddress(aeaMiddleAddress AeaMiddleAddress) {
	_, e := orm.NewOrm().InsertOrUpdate(&aeaMiddleAddress)
	if e != nil {
		fmt.Println("error->", e.Error())
	}
}

//获取 balance top 500
func FindAddressBalanceTopList() ([]AeaMiddleAddress, error) {
	var aeaMiddleAddress []AeaMiddleAddress
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM `aea_middle_address` order by  balance desc limit 500").QueryRows(&aeaMiddleAddress)
	return aeaMiddleAddress, err
}

func UpdateAddressTime(address string, time int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("aea_middle_address").Filter("address", address).Update(orm.Params{
		"update_time": time,
	})
	return err
}

func RegisterAeaMiddleAddressDB() {
	orm.RegisterModel(new(AeaMiddleAddress))
}
