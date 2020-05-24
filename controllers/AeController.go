package controllers

import (
	"ae/models"
	"bytes"
	"encoding/json"
	"fmt"
	"time"
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

//返回区块高度
func (c *ApiBlocksTopController) Get() {
	if c.verifyAppId() {
		height := models.ApiBlocksTop()
		var data = map[string]uint64{}
		data["height"] = height

		c.SuccessJson(data)
	} else {
		c.ErrorJson(-100, "appId verify error", JsonData{})
	}
}

//查询th
func (c *ApiThHashController) Get() {
	if c.verifyAppId() {
		th := c.GetString("th")
		if th == "" {
			c.ErrorJson(-200, "th is nil", JsonData{})
			return
		}
		t := models.ApiThHash(th)
		c.SuccessJson(t)
	} else {
		c.ErrorJson(-100, "appId verify error", JsonData{})
	}

}

//数据上链
func (c *ApiTransferController) Post() {
	data := c.GetString("data")
	if c.verifySecret() {
		if len(data) > 5000 || len(data) == 0 {
			c.ErrorJson(-100, "data len > 50000 or data len = 0", JsonData{})
			return
		}
		account := c.GetSecretAeAccount()
		if !models.Is1AE(account.Address) {
			c.ErrorJson(-500, "The balance should be greater than 1ae", JsonData{})
			return
		}
		tx, e := models.ApiSpend(account, "ak_wNL5NYtbr6AAuAWxKGF3ZwQNBeb7UMpu9BHoVb24pS9iWAQCo", 0.001, data)
		time.Sleep(3 * time.Second)
		if e == nil {
			c.SuccessJson(map[string]interface{}{"tx": tx})
		} else {
			c.ErrorJson(-500, e.Error(), JsonData{})
		}
	} else {
		c.ErrorJson(-100, "appId or secret verify error", JsonData{})
	}
}

func (c *WalletTransferRecordController) Post() {
	address := c.GetString("address")
	page, _ := c.GetInt("page", 1)
	if address == "" {
		c.ErrorJson(-100, "address = nil", JsonData{})
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
			if err != nil {
				// 错误处理
				fmt.Println("Decode", "error.")
			}
			txs = append(txs, mapObj)
		}
		c.SuccessJson(txs)
	} else {
		c.ErrorJson(-100, "appId or secret verify error", JsonData{})
	}
}

func (c *WalletTransferController) Post() {
	data := c.GetString("data")
	address := c.GetString("address")
	signingKey := c.GetString("signingKey")
	amount, _ := c.GetFloat("amount", 0.001)
	if address == "" || signingKey == "" {
		c.ErrorJson(-100, "address = nil", JsonData{})
		return
	}
	if c.verifyAppId() {
		if len(data) > 5000 {
			c.ErrorJson(-100, "data len > 50000 ", JsonData{})
			return
		}
		account, err := models.SigningKeyHexStringAccount(signingKey)
		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}

		tx, e := models.ApiSpend(account, address, amount, data)
		time.Sleep(3 * time.Second)
		if e == nil {
			c.SuccessJson(map[string]interface{}{"tx": tx})
		} else {
			c.ErrorJson(-500, e.Error(), JsonData{})
		}
	} else {
		c.ErrorJson(-100, "appId or secret verify error", JsonData{})
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
