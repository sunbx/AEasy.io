package controllers

import (
	"ae/models"
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
		name["owner"] = namesDb[i].Owner
		name["current_price"] = namesDb[i].CurrentPrice
		name["th_hash"] = namesDb[i].ThHash

		names = append(names, name)
	}
	c.SuccessJson(names)
}

func (c *NamesPriceController) Post() {
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
		name["owner"] = namesDb[i].Owner
		name["current_price"] = namesDb[i].CurrentPrice
		name["th_hash"] = namesDb[i].ThHash

		names = append(names, name)
	}
	c.SuccessJson(names)
}

func (c *NamesOverController) Post() {
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
		name["owner"] = namesDb[i].Owner
		name["current_price"] = namesDb[i].CurrentPrice
		name["th_hash"] = namesDb[i].ThHash

		names = append(names, name)
	}
	c.SuccessJson(names)
}


func (c *NamesMyRegisterController) Get() {
	signingKey := c.GetString("signingKey")
	page, _ := c.GetInt("page", 1)
	if signingKey == "" {
		c.ErrorJson(-301, "parameter is nul", JsonData{})
		return
	}
	aeasyAccount, e := models.SigningKeyHexStringAccount(signingKey)
	if e != nil {
		c.ErrorJson(-301, "Account signingKey error", e.Error())
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
		name["owner"] = namesDb[i].Owner
		name["current_price"] = namesDb[i].CurrentPrice
		name["th_hash"] = namesDb[i].ThHash

		names = append(names, name)
	}
	c.SuccessJson(names)
}


func (c *NamesMyOverController) Get() {
	signingKey := c.GetString("signingKey")
	page, _ := c.GetInt("page", 1)
	if signingKey == "" {
		c.ErrorJson(-301, "parameter is nul", JsonData{})
		return
	}
	aeasyAccount, e := models.SigningKeyHexStringAccount(signingKey)
	if e != nil {
		c.ErrorJson(-301, "Account signingKey error", e.Error())
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
		name["owner"] = namesDb[i].Owner
		name["current_price"] = namesDb[i].CurrentPrice
		name["th_hash"] = namesDb[i].ThHash

		names = append(names, name)
	}
	c.SuccessJson(names)
}
