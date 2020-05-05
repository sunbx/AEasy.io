package controllers

import (
	"ae/models"
	"ae/utils"
	"encoding/json"
	"fmt"
	"strconv"
)

type ArticleDataController struct {
	BaseController
}


type Article struct {
	Id            int64    `json:"id"`
	Title         Rendered `json:"title"`
	Content       Rendered `json:"content"`
	FeaturedMedia int      `json:"featured_media"`
	Date          string   `json:"date"`
}
type Rendered struct {
	Rendered string `json:"rendered"`
}
type Media struct {
	SourceUrl string `json:"source_url"`
}

type WealthData struct {
	Top500 []models.Top500 `json:"top500"`
}


func (c *ArticleDataController) Get() {
	response := utils.Get("https://www.aechina.io/wp-json/wp/v2/posts?page=1")
	var articles []Article
	err := json.Unmarshal([]byte(response), &articles)
	if err != nil {
		c.ErrorJson(500, "Umarshal failed", JsonData{})
		return
	}
	for i := 0; i < len(articles); i++ {
		response = utils.Get("https://www.aechina.io/wp-json/wp/v2/media/" + strconv.Itoa(articles[i].FeaturedMedia))
		var media Media
		err = json.Unmarshal([]byte(response), &media)
		if err != nil {
			c.ErrorJson(500, "Umarshal failed", JsonData{})
			return
		}
		fmt.Println(media)
		if media.SourceUrl != "" {
			_, _ = models.InsertArticle(articles[i].Id, media.SourceUrl, articles[i].Title.Rendered, articles[i].Date, articles[i].Content.Rendered)
		}

	}
	c.SuccessJson(JsonData{})
}


