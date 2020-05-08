package controllers

import (
	"ae/models"
	"ae/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aeternity/aepp-sdk-go/transactions"
	"github.com/shopspring/decimal"
	"math/big"
	"strconv"
)

type NamesAuctionsActiveController struct {
	BaseController
}

type NamesPriceController struct {
	BaseController
}

type NamesMyRegisterController struct {
	BaseController
}
type NamesMyOverController struct {
	BaseController
}
type NamesOverController struct {
	BaseController
}
type NamesUpdateController struct {
	BaseController
}

type NamesInfoController struct {
	BaseController
}
type NamesAddController struct {
	BaseController
}
type NamesTransferController struct {
	BaseController
}

type Names struct {
	Name           string `json:"name"`
	Expiration     int    `json:"expiration"`
	ExpirationTime string `json:"expiration_time"`
	WinningBid     string `json:"winning_bid"`
	WinningBidder  string `json:"winning_bidder"`
}

type NamesMy struct {
	Name             string `json:"name"`
	NameHash         string `json:"name_hash"`
	TxHash           string `json:"tx_hash"`
	CreatedAtHeight  int    `json:"created_at_height"`
	AuctionEndHeight int    `json:"auction_end_height"`
	Owner            string `json:"owner"`
	ExpiresAt        int    `json:"expires_at"`
	ExpirationTime   string `json:"expires_time"`
}

type NamesMyBid struct {
	NameAuctionEntry Names `json:"name_auction_entry"`
}

func (c *NamesAuctionsActiveController) Post() {
	if c.verifyAppId() {
		page, _ := c.GetInt("page", 1)
		height := int(models.ApiBlocksTop())
		namesDb, e := models.FindNameAuctionOver(page, height)
		if e != nil {
			c.ErrorJson(-500, e.Error(), JsonData{})
		}
		var names []map[string]interface{}
		for i := 0; i < len(namesDb); i++ {
			var name = map[string]interface{}{}
			name["name"] = namesDb[i].Name
			name["length"] = namesDb[i].Length
			name["start_height"] = namesDb[i].StartHeight
			name["end_height"] = namesDb[i].EndHeight
			name["over_height"] = namesDb[i].OverHeight
			name["current_height"] = height
			name["owner"] = namesDb[i].Owner
			name["current_price"] = utils.FormatTokens(namesDb[i].CurrentPrice)
			name["th_hash"] = namesDb[i].ThHash

			names = append(names, name)
		}
		if names == nil{
			c.SuccessJson([]JsonData{})
			return
		}
		c.SuccessJson(names)
	} else {
		c.ErrorJson(-100, "appId or secret verify error", JsonData{})
	}
}

func (c *NamesPriceController) Post() {
	if c.verifyAppId() {
		page, _ := c.GetInt("page", 1)
		height := int(models.ApiBlocksTop())
		namesDb, e := models.FindNameAuctionPrice(page, height)
		if e != nil {
			c.ErrorJson(-500, e.Error(), JsonData{})
		}
		var names []map[string]interface{}
		for i := 0; i < len(namesDb); i++ {
			var name = map[string]interface{}{}
			name["name"] = namesDb[i].Name
			name["length"] = namesDb[i].Length
			name["start_height"] = namesDb[i].StartHeight
			name["end_height"] = namesDb[i].EndHeight
			name["over_height"] = namesDb[i].OverHeight
			name["current_height"] = height
			name["owner"] = namesDb[i].Owner
			name["current_price"] =  utils.FormatTokens(namesDb[i].CurrentPrice)
			name["th_hash"] = namesDb[i].ThHash

			names = append(names, name)
		}
		if names == nil{
			c.SuccessJson([]JsonData{})
			return
		}
		c.SuccessJson(names)
	} else {
		c.ErrorJson(-100, "appId or secret verify error", JsonData{})
	}
}

func (c *NamesOverController) Post() {
	if c.verifyAppId() {
		page, _ := c.GetInt("page", 1)
		height := int(models.ApiBlocksTop())
		namesDb, e := models.FindNameOver(page, height)
		if e != nil {
			c.ErrorJson(-500, e.Error(), JsonData{})
		}
		var names []map[string]interface{}
		for i := 0; i < len(namesDb); i++ {
			var name = map[string]interface{}{}
			name["name"] = namesDb[i].Name
			name["length"] = namesDb[i].Length
			name["start_height"] = namesDb[i].StartHeight
			name["end_height"] = namesDb[i].EndHeight
			name["over_height"] = namesDb[i].OverHeight
			name["current_height"] = height
			name["owner"] = namesDb[i].Owner
			name["current_price"] =  utils.FormatTokens(namesDb[i].CurrentPrice)
			name["th_hash"] = namesDb[i].ThHash

			names = append(names, name)
		}
		if names == nil{
			c.SuccessJson([]JsonData{})
			return
		}
		c.SuccessJson(names)
	} else {
		c.ErrorJson(-100, "appId or secret verify error", JsonData{})
	}
}

func (c *NamesMyRegisterController) Post() {
	if c.verifyAppId() {
		signingKey := c.GetString("signingKey")
		page, _ := c.GetInt("page", 1)
		if signingKey == "" {
			c.ErrorJson(-500, "parameter is nul", JsonData{})
			return
		}
		aeasyAccount, e := models.SigningKeyHexStringAccount(signingKey)
		if e != nil {
			c.ErrorJson(-500, "Account signingKey error", e.Error())
			return
		}
		height := int(models.ApiBlocksTop())
		namesDb, e := models.FindNameMyRegister(aeasyAccount.Address, page, height)
		if e != nil {
			c.ErrorJson(-500, e.Error(), JsonData{})
		}
		var names []map[string]interface{}
		for i := 0; i < len(namesDb); i++ {
			var name = map[string]interface{}{}
			name["name"] = namesDb[i].Name
			name["length"] = namesDb[i].Length
			name["start_height"] = namesDb[i].StartHeight
			name["end_height"] = namesDb[i].EndHeight
			name["over_height"] = namesDb[i].OverHeight
			name["current_height"] = height
			name["owner"] = namesDb[i].Owner
			name["current_price"] =  utils.FormatTokens(namesDb[i].CurrentPrice)
			name["th_hash"] = namesDb[i].ThHash

			names = append(names, name)
		}
		if names == nil{
			c.SuccessJson([]JsonData{})
			return
		}
		c.SuccessJson(names)
	} else {
		c.ErrorJson(-100, "appId or secret verify error", JsonData{})
	}
}

func (c *NamesMyOverController) Post() {
	if c.verifyAppId() {
		signingKey := c.GetString("signingKey")
		page, _ := c.GetInt("page", 1)
		if signingKey == "" {
			c.ErrorJson(-500, "parameter is nul", JsonData{})
			return
		}
		aeasyAccount, e := models.SigningKeyHexStringAccount(signingKey)
		if e != nil {
			c.ErrorJson(-500, "Account signingKey error", e.Error())
			return
		}
		height := int(models.ApiBlocksTop())
		namesDb, e := models.FindNameMyRegisterIng(aeasyAccount.Address, page, height)
		if e != nil {
			c.ErrorJson(-500, e.Error(), JsonData{})
		}
		var names []map[string]interface{}
		for i := 0; i < len(namesDb); i++ {
			var name = map[string]interface{}{}
			name["name"] = namesDb[i].Name
			name["length"] = namesDb[i].Length
			name["start_height"] = namesDb[i].StartHeight
			name["end_height"] = namesDb[i].EndHeight
			name["over_height"] = namesDb[i].OverHeight
			name["current_height"] = height
			name["owner"] = namesDb[i].Owner
			name["current_price"] =  utils.FormatTokens(namesDb[i].CurrentPrice)
			name["th_hash"] = namesDb[i].ThHash

			names = append(names, name)
		}
		if names == nil{
			c.SuccessJson([]JsonData{})
			return
		}
		c.SuccessJson(names)
	} else {
		c.ErrorJson(-100, "appId or secret verify error", JsonData{})
	}
}

func (c *NamesUpdateController) Post() {
	if c.verifyAppId() {
		signingKey := c.GetString("signingKey")
		name := c.GetString("name")
		if signingKey == "" {
			c.ErrorJson(-500, "parameter is nul", JsonData{})
			return
		}
		aeasyAccount, e := models.SigningKeyHexStringAccount(signingKey)

		if e != nil {
			c.ErrorJson(-500, "Account signingKey error", e.Error())
			return
		}

		receipt, e := models.UpdateAENS(aeasyAccount, name)
		if e != nil {
			c.ErrorJson(-500, "parameter is nul", JsonData{})
		}
		c.SuccessJson(receipt)
	} else {
		c.ErrorJson(-100, "appId or secret verify error", JsonData{})
	}
}

func (c *NamesInfoController) Post() {
	if c.verifyAppId() {

		name := c.GetString("name")
		if name == "" {
			c.ErrorJson(-500, "parameter is nul", JsonData{})
			return
		}

		nameDb, err := models.FindNameName(name)
		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}

		if nameDb.CurrentPrice == 0 {
			c.ErrorJson(201, "name is no register", JsonData{})
			return
		}
		height := int(models.ApiBlocksTop())
		var nameMap = map[string]interface{}{}
		nameMap["name"] = nameDb.Name
		nameMap["length"] = nameDb.Length
		nameMap["start_height"] = nameDb.StartHeight
		nameMap["end_height"] = nameDb.EndHeight
		nameMap["over_height"] = nameDb.OverHeight
		nameMap["current_height"] = height
		nameMap["owner"] = nameDb.Owner

		nameMap["current_price"] = utils.FormatTokens(nameDb.CurrentPrice)
		nameMap["th_hash"] = nameDb.ThHash

		blocksDb, err := models.FindMicroBlockBlockNameorDatas(name)
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
			f, _ := mapObj["name_fee"].(json.Number).Float64()
			mapObj["name_fee"] = utils.FormatTokens(f)
			txs = append(txs, mapObj)
		}

		nameMap["claim"] = txs
		c.SuccessJson(nameMap)
	} else {
		c.ErrorJson(-100, "appId or secret verify error", JsonData{})
	}
}

func (c *NamesTransferController) Post() {
	if c.verifyAppId() {
		name := c.GetString("name")
		signingKey := c.GetString("signingKey")
		recipientAddress := c.GetString("recipientAddress")
		if name == "" || signingKey == "" {
			c.ErrorJson(-500, "parameter is nul", JsonData{})
			return
		}
		account, _ := models.SigningKeyHexStringAccount(signingKey)
		receipt, err := models.TransferAENS(account, recipientAddress, name)

		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}
		c.SuccessJson(receipt)
	} else {
		c.ErrorJson(-100, "appId or secret verify error", JsonData{})
	}
}

func (c *NamesAddController) Post() {
	if c.verifyAppId() {

		name := c.GetString("name")
		signingKey := c.GetString("signingKey")
		if name == "" || signingKey == "" {
			c.ErrorJson(-500, "parameter is nul", JsonData{})
			return
		}

		account, _ := models.SigningKeyHexStringAccount(signingKey)
		nameDb, err := models.FindNameName(name)
		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}
		aeHeight, _ := strconv.Atoi(strconv.FormatUint(models.ApiBlocksTop(), 10))
		if nameDb.CurrentPrice > 0 && nameDb.EndHeight < aeHeight {
			c.ErrorJson(-500, "names Has been registered ", JsonData{})
			return
		}

		var nameFee *big.Int
		if nameDb.CurrentPrice == 0 {
			nameFee = transactions.CalculateMinNameFee(name)
		} else {
			nameFee = decimal.NewFromFloat(nameDb.CurrentPrice).BigInt()
		}

		decimalValue := decimal.NewFromBigInt(nameFee, 0)
		decimalValueMul := decimalValue.Mul(decimal.NewFromFloat(0.05))
		decimalValue = decimalValue.Add(decimalValueMul)

		accountNet, e := models.ApiGetAccount(account.Address)
		if e != nil {
			c.ErrorJson(-500, e.Error(), JsonData{})
			return
		}
		tokens, err := strconv.ParseFloat(accountNet.Balance.String(), 64)
		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}
		f, _ := decimalValue.Float64()
		if tokens < f {
			c.ErrorJson(-500, "Lack of balance >"+utils.FormatTokens(f), JsonData{})
			return
		}

		receipt, err := models.ClaimAENS(account, name, decimalValue.BigInt())

		if err != nil {
			c.ErrorJson(-500, err.Error(), JsonData{})
			return
		}
		c.SuccessJson(receipt)
	} else {
		c.ErrorJson(-100, "appId or secret verify error", JsonData{})
	}
}
