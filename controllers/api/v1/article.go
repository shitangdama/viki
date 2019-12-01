package v1

import (
	"fmt"
	"net/http"
	"viki/model"
	"github.com/unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

type Article struct {
	Title       string    `form:"title" json:"title" binding:"exists,min=4"`
	Content     string    `form:"content" json:"content"`
	Source      string    `form:"source" json:"source"`
	Tags        pq.StringArray  `form:"tags" json:"tags"`
	Info        postgres.Jsonb `form:"info" json:"info"`
}

//获取多个文章标签
func GetArticles(c *gin.Context) {

	page := com.StrTo(c.DefaultQuery("page", "0")).MustInt()

	articles, _ := model.GetArticles(page,20,&model.Article{})

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": articles,
	})
}

// // 获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	fmt.Println(id)

	if article, err := model.GetArticle(id); err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "fail",
			"data": err,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": article,
		})
	}
}

func CreateArticle(c *gin.Context) {
	var article Article

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(article)
	articleData, _ := model.CreateArticle(article)

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": articleData,
	})
}

func UpdataArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	var article Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	articleData, _ := model.UpdataArticle(id, article)

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": articleData,
	})
}
