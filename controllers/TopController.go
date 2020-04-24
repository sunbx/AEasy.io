package controllers

import (
	"ae/models"
	"ae/utils"
	"encoding/json"
)

type WealthListController struct {
	BaseController
}
type BaseDataController struct {
	BaseController
}

func (c *WealthListController) Get() {
	page, _ := c.GetInt("page", 0)
	tops, err := models.GetTopAll(page)
	if err != nil {
		c.ErrorJson(-200, err.Error(), JsonData{})
		return
	}
	var topJsons []models.Top500Json
	for i := 0; i < len(tops); i++ {
		content := utils.FormatTokensP(tops[i].Balance, 5)
		//tops[i].Tokens = content + " AE"
		var top models.Top500Json
		top.Id = tops[i].Id
		top.Balance = content + " AE"
		top.Ak = tops[i].Ak
		top.Per = tops[i].Per
		top.LastUpdate = tops[i].LastUpdate

		topJsons = append(topJsons, top)

	}
	c.SuccessJson(topJsons)
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
