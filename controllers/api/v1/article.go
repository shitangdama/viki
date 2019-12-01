package v1

import (
	"errors"
	"net/http"
	"viki/model"
	"viki/pkg/logging"
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

// ShowAccount godoc
// @Summary 新增文章
// @Description 新增文章
func GetArticles(c *gin.Context) {

	page := com.StrTo(c.DefaultQuery("page", "0")).MustInt()

	articles, _ := model.GetArticles(page,20,&model.Article{})

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": articles,
	})
}

// ShowAccount godoc
// @Summary 新增文章
// @Description 新增文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()


	if article, err := model.GetArticle(id); err != nil {
		logging.Info(err)
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "fail",
			"data": err,
		})
	} else {
		logging.Info(err)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": article,
		})
	}
}

// ShowAccount godoc
// @Summary 新增文章
// @Description 新增文章
// @Accept  json
// @Produce  json
// @Param id path int true "Article ID"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/aticles [post]
func CreateArticle(c *gin.Context) {
	var article Article

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("Invalid param")})
		return
	}

	if articleData, err := model.CreateArticle(article); err != nil {
		c.JSON(200, gin.H{
			"code": 201,
			"msg":  "ok",
			"data": articleData,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "success",
			"data": err,
		})
	}


}

// ShowAccount godoc
// @Summary 修改文章
// @Description 修改文章
// @Accept  json
// @Produce  json
// @Param id path int true "Article ID"
// @Param tags body []string false "Tags"
// @Param title body string true "Title"
// @Param source body string false "Source"
// @Param content body string true "Content"
// @Param info body map[string]interface{} true "Info"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/aticles/:id [put]
func UpdataArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	var article Article
	if err := c.ShouldBindJSON(&article); err != nil {
		logging.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("Invalid param")})
		return
	}

	articleData, _ := model.UpdataArticle(id, article)

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": articleData,
	})
}
