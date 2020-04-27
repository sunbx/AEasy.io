package controllers

import "ae/models"

type ArticleListController struct {
	BaseController
}

func (c *ArticleListController) Post() {
	page, _ := c.GetInt("page", 1)
	articles, err := models.GetArticleAll(page)
	if err != nil {
		c.ErrorJson(-200, err.Error(), JsonData{})
		return
	}

	for i := 0; i < len(articles); i++ {
		articles[i].Content = ""
	}
	c.SuccessJson(articles)
}




