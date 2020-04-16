package models

import (
	"ae/utils"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)

type User struct {
	ID            int    `orm:"column(id)" json:"id"`
	Email         string `orm:"column(email)" json:"email"`
	Password      string `orm:"column(password)" json:"-"`
	Nickname      string `orm:"column(nickname)" json:"nickname"`
	Gender        int    `orm:"column(gender)" json:"-"`
	Status        int    `orm:"column(status)" json:"status"`
	LoginIp       string `orm:"column(login_ip)" json:"-" `
	LoginLastTime int64  `orm:"column(login_last_time)" json:"-"`
	CTime         int64  `orm:"column(create_time)" json:"-"`
	UTime         int64  `orm:"column(update_time)" json:"-"`
	DTime         int64  `orm:"column(delete_time)" json:"-"`
}

func (user *User) TableName() string {
	return "aea_user"
}

type Secret struct {
	ID         int     `orm:"column(id)" json:"id"`
	UserId     int64   `orm:"column(user_id)" json:"user_id"`
	AppId      string  `orm:"column(app_id)" json:"app_id"`
	AppSecret  string  `orm:"column(app_secret)" json:"app_secret"`
	Contracts  string  `orm:"column(contracts)" json:"contracts"`
	Type       int     `orm:"column(type)" json:"type"`
	Status     int     `orm:"column(status)" json:"status"`
	Mnemonic   string  `orm:"column(mnemonic)" json:"-"`
	Address    string  `orm:"column(address)" json:"-"`
	SigningKey string  `orm:"column(signing_key)" json:"-"`
	Tokens     float64 `orm:"column(tokens);digits(30);decimals(0)" json:"tokens"`
	IsShow     int     `orm:"column(is_show);digits(30);decimals(0)" json:"-"`
	CTime      int64   `orm:"column(create_time)" json:"create_time"`
}

func (secret *Secret) TableName() string {
	return "aea_secret"
}

//InstUser 插入用户
func InstUser(mail string, password string, addr string) (int64, error) {
	unix := time.Now().UnixNano() / 1e6
	user := User{
		Email:         mail,
		Password:      utils.Md5V(password),
		Nickname:      mail,
		Gender:        0,
		Status:        0,
		LoginIp:       addr,
		LoginLastTime: unix,
		CTime:         unix,
		UTime:         unix}
	id, err := orm.NewOrm().Insert(&user)

	if err == nil {
		appId := utils.UniqueId(string(id))
		appSecret := utils.Md5V(appId + string(id))
		mnemonic, signingKey, address := CreateAccountUtils()
		secret := Secret{
			UserId:     id,
			AppId:      appId,
			Mnemonic:   mnemonic,
			AppSecret:  appSecret,
			Address:    address,
			SigningKey: signingKey,
			CTime:      unix,
		}
		_, err = orm.NewOrm().Insert(&secret)
		return id, err
	}
	return id, err
}

//插入商户
func InsertSecret(id int64) (int64, error) {
	unix := time.Now().UnixNano() / 1e6
	appId := utils.UniqueId(string(id))
	secret := Secret{
		UserId:    id,
		AppId:     appId,
		AppSecret: utils.Md5V(appId + string(id)),
		CTime:     unix,
	}
	_, err := orm.NewOrm().Insert(&secret)
	return id, err
}

//通过email获取用户信息
func FindUserEmail(mail string) (User, error) {
	var user User
	qs := orm.NewOrm().QueryTable("aea_user")
	err := qs.
		Filter("email", mail).
		One(&user)
	return user, err
}

//通过id 获取用户信息
func FindUserID(id int64) (User, error) {
	var user User
	qs := orm.NewOrm().QueryTable("aea_user")
	err := qs.
		Filter("id", id).
		One(&user)
	return user, err
}

//通过账号密码 获取用户信息
func FindUserEmailPassword(mail string, password string) (User, error) {
	var user User
	qs := orm.NewOrm().QueryTable("aea_user")
	err := qs.
		Filter("email", mail).
		Filter("password", utils.Md5V(password)).
		One(&user)
	return user, err
}

//通过 用户id 获取商户信息
func FindSecretUserID(id int) (Secret, error) {
	var secret Secret
	qs := orm.NewOrm().QueryTable("aea_secret")
	err := qs.
		Filter("user_id", id).
		One(&secret)
	return secret, err
}

//通过 appId 和 appSecret获取商户信息
func FindSecretIdSecret(appId string, appSecret string) (Secret, error) {
	var secret Secret
	qs := orm.NewOrm().QueryTable("aea_secret")
	err := qs.
		Filter("app_id", appId).
		Filter("app_secret", appSecret).
		Filter("is_show", 1).
		Filter("status", 0).
		One(&secret)
	return secret, err
}

//通过 appId获取商户信息
func FindSecretId(appId string) (Secret, error) {
	var secret Secret
	qs := orm.NewOrm().QueryTable("aea_secret")
	err := qs.
		Filter("app_id", appId).
		Filter("is_show", 1).
		Filter("status", 0).
		One(&secret)
	return secret, err
}

//更新token数量信息
func UpdateSecretTokens(address string, token float64, isShow int) {
	qs := orm.NewOrm().QueryTable("aea_secret")
	_, _ = qs.
		Filter("address", address).Update(orm.Params{
		"tokens":  token,
		"is_show": isShow,
	})
}

//更新contracts 合约地址
func UpdateSecretContracts(userId int, contracts string) {
	qs := orm.NewOrm().QueryTable("aea_secret")
	_, _ = qs.
		Filter("user_id", userId).Update(orm.Params{
		"contracts":  contracts,
	})
}
func RegisterUserDB() {
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Secret))
}
