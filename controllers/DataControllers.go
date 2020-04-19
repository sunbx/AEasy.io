package controllers

import (
	"ae/models"
	"ae/utils"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

type ArticleDataController struct {
	BaseController
}
type WealthDataController struct {
	BaseController
}
type NameshDataController struct {
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
func (c *WealthDataController) Get() {
	response := utils.Get("https://www.aeknow.org/api/wealth500")
	var wealthData WealthData
	err := json.Unmarshal([]byte(response), &wealthData)
	if err != nil {
		c.ErrorJson(500, err.Error(), JsonData{})
		return
	}
	models.DeleteAllTop()
	models.InsertTop(wealthData.Top500)
	c.SuccessJson(JsonData{})
}

func (c *NameshDataController) Get() {
	response := utils.Get("https://mainnet.aeternal.io/middleware/names")
	var names []models.Names
	err := json.Unmarshal([]byte(response), &names)
	if err != nil {
		c.ErrorJson(500, err.Error(), JsonData{})
		return
	}
	models.DeleteAllNames()
	models.InsertNames(names)
	c.SuccessJson(JsonData{})
}




func download(url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	//函数结束后关闭相关链接
	defer resp.Body.Close()

	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("NewDocumentFromReader error", err)
		return
	}

	dom.Find("tr").Each(func(i int, selection *goquery.Selection) {
		val, _ := selection.Find("a").Attr("href")
		var address string
		var weatlth string
		var percentage string
		var time string
		selection.Find("td").Each(func(i int, selection *goquery.Selection) {
			//fmt.Println("i", i, "select ", selection.Text())
			if i == 1 {
				address = selection.Text()
			}
			if i == 2 {
				weatlth = selection.Text()
			}
			if i == 3 {
				percentage = strings.Replace(selection.Text(), " ", "", -1)
			}
			if i == 4 {
				time = selection.Text()
			}

		})
		fmt.Println("i", i, "select ", val, weatlth, percentage, time)
		//fmt.Println("i", i, "select ", val)
		//val3 := selection.Find("td").Next().Next().Text()
		//val4 := selection.Find("td").Next().Next().Next().Text()
		//fmt.Println("i", i, "select " ,text)
	})

	//c.SuccessJson(JsonData{})

	//links := collectlinks.All(resp.Body)
	//for _, link := range links {
	//	fmt.Println("parse url", link)
	//}
}
