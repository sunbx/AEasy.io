package models

import (
	"ae/utils"
	"github.com/astaxie/beego/orm"
	"time"
)

type AeasyAccount struct {
	ID         int64  `orm:"column(id)" json:"id"`
	UserId     int64  `orm:"column(user_id)" json:"user_id"`
	Status     int    `orm:"column(status)" json:"status"`
	Address    string `orm:"column(address)" json:"-"`
	Password   string `orm:"column(password)" json:"-"`
	SigningKey string `orm:"column(signing_key)" json:"-"`
	CTime      int64  `orm:"column(create_time)" json:"create_time"`
}

func (email *AeasyAccount) TableName() string {
	return "aea_account"
}

func RegisterAccountDB() {
	orm.RegisterModel(new(AeasyAccount))
}

func InsertAccount(userId int64, signingKey string, address string) (AeasyAccount, error) {
	unix := time.Now().UnixNano() / 1e6

	aeasyAccount := AeasyAccount{
		UserId:     userId,
		Address:    address,
		SigningKey: signingKey,
		CTime:      unix,
	}
	_, err := orm.NewOrm().Insert(&aeasyAccount)
	return aeasyAccount, err
}

//通过 用户id 获取 ae_account
func FindAccountSigningKey(signingKey string) (AeasyAccount, error) {
	var aesyAccount AeasyAccount
	qs := orm.NewOrm().QueryTable("aea_account")
	err := qs.
		Filter("signing_key", signingKey).
		One(&aesyAccount)
	return aesyAccount, err
}

//通过 用户openId 获取 ae_account
func FindAccountOpenId(openId string) (AeasyAccount, error) {
	var aesyAccount AeasyAccount
	qs := orm.NewOrm().QueryTable("aea_account")
	err := qs.
		Filter("open_id", openId).
		One(&aesyAccount)
	return aesyAccount, err
}

//通过 用户email 获取 ae_account
func FindAccountEmail(email string) (AeasyAccount, error) {
	var aesyAccount AeasyAccount
	qs := orm.NewOrm().QueryTable("aea_account")
	err := qs.
		Filter("email", email).
		One(&aesyAccount)
	return aesyAccount, err
}

//更新token数量信息
func UpdateAccountOpenIdToToken(address string, token float64) {
	qs := orm.NewOrm().QueryTable("aea_account")
	_, _ = qs.
		Filter("address", address).Update(orm.Params{
		"tokens": token,
	})
}

//更新token数量信息
func UpdateAccountOpenIdToEmailPassword(openId string, email string, password string) {
	qs := orm.NewOrm().QueryTable("aea_account")
	_, _ = qs.
		Filter("open_id", openId).Update(orm.Params{
		"email":    email,
		"password": utils.Md5V(password),
	})
}
