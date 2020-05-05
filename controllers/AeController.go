package controllers

import (
	"ae/models"
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
			c.ErrorJson(-301, "The balance should be greater than 1ae", JsonData{})
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

//创建用户
func (c *ApiCreateAccountController) Post() {
	accountCreate, mnemonic := models.CreateAccount()

	c.SuccessJson(map[string]string{
		"mnemonic":    mnemonic,
		"address":     accountCreate.Address,
		"signing_key": accountCreate.SigningKeyToHexString(),
	})
}

