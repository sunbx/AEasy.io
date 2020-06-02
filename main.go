package main

import (
	"ae/models"
	_ "ae/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	"github.com/beego/i18n"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"strconv"
	"strings"
)

//引入数据模型
func init() {
	orm.Debug = false


	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	host := beego.AppConfig.String("db::host")
	port := beego.AppConfig.String("db::port")
	dbname := beego.AppConfig.String("db::databaseName")
	user := beego.AppConfig.String("db::userName")
	pwd := beego.AppConfig.String("db::password")
	dbconnect := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	_ = orm.RegisterDataBase("default", "mysql", dbconnect /*"root:root@tcp(localhost:3306)/test?charset=utf8"*/) //密码为空格式

	// 注册数据库
	models.RegisterEmailDB()
	models.RegisterUserDB()
	models.RegisterOrderDB()
	models.RegisterAccountDB()
	models.RegisterTokenDB()
	models.RegisterArticleDB()
	models.RegisterTopDB()

	models.RegisterAeaMiddleBlockDB()
	models.RegisterAeaMiddleMicroBlockDB()
	models.RegisterAeaMiddleNamesDB()
	models.RegisterAeaMiddleAddressDB()

}

func main() {

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 7776000

	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	i18n.SetMessage("zh-CN", "conf/locale_zh-CN.ini")
	i18n.SetMessage("en-US", "conf/locale_en-US.ini")
	beego.AddFuncMap("i18n", i18n.Tr)

	task()
	beego.Run()
	task()

}

var isTask = true

func task() {
	tk := toolbox.NewTask("myTask", "0/10 * * * * *", func() error {
		if (isTask) {
			isTask = false
			fmt.Println("定时任务执行")
			SynAeBlock()
			fmt.Println("定时任务结束")
			isTask = true
		} else {
			fmt.Println("等等吧,有其他任务正在执行=========================等等吧,有其他任务正在执行=========================")
		}

		return nil
	})
	//funcName()
	//err := tk.Run()
	//	//if err != nil {
	//	//	fmt.Println(err)
	//	//}
	toolbox.AddTask("myTask", tk)
	toolbox.StartTask()
	fmt.Println("定时任务开启")




}


func funcUpdateAddressTime() {
	var aeaMiddleAddress []models.AeaMiddleAddress
	o := orm.NewOrm()
	_, err := o.Raw("select * FROM `aea_middle_address` where update_time=0").QueryRows(&aeaMiddleAddress)
	if err != nil {
		fmt.Println(err.Error())
	}
	for i := 0; i < len(aeaMiddleAddress); i++ {
		if aeaMiddleAddress[i].UpdateTime == 0 {
			var aeaMiddleMicroBlock []models.AeaMiddleMicroBlock
			_, err := o.Raw("SELECT *  FROM `aea_middle_micro_block` where recipient_id = '" + aeaMiddleAddress[i].Address + "' or sender_id ='" + aeaMiddleAddress[i].Address + "' order by time desc limit 1").QueryRows(&aeaMiddleMicroBlock)
			if err != nil {
				fmt.Println("aea_middle_micro_block", err.Error())
				return
			}

			err = models.UpdateAddressTime(aeaMiddleAddress[i].Address, aeaMiddleMicroBlock[0].Time)
			if err != nil {
				fmt.Println("UpdateAddressTime", err.Error())
				return
			}
			fmt.Println("sucess ->", aeaMiddleAddress[i].Address)

		} else {
			fmt.Println("time no 0 ->", aeaMiddleAddress[i].Address)
		}

	}
}

func funcAddress() {
	var aeaMiddleMicroBlock []models.AeaMiddleMicroBlock
	o := orm.NewOrm()
	_, err := o.Raw("SELECT distinct recipient_id FROM `aea_middle_micro_block` WHERE `recipient_id`  !='null' and type = 'SpendTx'").QueryRows(&aeaMiddleMicroBlock)
	if err != nil {
		fmt.Println(err.Error())
	}
	for i := 0; i < len(aeaMiddleMicroBlock); i++ {
		if strings.Contains(aeaMiddleMicroBlock[i].RecipientId, "ak_") {
			account, e := models.ApiGetAccount(aeaMiddleMicroBlock[i].RecipientId)
			if e != nil {
				fmt.Println(e.Error())
				return
			}
			var address models.AeaMiddleAddress
			address.Address = aeaMiddleMicroBlock[i].RecipientId

			tokens, e := strconv.ParseFloat(account.Balance.String(), 64)
			address.Balance = tokens
			models.InsertAddress(address)
			fmt.Println("sucess->", aeaMiddleMicroBlock[i].RecipientId, "->", address.Balance)
		} else {
			fmt.Println("error->", aeaMiddleMicroBlock[i].RecipientId)
		}

	}
}

func funcName() {
	var aeaMiddleMicroBlock []models.AeaMiddleMicroBlock
	o := orm.NewOrm()
	_, err := o.Raw("select * from aea_middle_micro_block  where  type='NameTransferTx' order by time desc").QueryRows(&aeaMiddleMicroBlock)
	if err != nil {
		fmt.Println(err.Error())
	}
	for i := 0; i < len(aeaMiddleMicroBlock); i++ {
		if aeaMiddleMicroBlock[i].NameId != "" {
			err := models.UpdateNameOwner(aeaMiddleMicroBlock[i].NameId, aeaMiddleMicroBlock[i].RecipientId)
			if err != nil {
				fmt.Println(err.Error())
			}

		}

		fmt.Println("sucess->", aeaMiddleMicroBlock[i].NameId, "---", aeaMiddleMicroBlock[i].RecipientId)
	}
}



type Foo struct {
	AccountID string `json:"account_id"`
	Fee       int64  `json:"fee"`
	Name      string `json:"name"`
	NameFee   int64  `json:"name_fee"`
	NameSalt  int64  `json:"name_salt"`
	Nonce     int64  `json:"nonce"`
	Type      string `json:"type"`
	Version   int64  `json:"version"`
}

type V2Name struct {
	ID       string     `json:"id"`
	Pointers []Pointers `json:"pointers"`
	TTL      int64      `json:"ttl"`
}

type Pointers struct {
	ID  string `json:"id"`
	Key string `json:"key"`
}
