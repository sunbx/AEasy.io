package controllers

import (
	 "ae/models"
	"ae/utils"
	"encoding/json"
	"fmt"

	"strconv"
	"strings"
)

type MainController struct {
	BaseController
}
type LanguageController struct {
	BaseController
}
type LoginController struct {
	BaseController
}

type UserController struct {
	BaseController
}
type ShowController struct {
	BaseController
}
type AccreditController struct {
	BaseController
}
type AccreditBindController struct {
	BaseController
}



type TestController5 struct {
	BaseController
}


type TokenController struct {
	BaseController
}

type ArticleInfoController struct {
	BaseController
}

type MetaInfo struct {
	Decimals int    `json:"decimals"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
}
type Balance struct {
	Some []float64 `json:"Some"`
}

func (c *TestController5) Get() {

	account, _ := models.SigningKeyHexStringAccount("0b5a745770a17cb571e58eac04b0e536db1818dc9036216252aa749e848563fb56a1fc4ef38cbc5ddcb6b23307b0c535402e72a932c9087c8a22cabe6edd030f")

	for i := 0; i <= 100; i++ {
		createAccount, _ := models.CreateAccount()
		tx, _ := models.ApiSpend(account, createAccount.Address, 0.0001, "")
		println("tx->"+strconv.Itoa(i)+" ", tx.Hash,)
	}

	//c.SuccessJson(map[string]interface{}{
	//	"sing": account.SigningKeyToHexString(),})
}

func (c *AccreditController) Get() {
	if c.verifyAppId() {
		c.TplName = "accredit.html"
	} else {
		c.TplName = "error.html"
	}
}
func (c *TokenController) Get() {
	if c.isLogin() {
		secret, e := models.FindSecretUserID(c.getCurrentUserId())
		if e != nil {
			c.TplName = "index.html"
			return
		}
		if secret.Contracts == "" {
			var balance string
			account, e := models.ApiGetAccount(secret.Address)
			if e != nil {
				balance = strconv.FormatFloat(secret.Tokens, 'f', 0, 64)
			} else {
				balance = account.Balance.String()
			}
			tokens, e := strconv.ParseFloat(balance, 64)
			if e == nil {
				if secret.IsShow == 0 && tokens/1000000000000000000 > 1 {
					secret.IsShow = 1
				}
				models.UpdateSecretTokens(secret.Address, tokens, secret.IsShow)
				content := utils.FormatTokens(tokens)
				c.Data["Token"] = content
			}
			c.TplName = "tokens_create.html"
		} else {
			account, _ := models.SigningKeyHexStringAccount(secret.SigningKey)

			if !models.Is1AE(account.Address) {
				c.TplName = "error2.html"
				return
			}
			//metaInfoCall, _ := models.CallContractFunction(account, secret.Contracts, "meta_info", []string{})
			//totalSupplyCall, _ := models.CallContractFunction(account, secret.Contracts, "total_supply", []string{})
			balanceCall, _ := models.CallContractFunction(account, secret.Contracts, "balance", []string{account.Address})
			//metaInfoJson, _ := json.Marshal(&metaInfoCall)
			//totalSupplyJson, _ := json.Marshal(&totalSupplyCall)
			balanceJson, _ := json.Marshal(&balanceCall)
			fmt.Println(balanceCall)
			//var metaInfo MetaInfo
			//var totalSupply float64
			var balance Balance
			//_ = json.Unmarshal(metaInfoJson, &metaInfo)
			//_ = json.Unmarshal(totalSupplyJson, &totalSupply)
			_ = json.Unmarshal(balanceJson, &balance)
			//fmt.Println(totalSupply)

			tokens, e := models.FindTokenUserId(c.getCurrentUserId())
			if e != nil {
				c.TplName = "error.html"
				return
			}
			count, _ := strconv.ParseFloat(tokens.Count, 64)
			c.Data["total_supply"] = utils.FormatTokens(count)
			c.Data["balance"] = utils.FormatTokens(balance.Some[0])
			c.Data["decimals"] = tokens.Decimals
			c.Data["name"] = tokens.Name
			c.Data["symbol"] = tokens.Name
			c.Data["contracts"] = secret.Contracts
			c.Data["address"] = tokens.Address
			c.TplName = "tokens_detail.html"
		}
	} else {
		c.TplName = "index.html"
		return
	}
}

func (c *AccreditBindController) Get() {
	if c.verifyAppId() {
		c.TplName = "accredit_bind.html"
	} else {
		c.TplName = "error.html"
	}
}
func (c *LoginController) Get() {

	if c.isLogin() {
		c.Redirect("/user", 302)
	} else {
		c.TplName = "login.html"
	}

}

func (c *ShowController) Get() {
	isShow := c.Ctx.GetCookie("isShow")
	if isShow == "show" {
		c.Ctx.SetCookie("isShow", "")
	} else {
		c.Ctx.SetCookie("isShow", "show")
	}
	c.Redirect("/user", 302)
}

func (c *UserController) Get() {
	if c.isLogin() {
		v := c.GetSession("user_id")
		c.Data["address"] = "-"
		c.Data["token"] = "-"
		c.Data["appId"] = "-"
		c.Data["appSecret"] = "-"

		if userId, ok := v.(int); ok == true {
			if userId > 0 {
				secret, e := models.FindSecretUserID(userId)
				if e != nil {
					c.Redirect("/", 302)
					return
				}
				var balance string
				account, e := models.ApiGetAccount(secret.Address)
				if e != nil {
					balance = strconv.FormatFloat(secret.Tokens, 'f', 0, 64)
				} else {
					balance = account.Balance.String()
				}

				tokens, e := strconv.ParseFloat(balance, 64)
				if e == nil {
					if secret.IsShow == 0 && tokens/1000000000000000000 >= 1 {
						secret.IsShow = 1
					}
					models.UpdateSecretTokens(secret.Address, tokens, secret.IsShow)
					c.Data["Address"] = secret.Address
					//c.Data["Token"] = secret.Tokens
					content := utils.FormatTokens(tokens)
					c.Data["Token"] = content
					if secret.IsShow == 0 {
						c.Data["AppId"] = "**** **** **** ****"
						c.Data["AppSecret"] = "**** **** **** **** **** **** **** ****"
						c.Data["is_zhegyan"] = "display: none;"
						c.Data["is_biyan"] = "display: none;"
					} else {

						isShow := c.Ctx.GetCookie("isShow")
						if isShow == "show" {
							c.Data["AppId"] = secret.AppId
							c.Data["AppSecret"] = secret.AppSecret
							c.Data["BY"] = "display: none"
						} else {
							c.Data["AppId"] = "**** **** **** ****"
							c.Data["AppSecret"] = "**** **** **** **** **** **** **** ****"
							c.Data["ZY"] = "display: none"
						}

					}
				}
			}
		}
		c.TplName = "user.html"

	} else {
		c.TplName = "login.html"
	}

}

func (c *MainController) Get() {
	if c.isLogin() {
		c.Redirect("/user", 302)
	} else {
		if utils.IsMobile(c.Ctx.Input.Header("user-agent")) {
			c.TplName = "index_mobile.html"
		} else {

			c.Data["bye"] = c.Tr("bye")
			c.TplName = "index.html"
		}
	}
}

func (c *LanguageController) Get() {

	var language = c.Ctx.GetCookie("language")
	if strings.Contains(language, "zh-CN") || strings.Contains(language, "zh-cn") {
		c.Lang = "en-US"
	} else {
		c.Lang = "zh-CN"
	}

	fmt.Printf("language", language)
	fmt.Printf("c.Lang", c.Lang)

	c.Ctx.SetCookie("language", c.Lang)
	c.Redirect("/", 302)
}

//func (c *PayController) Get() {
//	orderNo := c.GetString("order_no")
//	redirectUri := c.GetString("redirect_uri")
//	if orderNo == "" || redirectUri == "" {
//		c.Data["myAddress"] = "-"
//		c.Data["address"] = "-"
//		c.Data["tokens"] = "0"
//		c.TplName = "pay.html"
//		return
//	}
//	order, e := models.FindOrderOrderNo(orderNo)
//	if e != nil {
//		c.Data["myAddress"] = "-"
//		c.Data["address"] = "-"
//		c.Data["tokens"] = "0"
//		c.TplName = "pay.html"
//		return
//	}
//	c.Data["myAddress"] = order.SendAddress
//	c.Data["address"] = order.ReceiveAddress
//	c.Data["tokens"] = order.Tokens
//	c.TplName = "pay.html"
//}

func (c *ArticleInfoController) Get() {
	articleId := c.GetString("article_id")
	article, e := models.FindArticleId(articleId)
	if e != nil {
		c.TplName = "error.html"
	}
	c.Data["article"] = article.Content
	c.TplName = "article.html"
}
