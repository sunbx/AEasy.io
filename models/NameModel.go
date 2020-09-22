package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type AeaMiddleNames struct {
	Id           int64   `orm:"column(id)" json:"id"`
	CurrentPrice float64 `orm:"column(current_price)" json:"current_price"`
	EndHeight    int     `orm:"column(end_height)" json:"end_height"`
	Length       int     `orm:"column(length)" json:"length"`
	Name         string  `orm:"column(name)" json:"name"`
	NameID       string  `orm:"column(name_id)" json:"name_id"`
	OverHeight   int     `orm:"column(over_height)" json:"over_height"`
	Owner        string  `orm:"column(owner)" json:"owner"`
	Price        float64 `orm:"column(price)" json:"price"`
	StartHeight  int     `orm:"column(start_height)" json:"start_height"`
	ThHash       string  `orm:"column(th_hash)" json:"th_hash"`
}

// TableName sets the insert table name for this struct type
func (a *AeaMiddleNames) TableName() string {
	return "aea_middle_names"
}

func InsertName(blockHeight int64, names AeaMiddleNames) {
	middleNames, e := FindNameName(names.Name)
	if e == nil {
		if middleNames.CurrentPrice < names.CurrentPrice || int(blockHeight) > middleNames.EndHeight {
			_, _ = orm.NewOrm().InsertOrUpdate(&names)
		}
		return
	}
	_, _ = orm.NewOrm().InsertOrUpdate(&names)

}

//拍卖中 - 即将拍卖结束
func FindNameIdIsNull(height int) ([]AeaMiddleNames, error) {
	var aeaMiddleNames []AeaMiddleNames
	o := orm.NewOrm()
	//_, err := o.Raw("SELECT * FROM `aea_middle_names` where   ? > end_height", height).QueryRows(&aeaMiddleNames)
	_, err := o.Raw("SELECT * FROM `aea_middle_names` where name_id = '' and ? > end_height", height).QueryRows(&aeaMiddleNames)
	return aeaMiddleNames, err
}

//拍卖中 - 即将拍卖结束
func FindNameAuctionOver(page int, height int) ([]AeaMiddleNames, error) {
	var aeaMiddleNames []AeaMiddleNames
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM `aea_middle_names` where name_id = '' and end_height>? order by end_height limit ?,?", height, (page-1)*20, 20).QueryRows(&aeaMiddleNames)
	return aeaMiddleNames, err
}

//拍卖中 - 价格最贵的域名
func FindNameAuctionPrice(page int, height int) ([]AeaMiddleNames, error) {
	var aeaMiddleNames []AeaMiddleNames
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM `aea_middle_names` where name_id = '' and end_height>? order by current_price desc limit ?,?", height, (page-1)*20, 20).QueryRows(&aeaMiddleNames)
	return aeaMiddleNames, err
}

//即将过期 未续费的域名
func FindNameOver(page int, height int) ([]AeaMiddleNames, error) {
	var aeaMiddleNames []AeaMiddleNames
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM `aea_middle_names` where over_height>? order by over_height  limit ?,?", height, (page-1)*20, 20).QueryRows(&aeaMiddleNames)
	return aeaMiddleNames, err
}

//我的 - 已注册的域名
func FindNameMyRegister(address string, page int, height int) ([]AeaMiddleNames, error) {
	var aeaMiddleNames []AeaMiddleNames
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM `aea_middle_names` where owner=? and end_height<? and over_height>? order by over_height  limit ?,?", address, height,height, (page-1)*20, 20).QueryRows(&aeaMiddleNames)
	return aeaMiddleNames, err
}

//我的 - 注册中的域名
func FindNameMyRegisterIng(address string, page int, height int) ([]AeaMiddleNames, error) {
	var aeaMiddleNames []AeaMiddleNames
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM `aea_middle_names` where owner=? and end_height>? order by end_height  limit ?,?", address, height, (page-1)*20, 20).QueryRows(&aeaMiddleNames)
	return aeaMiddleNames, err
}

type NameBase struct {
	Count    int       `json:"count"`
	Sum      int       `json:"sum"`
	SumPrice float64   `json:"sum_price"`
	Ranking  []Ranking `json:"ranking"`
}
type Ranking struct {
	Owner    string  `json:"owner"`
	NameNum  int     `json:"name_num"`
	SumPrice float64 `json:"sum_price"`
}

//域名基础数据
func FindNameBase() (NameBase, error) {
	var data NameBase
	var ranking []Ranking
	o := orm.NewOrm()
	err := o.Raw("select count(distinct(owner)) as count from aea_middle_names").QueryRow(&data)
	fmt.Println("data=>", data)
	err = o.Raw("SELECT SUM(current_price/1000000000000000000) as sum_price FROM aea_middle_names").QueryRow(&data)
	fmt.Println("data=>", data)
	err = o.Raw("SELECT COUNT(*) as sum FROM aea_middle_names").QueryRow(&data)
	fmt.Println("data=>", data)
	_, err = o.Raw("SELECT `owner`,count(OWNER) AS name_num,SUM(current_price/1000000000000000000) AS sum_price FROM aea_middle_names GROUP BY `owner` ORDER BY sum_price DESC LIMIT 20").QueryRows(&ranking)
	fmt.Println("data=>", data)
	data.Ranking = ranking
	return data, err
}

//通过id 获取信息
func FindNameId(nameId string) (AeaMiddleNames, error) {
	var aeaMiddleNames []AeaMiddleNames
	o := orm.NewOrm()
	_, err := o.Raw("select * from aea_middle_names where name_id = ? limit 1", nameId).QueryRows(&aeaMiddleNames)
	if len(aeaMiddleNames) > 0 {
		return aeaMiddleNames[0], err
	} else {
		return AeaMiddleNames{}, err
	}

}

//通过name 获取信息
func FindNameName(name string) (AeaMiddleNames, error) {
	var aeaMiddleNames []AeaMiddleNames
	o := orm.NewOrm()
	_, err := o.Raw("select * from aea_middle_names where name = ? limit 1", name).QueryRows(&aeaMiddleNames)
	if len(aeaMiddleNames) > 0 {
		return aeaMiddleNames[0], err
	} else {
		return AeaMiddleNames{}, err
	}

}

func UpdateNameOwner(nameId string, owner string) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("aea_middle_names").Filter("name_id", nameId).Update(orm.Params{
		"owner": owner,
	})
	return err
}

func UpdateNameOwnerAndIdAndTTL(name string, nameId string, owner string, overHeight int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("aea_middle_names").Filter("name", name).Update(orm.Params{
		"owner":       owner,
		"name_id":     nameId,
		"over_height": overHeight,
	})
	return err
}

func UpdateNameAndIdAndTTL(name string, nameId string, overHeight int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("aea_middle_names").Filter("name", name).Update(orm.Params{
		"name_id":     nameId,
		"over_height": overHeight,
	})
	return err
}

func UpdateNameHeight(nameId string, overHeight int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("aea_middle_names").Filter("name_id", nameId).Update(orm.Params{
		"over_height": overHeight,
	})
	return err
}

func RegisterAeaMiddleNamesDB() {
	orm.RegisterModel(new(AeaMiddleNames))
}
