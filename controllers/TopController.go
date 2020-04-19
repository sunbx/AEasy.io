package controllers

import (
	"ae/models"
	"ae/utils"
)


type WealthListController struct {
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
