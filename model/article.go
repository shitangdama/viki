package model

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"

	"viki/common"
)

type Article struct {
	gorm.Model

	Title   string         `gorm:"not null"`
	Content string         `gorm:"type:text"`
	Source  string         `gorm:"type:text"`
	Tags    pq.StringArray `gorm:"type:varchar(100)[]"`
	Info    postgres.Jsonb
}

// GetArticleTotal gets the total number of articles based on the constraints
func GetArticleTotal(maps interface{}) (int, error) {
	var count int
	db := common.GetDB()
	if err := db.Model(&Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetArticles gets a list of articles based on paging constraints
func GetArticles(pageNum int, pageSize int, maps interface{}) ([]*Article, error) {
	var articles []*Article
	db := common.GetDB()
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articles, nil
}

func GetArticle(id int) (*Article, error) {
	var article Article
	db := common.GetDB()
	err := db.First(&article, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}


func CreateArticle(data interface{}) (*Article, error) {
	var article Article
	db := common.GetDB()
	db.Create(data).Find(&article)

	return &article, nil
}

func UpdataArticle(id int, data interface{}) (*Article, error) {
	var article Article
	db := common.GetDB()

	if err  := db.First(&article, id).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	db.Model(&article).Updates(data).Find(&article)
	return &article, nil
}
// func (article *Article) UpdateArticle() (*Article, error) {


// }