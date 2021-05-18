package main

import (
	"ae/models"
	_ "ae/routers"
	"ae/utils"
	"encoding/json"
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
	_ = orm.RegisterDataBase("default", "mysql", dbConnect)

	// 注册数据库
	models.RegisterEmailDB()
	models.RegisterUserDB()
	//models.RegisterOrderDB()
	//models.RegisterAccountDB()
	models.RegisterTokenDB()
	//models.RegisterArticleDB()
	//models.RegisterTopDB()
	models.RegisterAeaMiddleBlockDB()
	models.RegisterAeaMiddleMicroBlockDB()
	models.RegisterAeaMiddleNamesDB()
	models.RegisterAeaMiddleAddressDB()
	models.RegisterContractDB()
}

func main() {

	//设置session
	beego.BConfig.WebConfig.Session.SessionOn = true

	//设置缓冲目录
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"

	//设置国际化
	_ = i18n.SetMessage("zh-CN", "conf/locale_zh-CN.ini")
	_ = i18n.SetMessage("en-US", "conf/locale_en-US.ini")
	_ = beego.AddFuncMap("i18n", i18n.Tr)

	taskNodeSyn()
	taskNamesSyn()
	beego.Run()
	taskNodeSyn()

}

//当前是否有定时任务的标记
var isTaskNode = true

//定时任务, 从节点抓取数据
func taskNodeSyn() {
	tk := toolbox.NewTask("AEASY_NODE_SYN", "0/10 * * * * *", func() error {
		if isTaskNode {
			isTaskNode = false
			fmt.Println("taskNodeSyn start")
			SynAeBlock()
			fmt.Println("taskNodeSyn end")
			isTaskNode = true
		} else {
		}

		return nil
	})
	toolbox.AddTask("AEASY_NODE_SYN", tk)
	toolbox.StartTask()
}

//当前是否有定时任务的标记
var isTaskNames = true

//定时任务, 从节点抓取数据
func taskNamesSyn() {
	tk := toolbox.NewTask("AEASY_NAMES_SYN", "0 */10 * * * *", func() error {
		if isTaskNames {
			isTaskNames = false

			names, e := models.GetNameAll()
			if e == nil {
				for i := 0; i < len(names); i++ {
					names, err := models.FindNameId(names[i].NameID)
					if err != nil {
						continue
					}
					var response = utils.Get(models.NodeURL + "/v2/names/" + names.Name)
					var v2Name V2Name
					err = json.Unmarshal([]byte(response), &v2Name)
					if err != nil {
						continue
					}
					if v2Name.Owner == "" {
						continue
					}
					err = models.UpdateNameAndOwner(names.Name, v2Name.Owner, v2Name.ID, int(v2Name.TTL))
					if err == nil{
						fmt.Println("name update success "+names.Name)
					}else{
						fmt.Errorf("name update error "+names.Name)
					}
				}
			}

			isTaskNames = true
		} else {
		}

		return nil
	})
	toolbox.AddTask("AEASY_NAMES_SYN", tk)
	toolbox.StartTask()
}
