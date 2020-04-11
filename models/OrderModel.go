package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	tsgutils "github.com/typa01/go-utils"
	"time"
)

type Order struct {
	ID             int64   `orm:"column(id)" json:"id"`
	OrderNo        string  `orm:"column(order_no)" json:"-"`
	Body           string  `orm:"column(body)" json:"-"`
	Data           string  `orm:"column(data)" json:"-"`
	Tokens         float64 `orm:"column(tokens);digits(30);decimals(0)" json:"-"`
	PayTime        int64   `orm:"column(pay_time)" json:"-"`
	PayStatus      int     `orm:"column(pay_status)" json:"-"`
	AppId          string  `orm:"column(app_id)" json:"-"`
	OpenId         string  `orm:"column(open_id)" json:"-"`
	SendAddress    string  `orm:"column(send_address)" json:"-"`
	ReceiveAddress string  `orm:"column(receive_address)" json:"-"`
	ThPlatform     string  `orm:"column(th_platform)" json:"-"`
	ThMerchants    string  `orm:"column(th_merchants)" json:"-"`
	UpdateTime     int64   `orm:"column(update_time)" json:"update_time"`
	CreateTime     int64   `orm:"column(create_time)" json:"create_time"`
}

func (email *Order) TableName() string {
	return "aea_order"
}

func RegisterOrderDB() {
	orm.RegisterModel(new(Order))
}

func InsertOrder(body string, data string, tokens float64, appId string, openId string, sendAddress string, receiveAddress string) (Order, error) {
	unix := time.Now().UnixNano() / 1e6
	order := Order{
		OrderNo:        "ae_" + tsgutils.GUID(),
		Body:           body,
		Data:           data,
		Tokens:         tokens,
		PayTime:        0,
		PayStatus:      1,
		AppId:          appId,
		OpenId:         openId,
		SendAddress:    sendAddress,
		ReceiveAddress: receiveAddress,
		CreateTime:     unix,
	}
	_, err := orm.NewOrm().Insert(&order)
	return order, err
}

//通过 用户order_id 获取 order
func FindOrderOrderNo(orderNo string) (Order, error) {
	var order Order
	qs := orm.NewOrm().QueryTable("aea_order")
	err := qs.
		Filter("order_no", orderNo).
		One(&order)
	return order, err
}

//通过 用户order_id 获取 order
func UpdateOrderOrderNo(orderNo string, thPlatform string, thMerchants string) {
	unix := time.Now().UnixNano() / 1e6
	qs := orm.NewOrm().QueryTable("aea_order")
	_, _ = qs.
		Filter("order_no", orderNo).Update(orm.Params{
		"th_platform":  thPlatform,
		"th_merchants": thMerchants,
		"pay_time":     unix,
		"pay_status":   2,
		"update_time":  unix,
	})
}
