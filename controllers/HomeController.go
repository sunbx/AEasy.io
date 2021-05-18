package controllers

import (
	"ae/models"
	"ae/utils"
	"encoding/json"
	"fmt"
	"github.com/aeternity/aepp-sdk-go/config"
	"github.com/aeternity/aepp-sdk-go/naet"
	"github.com/shopspring/decimal"
	"io/ioutil"
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

	//node := naet.NewNode(models.NodeURL, false)
	//ttler := transactions.CreateTTLer(node)
	//noncer := transactions.CreateNoncer(node)
	//ttlNoncer := transactions.CreateTTLNoncer(ttler, noncer)
	//spendTx, _ := transactions.NewSpendTx("ak_idkx6m3bgRr7WiKXuB8EBYBoRqVsaSc6qo4dsd23HKgj3qiCF", "ak_CNcf2oywqbgmVg3FfKdbHQJfB959wrVwqfzSpdWVKZnep7nj4", utils.GetRealAebalanceBigInt(0.0001), []byte(""), ttlNoncer)
	//txRaw, _ := rlp.EncodeToBytes(spendTx)
	//msg := append([]byte("ae_mainnet"), txRaw...)
	//serializeTx, _ := transactions.SerializeTx(spendTx)
	//decodeMsg := hex.EncodeToString(msg)
	//c.SuccessJson(map[string]interface{}{
	//	"tx":  serializeTx,
	//	"msg": decodeMsg})
	//address := ""
	//for {
	//	account, s := models.CreateAccount()
	//	address = account.Address
	//	content := address[ len(address)-6 : len(address)]
	//	fmt.Println(address," - ",s," - ",content)
	//	res := address+" - "+s+" - "+content
	//	c.Ctx.WriteString(string(res))
	//	if strings.ContainsAny(content, "baixin"){
	//		break
	//	}
	//	time.Sleep(100)
	//
	//}
	//从 node 获取微块详细信息
	response := utils.Get(models.NodeURL + "/v2/micro-blocks/hash/" + "mh_G4kfZw6bjazQL3rNTkkZsdCD9k8s3pRRcvEUDdMePKXTtS4gd" + "/transactions")

	//解析微块信息
	var block MicroBlock
	err := json.Unmarshal([]byte(response), &block)
	if err != nil {
		return
	}

	mapObj, err := Obj2map(block.Transactions[1].Tx)
	if err != nil {
		fmt.Println("Obj2map error", err.Error())
		return
	}
	responseContractCode := utils.Get(models.NodeURL + "/v2/contracts/" + mapObj["contract_id"].(string) + "/code")

	var code interface{}
	err = json.Unmarshal([]byte(responseContractCode), &code)
	if err != nil {
		return
	}
	codeMap, err := Obj2map(code)
	compile := naet.NewCompiler("https://compiler.aeasy.io", false)
	decodedData, err := compile.DecodeCalldataBytecode(codeMap["bytecode"].(string), mapObj["call_data"].(string), config.Compiler.Backend)
	decodedDataJson, _ := json.Marshal(decodedData)
	var contractDecode ContractDecode
	err = json.Unmarshal(decodedDataJson, &contractDecode)
	if err != nil {
		fmt.Println("Obj2map error", err.Error())
		return
	}

	//只有aex9合约才记录
	aex9Amount := 0.0
	amount := 0.0
	aex9ReceiveAddress := ""

	amountFloat, _ := mapObj["amount"].(float64)
	amountFrom := decimal.NewFromFloat(amountFloat)
	amount, _ = amountFrom.Float64()
	feeFloat, _ := mapObj["fee"].(float64)
	feeForm := decimal.NewFromFloat(feeFloat)
	fee, _ := feeForm.Float64()
	hash := block.Transactions[1].Hash
	function := contractDecode.Function
	decodeJson := string(decodedDataJson)
	contractId := mapObj["contract_id"].(string)
	callAddress := mapObj["caller_id"].(string)
	tokens, _ := ioutil.ReadFile("conf/tokens.json")

	if strings.Contains(string(tokens), contractId) && function == "transfer" {
		argumentsAddress := decodedData.Arguments[0].(map[string]interface{})
		argumentsAmount := decodedData.Arguments[1].(map[string]interface{})
		aex9ReceiveAddress = argumentsAddress["value"].(string)
		aex9AmountFloat, _ := argumentsAmount["value"].(json.Number).Float64()
		aex9AmountFloatDecimal := decimal.NewFromFloat(aex9AmountFloat)
		aex9Amount, _ = aex9AmountFloatDecimal.Float64()
	}

	responseContractInfo := utils.Get(models.NodeURL + "/v2/transactions/" + hash + "/info")
	var info ResultInfo
	err = json.Unmarshal([]byte(responseContractInfo), &info)
	if err != nil {
		fmt.Println("Obj2map error", err.Error())
		return
	}

	returnType := info.CallInfo.ReturnType
	_, err = models.InsertContract("mh_2URQJSVGn8o7tCAdDPGXRLNzG82t53pYJwbt5dC2J4aTwEdhTf", 1, hash, function, decodeJson, contractId, callAddress, aex9Amount, amount, fee, returnType, aex9ReceiveAddress, 0)
	if err != nil {
		fmt.Println("Obj2map error", err.Error())
		return
	}
	c.SuccessJson(contractDecode)

}

func Obj2map(obj interface{}) (mapObj map[string]interface{}, err error) {
	// 结构体转json
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type ResultInfo struct {
	CallInfo struct {
		CallerID    string        `json:"caller_id"`
		CallerNonce int           `json:"caller_nonce"`
		ContractID  string        `json:"contract_id"`
		GasPrice    int           `json:"gas_price"`
		GasUsed     int           `json:"gas_used"`
		Height      int           `json:"height"`
		Log         []interface{} `json:"log"`
		ReturnType  string        `json:"return_type"`
		ReturnValue string        `json:"return_value"`
	} `json:"call_info"`
}

type ContractDecode struct {
	Arguments []interface{} `json:"arguments"`
	Function  string        `json:"function"`
}

type MicroBlock struct {
	Transactions []Transactions `json:"transactions"`
}

type Transactions struct {
	BlockHash   string      `json:"block_hash"`
	BlockHeight int64       `json:"block_height"`
	Hash        string      `json:"hash"`
	Signatures  []string    `json:"signatures"`
	Tx          interface{} `json:"tx"`
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
