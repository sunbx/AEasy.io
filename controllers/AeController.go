package controllers

import (
	"ae/models"
	"ae/utils"
	"bytes"
	"encoding/json"
	"github.com/beego/i18n"
	"sync"
	"time"
	aemodels "github.com/aeternity/aepp-sdk-go/swagguard/node/models"
)

//获取区块高度
type ApiBlocksTopController struct {
	BaseController
}

//查询ae th
type ApiThHashController struct {
	BaseController
}

//转账
type ApiTransferController struct {
	BaseController
}

//创建账户
type ApiCreateAccountController struct {
	BaseController
}

//钱包转账
type WalletTransferController struct {
	BaseController
}

//钱包转账记录
type WalletTransferRecordController struct {
	BaseController
}

type OracleRegisterController struct {
	BaseController
}

type OracleQueryController struct {
	BaseController
}

type OracleResponseController struct {
	BaseController
}
type OracleListController struct {
	BaseController
}

type OracleQueryDetailController struct {
	BaseController
}


var lock sync.Mutex

//返回区块高度
func (c *ApiBlocksTopController) Post() {
	if c.verifyAppId() {
		height := models.ApiBlocksTop()
		var data = map[string]uint64{}
		data["height"] = height

		c.SuccessJson(data)
	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(),"appId verify error"), JsonData{})
	}
}

//查询th
func (c *ApiThHashController) Post() {
	if c.verifyAppId() {
		th := c.GetString("th")
		if th == "" {
			c.ErrorJson(-200,  i18n.Tr(c.getHeaderLanguage(),"parameter is nul"), JsonData{})
			return
		}
		t := models.ApiThHash(th)
		c.SuccessJson(t)
	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(),"appId verify error"), JsonData{})
	}

}

//数据上链
func (c *ApiTransferController) Post() {
	data := c.GetString("data")
	if c.verifySecret() {
		if len(data) > 5000 || len(data) == 0 {
			c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(),"Len is greater than 50000 or len is equal to 0"), JsonData{})
			return
		}
		account := c.GetSecretAeAccount()
		if !models.Is1AE(account.Address) {
			c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(),"The balance should be greater than 1ae"), JsonData{})
			return
		}
		lock.Lock()
		tx, e := models.ApiSpend(account, "ak_wNL5NYtbr6AAuAWxKGF3ZwQNBeb7UMpu9BHoVb24pS9iWAQCo", 0.001, data)
		time.Sleep(3 * time.Second)
		lock.Unlock()
		if e == nil {
			c.SuccessJson(map[string]interface{}{"tx": tx})
		} else {
			c.ErrorJson(-500, e.Error(), JsonData{})
		}
	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(),"appId or secret verify error"), JsonData{})
	}
}

func (c *WalletTransferRecordController) Post() {
	address := c.GetString("address")
	page, _ := c.GetInt("page", 1)
	if address == "" {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(),"parameter is nul"), JsonData{})
		return
	}
	if c.verifyAppId() {

		blocksDb, err := models.FindMicroBlockBlockList(address, page, "all")

		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}
		var txs []map[string]interface{}
		for i := 0; i < len(blocksDb); i++ {
			mapObj := make(map[string]interface{})

			// body是后端的http返回结果
			d := json.NewDecoder(bytes.NewReader([]byte(blocksDb[i].Tx)))
			d.UseNumber()
			err = d.Decode(&mapObj)
			txs = append(txs, mapObj)
		}
		c.SuccessJson(txs)
	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(),"appId or secret verify error"), JsonData{})
	}
}

func (c *WalletTransferController) Post() {
	data := c.GetString("data")
	address := c.GetString("address")
	signingKey := c.GetString("signingKey")
	amount, _ := c.GetFloat("amount", 0.001)
	if address == "" || signingKey == "" {
		c.ErrorJson(-100,  i18n.Tr(c.getHeaderLanguage(),"parameter is nul"), JsonData{})
		return
	}
	if c.verifyAppId() {
		if len(data) > 5000 {
			c.ErrorJson(-100,  i18n.Tr(c.getHeaderLanguage(),"Len is greater than 50000 or len is equal to 0"), JsonData{})
			return
		}
		account, err := models.SigningKeyHexStringAccount(signingKey)
		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}
		lock.Lock()
		tx, e := models.ApiSpend(account, address, amount, data)
		lock.Unlock()
		if e == nil {
			c.SuccessJson(map[string]interface{}{"tx": tx})
		} else {
			c.ErrorJson(-500, e.Error(), JsonData{})
		}
	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(),"appId or secret verify error"), JsonData{})
	}
}

//创建用户
func (c *ApiCreateAccountController) Post() {
	accountCreate, mnemonic := models.CreateAccount()

	c.SuccessJson(map[string]string{
		"mnemonic":    mnemonic,
		"address":     accountCreate.Address,
		"signing_key": accountCreate.SigningKeyToHexString(),
	})
}


func (c *OracleRegisterController) Get() {
	querySpace := c.GetString("querySpace")
	responseSpec := c.GetString("responseSpec")
	signingKey := c.GetString("signingKey")
	queryFee, _ := c.GetFloat("queryFee", 0)
	if querySpace == "" || responseSpec == "" || signingKey == "" {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(), "parameter is nul"), JsonData{})
		return
	}
	//获取用户
	account, err := models.SigningKeyHexStringAccount(signingKey)
	if err != nil {
		c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "Account signingKey error"), []JsonData{})
		return
	}

	if !models.Is1AE(account.Address) {
		c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "The balance is insufficient please keep the number of ae greater than 1"), JsonData{})
		return
	}

	//注册预言机
	hash, oraclePubKey := models.OracleRegister(account, querySpace, responseSpec, queryFee)
	c.SuccessJson(map[string]interface{}{
		"hash":         hash,
		"oraclePubKey": oraclePubKey})
}

func (c *OracleQueryController) Post() {
	oracleID := c.GetString("oracleID")
	querySpec := c.GetString("querySpec")
	signingKey := c.GetString("signingKey")
	queryFee, _ := c.GetFloat("queryFee", 0)
	if oracleID == "" || querySpec == "" || signingKey == "" {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(), "parameter is nul"), JsonData{})
		return
	}
	account, err := models.SigningKeyHexStringAccount(signingKey)

	if err != nil {
		c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "Account signingKey error"), []JsonData{})
		return
	}

	if !models.Is1AE(account.Address) {
		c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "The balance is insufficient please keep the number of ae greater than 1"), JsonData{})
		return
	}

	//这可能要判断一下
	hash, opId := models.OracleQuery(account, oracleID, querySpec, queryFee)

	c.SuccessJson(map[string]interface{}{
		"hash": hash,
		"oqID": opId})

}
func (c *OracleResponseController) Get() {

	oracleID := c.GetString("oracleID")
	oqID := c.GetString("oqID")
	signingKey := c.GetString("signingKey")
	response := c.GetString("response")
	if oracleID == "" || oqID == "" || signingKey == "" || response == "" {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(), "parameter is nul"), JsonData{})
		return
	}

	account, err := models.SigningKeyHexStringAccount(signingKey)

	if err != nil {
		c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "Account signingKey error"), []JsonData{})
		return
	}

	if !models.Is1AE(account.Address) {
		c.ErrorJson(-500, i18n.Tr(c.getHeaderLanguage(), "The balance is insufficient please keep the number of ae greater than 1"), JsonData{})
		return
	}

	hash := models.OracleResponse(account, oracleID, oqID, response)

	c.SuccessJson(map[string]interface{}{
		"hash": hash,})
}



func (c *OracleListController) Get() {
	oracleID := c.GetString("oracleID")
	t := c.GetString("type")
	if t != "open" && t != "closed" {
		t = "open"
	}

	if oracleID == "" {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(), "parameter is nul"), JsonData{})
		return
	}

	response := utils.Get(models.NodeURL + "/v2/oracles/" + oracleID + "/queries?type=" + t)
	var oracleQueries aemodels.OracleQueries
	err := json.Unmarshal([]byte(response), &oracleQueries)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.SuccessJson(oracleQueries)
}

func (c *OracleQueryDetailController) Get() {
	oracleID := c.GetString("oracleID")
	oqID := c.GetString("oqID")

	if oracleID == "" {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(), "parameter is nul"), JsonData{})
		return
	}

	response := utils.Get(models.NodeURL + "/v2/oracles/" + oracleID + "/queries/" + oqID)
	var oracleQuery aemodels.OracleQuery
	err := json.Unmarshal([]byte(response), &oracleQuery)
	if err != nil {
		c.ErrorJson(-500, err.Error(), JsonData{})
		return
	}
	c.SuccessJson(oracleQuery)
}
