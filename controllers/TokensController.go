package controllers

import (
	"ae/models"
	"fmt"
	"github.com/beego/i18n"
	"strconv"
)

//订单支付
type TokenCreateController struct {
	BaseController
}

type TokenTransferController struct {
	BaseController
}


//tokens 创建
func (c *TokenCreateController) Post() {
	name := c.GetString("name")
	countString := c.GetString("count")
	if countString == "" || name == "" {
		c.ErrorJson(-301, i18n.Tr(c.getHeaderLanguage(),"parameter is nul"), JsonData{})
		return
	}
	if len(name) < 2 || len(name) > 5 {
		c.ErrorJson(-301, "parameter is name len < 2 and > 5", JsonData{})
		return
	}
	count, err := strconv.Atoi(countString)
	if err != nil {
		c.ErrorJson(-301, err.Error(), JsonData{})
		return
	}
	if count > 999999999999 {
		c.ErrorJson(-301, "parameter is count > 999999999999", JsonData{})
		return
	}
	secret, err := models.FindSecretUserID(c.getCurrentUserId())
	if err != nil {
		c.ErrorJson(-301, err.Error(), JsonData{})
		return
	}
	stringAccount, err := models.SigningKeyHexStringAccount(secret.SigningKey)
	if err != nil {
		c.ErrorJson(-301, err.Error(), JsonData{})

		return
	}
	//_, err = models.ApiSpend(stringAccount, "ak_wNL5NYtbr6AAuAWxKGF3ZwQNBeb7UMpu9BHoVb24pS9iWAQCo", 1000, "")
	//if err != nil {
	//	c.ErrorJson(-301, err.Error(), JsonData{})
	//	return
	//}

	if !models.Is1AE(stringAccount.Address) {
		c.ErrorJson(-301, i18n.Tr(c.getHeaderLanguage(),"The balance should be greater than 1ae"), JsonData{})
		return
	}

	s, e := models.CompileContractInit(stringAccount, name, strconv.Itoa(count)+"000000000000000000")

	if e == nil {
		models.UpdateSecretContracts(c.getCurrentUserId(), s)
		_, err = models.InsertToken(c.getCurrentUserId(), name, s, secret.Address, strconv.Itoa(count)+"000000000000000000")
		c.SuccessJson(map[string]interface{}{
			"contract": s,
		})
	} else {
		c.ErrorJson(-500, e.Error(), nil)
	}
}

func (c *TokenTransferController) Post() {
	address := c.GetString("address")
	countString := c.GetString("count")
	if c.isLogin() {
		if countString == "" || address == "" {
			c.ErrorJson(-301, i18n.Tr(c.getHeaderLanguage(),"parameter is nul"), JsonData{})
			return
		}

		count, err := strconv.ParseFloat(countString, 64)
		if err != nil {
			c.ErrorJson(-301, err.Error(), JsonData{})
			return
		}
		if count > 999999999999 {
			c.ErrorJson(-301, "parameter is count > 999999999999", JsonData{})
			return
		}
		secret, e := models.FindSecretUserID(c.getCurrentUserId())
		if e != nil {
			c.ErrorJson(-500, e.Error(), JsonData{})
			return
		}
		if secret.Contracts == "" {
			c.ErrorJson(-500, "contracts error", JsonData{})
			return
		}
		account, _ := models.SigningKeyHexStringAccount(secret.SigningKey)
		fmt.Println(strconv.FormatFloat(count*1000000000000000000, 'f', -1, 64))
		_, err = models.CallContractFunction(account, secret.Contracts, "transfer", []string{address, strconv.FormatFloat(count*1000000000000000000, 'f', -1, 64)})

		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}
		c.SuccessJson(JsonData{})
	} else {
		c.TplName = "index.html"
		return
	}
}

