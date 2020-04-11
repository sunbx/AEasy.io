package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type Tokens struct {
	ID        int64  `orm:"column(id)" json:"id"`
	UserId    int    `orm:"column(user_id)" json:"user_id"`
	Name      string `orm:"column(name)" json:"name"`
	Decimals  int64  `orm:"column(decimals)" json:"decimals"`
	Contracts string `orm:"column(contracts)" json:"contracts"`
	Address   string `orm:"column(address)" json:"address"`
	Count     string `orm:"column(count)" json:"-"`
	CTime     int64  `orm:"column(create_time)" json:"create_time"`
}

func (email *Tokens) TableName() string {
	return "aea_tokens"
}

func RegisterTokenDB() {
	orm.RegisterModel(new(Tokens))
}

func InsertToken(userId int, name string, contracts string, address string, count string) (Tokens, error) {
	unix := time.Now().UnixNano() / 1e6

	tokens := Tokens{
		UserId:    userId,
		Name:      strings.ToLower(name),
		Decimals:  18,
		Contracts: contracts,
		Address:   address,
		Count:     count,
		CTime:     unix,
	}
	_, err := orm.NewOrm().Insert(&tokens)
	return tokens, err
}

//通过 userId 获取 token
func FindTokenUserId(userId int) (Tokens, error) {
	var tokens Tokens
	qs := orm.NewOrm().QueryTable("aea_tokens")
	err := qs.
		Filter("user_id", userId).
		One(&tokens)
	return tokens, err
}
