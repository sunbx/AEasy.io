package controllers

import (
	"ae/models"
	"ae/utils"
	"github.com/beego/i18n"
	"github.com/shopspring/decimal"
	_ "github.com/typa01/go-utils"
	"strconv"
)

type AccreditLoginController struct {
	BaseController
}
type AccreditRegisterController struct {
	BaseController
}

type AccreditInfoController struct {
	BaseController
}

//订单支付
type PayBuyController struct {
	BaseController
}

//用户信息
func (c *AccreditInfoController) Post() {
	if c.verifyAppId() {
		signingKey := c.GetString("signingKey")
		address := c.GetString("address")
		if signingKey == "" && address == "" {
			c.ErrorJson(-301,  i18n.Tr(c.getHeaderLanguage(),"parameter is nul"), JsonData{})
			return
		}

		if signingKey != "" {
			aeasyAccount, e := models.SigningKeyHexStringAccount(signingKey)
			if e == nil {
				address = aeasyAccount.Address
			}
		}
		accountNet, e := models.ApiGetAccount(address)
		if e != nil {
			if e.Error() == "Error: Account not found" {
				c.SuccessJson(map[string]string{
					"balance": "0.00000",
					"address": address,
				})
			} else {
				c.ErrorJson(-500, e.Error(), JsonData{})
			}
			return
		}
		decimalValue, _ := decimal.NewFromString(accountNet.Balance.String())
		f, _ := decimalValue.Float64()
		c.SuccessJson(map[string]string{
			"balance": utils.FormatTokens(f),
			"address": address,
		})
	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(),"appId or secret verify error"), JsonData{})
	}

}

//授权登录
func (c *AccreditLoginController) Post() {
	//验证key
	if c.verifyAppId() {
		//获取参数
		redirectUri := c.GetString("redirect_uri")
		appId := c.GetString("app_id")
		mnemonic := c.GetString("mnemonic")
		indexAddress, _ := c.GetUint32("index_address", 1)
		//验证助记词
		account, e := models.MnemonicAccount(mnemonic, indexAddress)
		if e == nil {
			//查询数据库,获取ae账户
			_, e := models.FindAccountSigningKey(utils.Md5V(account.SigningKeyToHexString() + "aeasy"))
			//没有错误,表示查到数据了,直接返回对接方
			if e == nil {
				c.SuccessJson(map[string]string{
					"redirectUri": redirectUri,
					"mnemonic":    mnemonic,
					"address":     account.Address,
					"signingKey":  account.SigningKeyToHexString(),
				})
			} else {
				//查询商户信息
				secret, _ := models.FindSecretId(appId)
				//插入数据库
				_, e = models.InsertAccount(secret.UserId, utils.Md5V(account.SigningKeyToHexString()+"aeasy"), account.Address)
				if e == nil {
					c.SuccessJson(map[string]string{
						"redirectUri": redirectUri,
						"mnemonic":    mnemonic,
						"address":     account.Address,
						"signingKey":  account.SigningKeyToHexString(),
					})
				} else {
					c.ErrorJson(-500, e.Error(), JsonData{})
				}
			}
		} else {
			c.ErrorJson(-500, e.Error(), JsonData{})
		}

	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(),"appId or secret verify error"), JsonData{})
	}
}

//授权注册
func (c *AccreditRegisterController) Post() {
	if c.verifyAppId() {
		redirectUri := c.GetString("redirect_uri")
		appId := c.GetString("app_id")

		accountCreate, mnemonic := models.CreateAccount()

		secret, _ := models.FindSecretId(appId)
		_, e := models.InsertAccount(secret.UserId, utils.Md5V(accountCreate.SigningKeyToHexString()+"aeasy"), accountCreate.Address)
		if e == nil {
			c.SuccessJson(map[string]string{
				"redirectUri": redirectUri,
				"mnemonic":    mnemonic,
				"address":     accountCreate.Address,
				"signingKey":  accountCreate.SigningKeyToHexString(),
			})
		} else {
			c.ErrorJson(-500, e.Error(), JsonData{})
		}
	} else {
		c.ErrorJson(-100, i18n.Tr(c.getHeaderLanguage(),"appId or secret verify error"), JsonData{})
	}
}

//支付订单
func (c *PayBuyController) Post() {
	orderNo := c.GetString("order_no")
	password := c.GetString("password")
	if orderNo == "" {
		c.ErrorJson(-301, i18n.Tr(c.getHeaderLanguage(),"parameter is nul"), JsonData{})
		return
	}
	order, e := models.FindOrderOrderNo(orderNo)
	if e == nil {
		if order.PayStatus != 1 {
			c.ErrorJson(-501, "The order has been paid.", JsonData{})
			return
		}
		account, e := models.FindAccountOpenId(order.OpenId)
		if e != nil {
			c.ErrorJson(-501, e.Error(), JsonData{})
			return
		}
		if utils.Md5V(password) != account.Password {
			c.ErrorJson(-511, "password error", JsonData{})
			return
		}
		//获取账户
		aeasyAccount, e := models.SigningKeyHexStringAccount(account.SigningKey)
		if e != nil {
			c.ErrorJson(-502, e.Error(), JsonData{})
			return
		}
		//获取账户信息
		accountNet, e := models.ApiGetAccount(account.Address)
		if e != nil {
			c.ErrorJson(-503, e.Error(), JsonData{})
			return
		}
		tokens, e := strconv.ParseFloat(accountNet.Balance.String(), 64)
		if e != nil {
			c.ErrorJson(-504, e.Error(), JsonData{})
		}
		if tokens/1000000000000000000 >= order.Tokens {

			thMerchants, e := models.ApiSpend(aeasyAccount, order.ReceiveAddress, order.Tokens, order.Data)
			if e != nil {
				c.ErrorJson(-507, e.Error(), JsonData{})
				return
			}

			models.UpdateOrderOrderNo(orderNo, "", thMerchants.Hash)
			c.SuccessJson(map[string]interface{}{
				"order_no":     orderNo,
				"th_merchants": thMerchants,
			})
		} else {
			c.ErrorJson(-500, "The account balance is insufficient", JsonData{})
		}
	} else {
		c.ErrorJson(-500, e.Error(), JsonData{})
	}

}
