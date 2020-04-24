package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Article struct {
	ID        int64  `orm:"column(id)" json:"id"`
	ArticleId int64  `orm:"column(article_id)" json:"article_id"`
	Image     string `orm:"column(image)" json:"image"`
	Title     string `orm:"column(title)" json:"title"`
	Date      string `orm:"column(date)" json:"date"`
	Content   string `orm:"column(content)" json:"content"`
	CTime     int64  `orm:"column(create_time)" json:"create_time"`
}

func InsertArticle(articleId int64, image string, title string, date string, content string) (Article, error) {
	unix := time.Now().UnixNano() / 1e6

	article := Article{
		ArticleId: articleId,
		Image:     image,
		Title:     title,
		Date:      date,
		Content:   content,
		CTime:     unix,
	}
	_, err := orm.NewOrm().Insert(&article)
	return article, err
}

func GetArticleAll(page int) ([]Article, error) {
	var articles []Article
	qs := orm.NewOrm().QueryTable("aea_article")
	_, err := qs.Limit(20).Offset((page - 1) * 20).OrderBy("-article_id").All(&articles)
	return articles, err
}

func FindArticleId(articleId string) (Article, error) {
	var article Article
	qs := orm.NewOrm().QueryTable("aea_article")
	err := qs.
		Filter("article_id", articleId).
		One(&article)
	return article, err
}

func (email *Article) TableName() string {
	return "aea_article"
}

func RegisterArticleDB() {
	orm.RegisterModel(new(Article))
}
