package controllers

import (
	"ae/models"
	"ae/utils"
	"encoding/json"
)

type WalletListController struct {
	BaseController
}
type BaseDataController struct {
	BaseController
}

func (c *WalletListController) Post() {
	addresses, e := models.FindAddressBalanceTopList()
	if e != nil {

	}

	var walletList []map[string]interface{}
	for i := 0; i < len(addresses); i++ {
		var balanceStr = utils.FormatTokens(addresses[i].Balance)
		addresses[i].BalanceStr = balanceStr
		var wallet = map[string]interface{}{}
		wallet["address"] = addresses[i].Address
		wallet["balance"] = balanceStr
		wallet["update_time"] = addresses[i].UpdateTime
		wallet["percentage"] = utils.FormatTokensP(addresses[i].Balance/355005806*100, 2)

		walletList = append(walletList, wallet)
	}

	c.SuccessJson(walletList)
}

type AeKnowAPI struct {
	PriceUsdt         string `json:"price_usdt"`
	PriceBtc          string `json:"price_btc"`
	BlockHeight       string `json:"block_height"`
	TotalTransactions string `json:"total_transactions"`
	MaxTps            string `json:"max_tps"`
	MarketCap         string `json:"market_cap"`
	TotalCoins        string `json:"total_coins"`
	AensTotal         string `json:"aens_total"`
	OraclesTotal      string `json:"oracles_total"`
	ContractsTotal    string `json:"contracts_total"`
	AccountsTotal     string `json:"accounts_total"`
}

func (c *BaseDataController) Post() {
	response := utils.Get("https://www.aeknow.org/api")
	var knowApi AeKnowAPI
	err := json.Unmarshal([]byte(response), &knowApi)
	if err != nil {
		c.ErrorJson(500, "Umarshal failed", JsonData{})
		return
	}
	c.SuccessJson(knowApi)
}
