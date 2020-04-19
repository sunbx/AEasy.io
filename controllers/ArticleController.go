package controllers

import "ae/models"

type ArticleListController struct {
	BaseController
}

func (c *ArticleListController) Get() {
	page, _ := c.GetInt("page", 0)
	articles, err := models.GetArticleAll(page)
	if err != nil {
		c.ErrorJson(-200, err.Error(), JsonData{})
		return
	}
	c.SuccessJson(articles)
}
