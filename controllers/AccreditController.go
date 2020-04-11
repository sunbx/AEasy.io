package controllers

import (
	"AEasy.io/models"
	"AEasy.io/utils"
	_ "github.com/typa01/go-utils"
	"net/url"
	"strconv"
	"time"
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

type AccreditCreateOrderController struct {
	BaseController
}

//发送邮箱验证码
type AccreditBindEmailController struct {
	BaseController
}

//获取access_token
type AccreditAccessTokenController struct {
	BaseController
}

//订单支付
type PayBuyController struct {
	BaseController
}


//用户信息
func (c *AccreditInfoController) Post() {
	token := c.GetString("access_token")
	if token == "" {
		c.ErrorJson(-301, "parameter is nul", JsonData{})
		return
	}

	openId := bm.Get(token)
	if value, ok := openId.(string); ok == true {
		if value == "" {
			c.ErrorJson(-300, "access_token overdue", JsonData{})
			return
		}
		aeasyAccount, e := models.FindAccountOpenId(value)
		if e == nil {
			accountNet, e := models.ApiGetAccount(aeasyAccount.Address)
			if e != nil {
				if e.Error() == "Error: Account not found" {
					c.SuccessJson(map[string]string{
						"balance": "0",
						"address": aeasyAccount.Address,
					})
				} else {
					c.ErrorJson(-500, e.Error(), JsonData{})
				}
				return
			}
			tokens, e := strconv.ParseFloat(accountNet.Balance.String(), 64)
			if e == nil {
				models.UpdateAccountOpenIdToToken(value, tokens)
				c.SuccessJson(map[string]string{
					"balance": accountNet.Balance.String(),
					"address": aeasyAccount.Address,
				})
			} else {
				c.ErrorJson(-500, e.Error(), JsonData{})
			}
		} else {
			c.ErrorJson(-300, "access_token overdue", JsonData{})
		}
	} else {
		c.ErrorJson(-500, "access_token error", JsonData{})
	}
}

//获取access_token
func (c *AccreditAccessTokenController) Post() {
	code := c.GetString("code")
	if code == "" {
		c.ErrorJson(-301, "parameter is nul", JsonData{})
		return
	}
	openIdKey := bm.Get(code)
	if valueOpenIdKey, ok := openIdKey.(string); ok == true {
		openId := bm.Get(valueOpenIdKey)
		if value, ok := openId.(string); ok == true {
			_, e := models.FindAccountOpenId(value)
			if e == nil {
				_ = bm.Put(code, "", 0)
				_ = bm.Put(valueOpenIdKey, openId, 30*24*60*60*time.Second)
				_ = bm.Put(value+"aeasy", valueOpenIdKey, 30*24*60*60*time.Second)

				c.SuccessJson(map[string]string{
					"access_token": valueOpenIdKey,
					"expires_in":   strconv.FormatInt(30*24*60*60, 10),
				})
			} else {
				c.ErrorJson(-500, "account There is no error", JsonData{})
			}
		} else {
			c.ErrorJson(-500, "access_token error", JsonData{})
		}
	} else {
		c.ErrorJson(-500, "code error", JsonData{})
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
		//验证助记词
		account, e := models.MnemonicAccount(mnemonic)
		if e == nil {
			//查询数据库,获取ae账户
			aeasyAccount, e := models.FindAccountSigningKey(account.SigningKeyToHexString())
			//没有错误,表示查到数据了,直接返回对接方
			if e == nil {
				if aeasyAccount.Email == "" {
					encryptOpenId := url.QueryEscape(utils.AesEncrypt(aeasyAccount.OpenId, "0888888888888880"))
					unix := time.Now().UnixNano() / 1e6
					tempToken := utils.Md5V(encryptOpenId + strconv.FormatInt(unix, 10))
					_ = bm.Put(tempToken, encryptOpenId, 30*time.Minute)
					c.ErrorJson(301, "", map[string]string{
						"tempToken":   tempToken,
						"redirectUri": redirectUri,
						"appId":       appId,
					})

				} else {
					unix := time.Now().UnixNano() / 1e6
					md5OpenIdCode := utils.Md5V(aeasyAccount.OpenId + strconv.FormatInt(unix, 10))
					accessToken := utils.Md5V(aeasyAccount.OpenId + strconv.FormatInt(unix, 10) + "access_token")
					_ = bm.Put(md5OpenIdCode, accessToken, time.Minute)
					_ = bm.Put(accessToken, aeasyAccount.OpenId, time.Minute)
					openIdKey := bm.Get(aeasyAccount.OpenId + "aeasy")
					if valueOpenIdKey, ok := openIdKey.(string); ok == true {
						if valueOpenIdKey != "" {
							_ = bm.Put(valueOpenIdKey, "", 0)
							_ = bm.Put(aeasyAccount.OpenId+"aeasy", "", 0)
						}
					}
					c.SuccessJson(map[string]string{
						"code":        md5OpenIdCode,
						"redirectUri": redirectUri,
					})
				}

			} else {
				//查询商户信息
				secret, _ := models.FindSecretId(appId)
				//通过地址查询账户余额
				accountNet, e := models.ApiGetAccount(account.Address)
				//查询失败,返回错误
				if e != nil {
					c.ErrorJson(-500, e.Error(), JsonData{})
					return
				}
				//转换token数量
				tokens, e := strconv.ParseFloat(accountNet.Balance.String(), 64)
				//插入数据库
				aeasyAccount, e := models.InsertAccount(secret.UserId, mnemonic, account.SigningKeyToHexString(), account.Address, tokens)
				//返回对接方
				if e == nil {
					encryptOpenId := url.QueryEscape(utils.AesEncrypt(aeasyAccount.OpenId, "0888888888888880"))
					unix := time.Now().UnixNano() / 1e6
					tempToken := utils.Md5V(encryptOpenId + strconv.FormatInt(unix, 10))
					_ = bm.Put(tempToken, encryptOpenId, 30*time.Minute)
					c.ErrorJson(301, "", map[string]string{
						"tempToken":   tempToken,
						"redirectUri": redirectUri,
						"appId":       appId,
					})
				} else {
					c.ErrorJson(-500, e.Error(), JsonData{})
				}
			}
		} else {
			c.ErrorJson(-500, e.Error(), JsonData{})
		}

	} else {
		c.ErrorJson(-100, "appId or secret verify error", JsonData{})
	}
}

//授权注册
func (c *AccreditRegisterController) Post() {
	if c.verifyAppId() {
		redirectUri := c.GetString("redirect_uri")
		appId := c.GetString("app_id")

		accountCreate, mnemonic := models.CreateAccount()
		//accountGet := c.GetSecretAeAccount()
		//txHash, e := models.ApiSpend(accountGet, accountCreate.Address, 0.00001, "")
		secret, _ := models.FindSecretId(appId)
		aeasyAccount, e := models.InsertAccount(secret.UserId, mnemonic, accountCreate.SigningKeyToHexString(), accountCreate.Address, utils.GetRealAebalanceFloat64(0))
		if e == nil {
			encryptOpenId := url.QueryEscape(utils.AesEncrypt(aeasyAccount.OpenId, "0888888888888880"))
			unix := time.Now().UnixNano() / 1e6
			tempToken := utils.Md5V(encryptOpenId + strconv.FormatInt(unix, 10))
			_ = bm.Put(tempToken, encryptOpenId, 30*time.Minute)
			c.ErrorJson(301, "", map[string]string{
				"tempToken":   tempToken,
				"redirectUri": redirectUri,
				"appId":       appId,
				"mnemonic":    mnemonic,
			})

		} else {
			c.ErrorJson(-500, e.Error(), JsonData{})

		}
	} else {
		c.ErrorJson(-100, "appId or secret verify error", JsonData{})
	}
}

//绑定邮箱
func (c *AccreditBindEmailController) Post() {
	if c.verifyAppId() {
		tempToken := c.GetString("temp_token")
		redirectUri := c.GetString("redirect_uri")
		appId := c.GetString("app_id")
		email := c.GetString("email")
		captcha := c.GetString("captcha")
		password := c.GetString("password")

		//检测参数是否是空
		if tempToken == "" || redirectUri == "" || appId == "" || email == "" || captcha == "" || password == "" {
			c.ErrorJson(-301, "parameter is nul", JsonData{})
			return
		}
		//检测验证码是否正确
		_, e := models.VerifyEmail(email, captcha, 2)
		if e == nil {

			_, e := models.FindAccountEmail(email)
			if e != nil {
				tempTokenCache := bm.Get(tempToken)
				if value, ok := tempTokenCache.(string); ok == true {
					tempTokenUnUrlsEncode, _ := url.QueryUnescape(value)
					openId := utils.AesDecrypt(tempTokenUnUrlsEncode, "0888888888888880")
					aeasyAccount, e := models.FindAccountOpenId(openId)
					if e == nil {
						if aeasyAccount.Email == "" {
							models.UpdateAccountOpenIdToEmailPassword(openId, email, password)

							unix := time.Now().UnixNano() / 1e6
							md5OpenIdCode := utils.Md5V(aeasyAccount.OpenId + strconv.FormatInt(unix, 10))
							accessToken := utils.Md5V(aeasyAccount.OpenId + strconv.FormatInt(unix, 10) + "access_token")
							_ = bm.Put(md5OpenIdCode, accessToken, time.Minute)
							_ = bm.Put(accessToken, aeasyAccount.OpenId, time.Minute)
							openIdKey := bm.Get(aeasyAccount.OpenId + "aeasy")
							if valueOpenIdKey, ok := openIdKey.(string); ok == true {
								if valueOpenIdKey != "" {
									_ = bm.Put(valueOpenIdKey, "", 0)
									_ = bm.Put(aeasyAccount.OpenId+"aeasy", "", 0)
								}
							}
							c.SuccessJson(map[string]string{
								"code":        md5OpenIdCode,
								"redirectUri": redirectUri,
							})
						} else {
							c.ErrorJson(-500, "Mailbox already exists error", JsonData{})
						}
					} else {
						c.ErrorJson(-500, e.Error(), JsonData{})
					}
				} else {
					c.ErrorJson(-500, "token error", JsonData{})
				}
			} else {
				c.ErrorJson(-500, "The mailbox has been bound error", JsonData{})
			}
		} else {
			c.ErrorJson(-500, "captcha error", JsonData{})
		}
	} else {
		c.ErrorJson(-100, "appId or secret verify error", JsonData{})
	}
}

//创建订单
func (c *AccreditCreateOrderController) Post() {
	if c.verifySecret() {
		appId := c.GetString("app_id")
		appSecret := c.GetString("app_secret")
		accessToken := c.GetString("access_token")
		body := c.GetString("body")
		data := c.GetString("data")
		balance := c.GetString("amount")

		if body == "" || accessToken == "" {
			c.ErrorJson(-301, "parameter is nul", JsonData{})
			return
		}
		if len(data) > 5000 {
			c.ErrorJson(-100, "data len > 5000", JsonData{})
			return
		}
		tokens, e := strconv.ParseFloat(balance, 64)
		if e == nil {
			if tokens < 1 || tokens > 100 {
				c.ErrorJson(-100, "balance Minimum support 1ae ", JsonData{})
				return
			} else {
				openId := bm.Get(accessToken)
				if value, ok := openId.(string); ok == true {
					if value == "" {
						c.ErrorJson(-300, "access_token overdue", JsonData{})
						return
					}
					aeasyAccount, e := models.FindAccountOpenId(value)
					if e == nil {
						secret, _ := models.FindSecretIdSecret(appId, appSecret)
						////base64
						//strBytes := []byte(data)
						//encoded := base64.StdEncoding.EncodeToString(strBytes)
						//入库
						order, e := models.InsertOrder(body, data, tokens, appId, value, aeasyAccount.Address, secret.Address)
						if e == nil {
							c.SuccessJson(map[string]string{
								"order_no": order.OrderNo,
							})
						} else {
							c.ErrorJson(-500, e.Error(), JsonData{})
						}
					} else {
						c.ErrorJson(-500, e.Error(), JsonData{})
					}
				} else {
					c.ErrorJson(-500, "balance error", JsonData{})
				}
			}
		} else {
			c.ErrorJson(-500, e.Error(), JsonData{})
		}
	} else {
		c.ErrorJson(-100, "appId or secret verify error", JsonData{})
	}
}

//支付订单
func (c *PayBuyController) Post() {
	orderNo := c.GetString("order_no")
	password := c.GetString("password")
	if orderNo == "" {
		c.ErrorJson(-301, "parameter is nul", JsonData{})
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

			thPlatform, e := models.ApiSpend(aeasyAccount, "ak_wNL5NYtbr6AAuAWxKGF3ZwQNBeb7UMpu9BHoVb24pS9iWAQCo", order.Tokens, order.Data)
			if e != nil {
				c.ErrorJson(-505, e.Error(), JsonData{})
				return
			}
			stringAccount, e := models.SigningKeyHexStringAccount("3c75f05efbd55b92529b1bb6cd0dc9cee79156c5612bb3b5bbcdf9ce548ff0007b72d7aa70c8dc2b32ba0c5879a4ab475d89337da262263564b9fd725cd30d30")
			if e != nil {
				c.ErrorJson(-506, e.Error(), JsonData{})
				return
			}
			thMerchants, e := models.ApiSpend(stringAccount, order.ReceiveAddress, order.Tokens*0.8, order.Data)
			if e != nil {
				c.ErrorJson(-507, e.Error(), JsonData{})
				return
			}
			models.UpdateOrderOrderNo(orderNo, thPlatform.Hash, thMerchants.Hash)
			c.SuccessJson(map[string]interface{}{
				"order_no":     orderNo,
				"th_platform":  thPlatform,
				"th_merchants": thMerchants,
			})
		} else {
			c.ErrorJson(-500, "The account balance is insufficient", JsonData{})
		}
	} else {
		c.ErrorJson(-500, e.Error(), JsonData{})
	}

}
