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
)

//引入数据模型
func init() {
	orm.Debug = false

	//注册驱动
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)

	//注册默认数据库
	host := beego.AppConfig.String("db::host")
	port := beego.AppConfig.String("db::port")
	dbname := beego.AppConfig.String("db::databaseName")
	user := beego.AppConfig.String("db::userName")
	pwd := beego.AppConfig.String("db::password")
	dbConnect := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	_ = orm.RegisterDataBase("default", "mysql", dbConnect )

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

	//设置session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 7776000

	//设置缓冲目录
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"

	//设置国际化
	_ = i18n.SetMessage("zh-CN", "conf/locale_zh-CN.ini")
	_ = i18n.SetMessage("en-US", "conf/locale_en-US.ini")
	_ = beego.AddFuncMap("i18n", i18n.Tr)

	task()
	beego.Run()
	task()

}

//当前是否有定时任务的标记
var isTask = true

//定时任务, 从节点抓取数据
func task() {

	tk := toolbox.NewTask("myTask", "0/10 * * * * *", func() error {
		if isTask {
			isTask = false
			fmt.Println("task start")
			SynAeBlock()
			fmt.Println("task end")
			isTask = true
		} else {
			fmt.Println("task wait")
		}

		return nil
	})
	toolbox.AddTask("myTask", tk)
	toolbox.StartTask()
}



