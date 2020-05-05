package main

import (
	"ae/models"
	_ "ae/routers"
	"ae/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/shopspring/decimal"
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

//拽出来 aens names
func namesData() {
	blocks, e := models.FindMicroBlockBlockNames()
	if e == nil {
		for i := 0; i < len(blocks); i++ {
			block, _ := models.FindMicroBlockBlockNameorData(blocks[i].Name)
			response := utils.Get("http://node.aechina.io:3013/v2/names/" + blocks[i].Name)
			var v2Name V2Name
			err := json.Unmarshal([]byte(response), &v2Name)
			if err != nil {
				fmt.Println(500, err.Error())
				return
			}

			var endHeight int
			if len(blocks[i].Name)-6 <= 4 {
				endHeight = int(block.BlockHeight + 29760)
			} else if len(blocks[i].Name)-6 >= 5 && len(blocks[i].Name)-6 <= 8 {
				endHeight = int(block.BlockHeight + 14880)
			} else if len(blocks[i].Name)-6 >= 9 && len(blocks[i].Name)-6 <= 12 {
				endHeight = int(block.BlockHeight + 480)
			} else {
				endHeight = int(block.BlockHeight + 0)
			}
			if v2Name.TTL == 0 {
				v2Name.TTL = int64(endHeight + 50000)
			}

			var price string
			if len(blocks[i].Name)-6 == 1 {
				price = "570288700000000000000"
			} else if len(blocks[i].Name)-6 == 2 {
				price = "352457800000000000000"
			} else if len(blocks[i].Name)-6 == 3 {
				price = "217830900000000000000"
			} else if len(blocks[i].Name)-6 == 4 {
				price = "134626900000000000000"
			} else if len(blocks[i].Name)-6 == 5 {
				price = "83204000000000000000"
			} else if len(blocks[i].Name)-6 == 6 {
				price = "51422900000000000000"
			} else if len(blocks[i].Name)-6 == 7 {
				price = "31781100000000000000"
			} else if len(blocks[i].Name)-6 == 8 {
				price = "19641800000000000000"
			} else if len(blocks[i].Name)-6 == 9 {
				price = "12139300000000000000"
			} else if len(blocks[i].Name)-6 == 10 {
				price = "7502500000000000000"
			} else if len(blocks[i].Name)-6 == 11 {
				price = "4636800000000000000"
			} else if len(blocks[i].Name)-6 == 12 {
				price = "2865700000000000000"
			} else if len(blocks[i].Name)-6 >= 13 {
				price = "2865700000000000000"
			}
			priceFloat, _ := strconv.ParseFloat(price, 64)
			priceFormat := utils.FormatTokensP(priceFloat, 4)

			mapObj := make(map[string]interface{})

			// body是后端的http返回结果
			d := json.NewDecoder(bytes.NewReader([]byte(block.Tx)))
			d.UseNumber()
			err = d.Decode(&mapObj)
			if err != nil {
				// 错误处理
				fmt.Println("Decode", "error.")
			}
			fmt.Println(mapObj)
			f, _ := mapObj["name_fee"].(json.Number).Float64()
			decimalNum := decimal.NewFromFloat(f)
			priceFloat2, _ := strconv.ParseFloat(decimalNum.String(), 64)
			priceFormat2 := utils.FormatTokensP(priceFloat2, 4)
			fmt.Println("name->" + blocks[i].Name)
			fmt.Println("createHeight->" + strconv.Itoa(int(block.BlockHeight)))
			fmt.Println("endHeight->" + strconv.Itoa(endHeight))
			fmt.Println("orderHeight->" + strconv.Itoa(int(v2Name.TTL)))
			fmt.Println("th->" + block.Hash)
			fmt.Println("len->" + strconv.Itoa(len(blocks[i].Name)-6))
			fmt.Println("name_id->" + v2Name.ID)
			fmt.Println("owner->" + block.AccountId)
			fmt.Println("price->" + priceFormat)
			fmt.Println("priceCurrent->" + priceFormat2)
			//fmt.Println(priceFormat2)
			fmt.Println("=============================")

			var aeaMiddleName models.AeaMiddleNames
			aeaMiddleName.Name = blocks[i].Name
			aeaMiddleName.StartHeight = int(block.BlockHeight)
			aeaMiddleName.EndHeight = endHeight
			aeaMiddleName.OverHeight = int(v2Name.TTL)
			aeaMiddleName.ThHash = block.Hash
			aeaMiddleName.Length = len(blocks[i].Name) - 6
			aeaMiddleName.NameID = v2Name.ID
			aeaMiddleName.Owner = block.AccountId
			aeaMiddleName.Price = priceFloat
			//f2, _ := decimalNum.Float64()

			aeaMiddleName.CurrentPrice = priceFloat2

			models.InsertName(aeaMiddleName)

		}
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
