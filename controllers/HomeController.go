package controllers

import "C"
import (
	"ae/models"
	"ae/utils"
	"encoding/json"
	"strconv"
)

type MainController struct {
	BaseController
}
type LoginController struct {
	BaseController
}

type UserController struct {
	BaseController
}
type AccreditController struct {
	BaseController
}
type AccreditBindController struct {
	BaseController
}

type TestController struct {
	BaseController
}
type TestController2 struct {
	BaseController
}
type PayController struct {
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

			//var metaInfo MetaInfo
			//var totalSupply float64
			var balance Balance
			//_ = json.Unmarshal(metaInfoJson, &metaInfo)
			//_ = json.Unmarshal(totalSupplyJson, &totalSupply)
			_ = json.Unmarshal(balanceJson, &balance)
			//fmt.Println(totalSupply)

			tokens, e := models.FindTokenUserId(c.getCurrentUserId())
			if e != nil {
				c.TplName = "error2.html"
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

	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"

	if c.isLogin() {
		c.Redirect("/user", 302)
	} else {
		c.TplName = "login.html"
	}

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
					} else {
						c.Data["AppId"] = secret.AppId
						c.Data["AppSecret"] = secret.AppSecret
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

	//isSucess := models.Send("283122529@qq.com", "082273")
	//if(isSucess){
	//	c.Ctx.WriteString("发送成功")
	//}else{
	//	c.Ctx.WriteString("发送失败")
	//}

	//c.JsonData["Website"] = "beego.me"
	//c.JsonData["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"

	//entropy, _ := bip39.NewEntropy(128)
	//mnemonic, _ := bip39.NewMnemonic(entropy)
	//account, err, msg := models.MnemonicAccount("tail disagree oven fit state cube rule test economy claw nice stable")
	//signingKeyHex := account.SigningKeyToHexString()
	//if (err != nil) {
	//	c.Ctx.WriteString(msg)
	//}
	//c.Ctx.WriteString(account.Address+"-----")
	//c.Ctx.WriteString(signingKeyHex+"-----")
	//
	//acc :=models.SigningKeyHexStringAccount("659cc72a32b75ec5c488e7b2a788f6b94d9a8a6b5d8c0fcbaf3518a3d8f8eb8ef050d196f8a70d41b1ad22998be3020710f553e31033f9fa6f02752e4630d049")
	//
	//c.Ctx.WriteString(acc.Address+"-----")
	if c.isLogin() {
		c.Redirect("/user", 302)
	} else {
		if utils.IsMobile(c.Ctx.Input.Header("user-agent")) {
			c.TplName = "index_mobile.html"
		} else {
			c.TplName = "index.html"
		}

	}
}
func (c *PayController) Get() {
	orderNo := c.GetString("order_no")
	redirectUri := c.GetString("redirect_uri")
	if orderNo == "" || redirectUri == "" {
		c.Data["myAddress"] = "-"
		c.Data["address"] = "-"
		c.Data["tokens"] = "0"
		c.TplName = "pay.html"
		return
	}
	order, e := models.FindOrderOrderNo(orderNo)
	if e != nil {
		c.Data["myAddress"] = "-"
		c.Data["address"] = "-"
		c.Data["tokens"] = "0"
		c.TplName = "pay.html"
		return
	}
	c.Data["myAddress"] = order.SendAddress
	c.Data["address"] = order.ReceiveAddress
	c.Data["tokens"] = order.Tokens
	c.TplName = "pay.html"
}

func (c *ArticleInfoController) Get() {
	articleId := c.GetString("article_id")
	article, e := models.FindArticleId(articleId)
	if e != nil {
		c.TplName = "error.html"
	}
	c.Data["article"] = article.Content
	c.TplName = "article.html"
}
