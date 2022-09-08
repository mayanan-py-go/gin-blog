package models

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag Tag `json:"tag"`

	Title string `json:"title" gorm:"index"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}
func (a *Article) BeforeCreate(tx *gorm.DB) error {
	a.CreatedOn = time.Now().Unix()
	return nil
}
func (a *Article) AfterUpdate(tx *gorm.DB) error {
	tx.Model(&Article{}).Where("id = ?", a.ID).UpdateColumn("modified_on", time.Now().Unix())
	return nil
}
func ExistArticleByID(id int) bool {
	var art Article
	db.Take(&art, id)
	return art.ID > 0
}
func GetArticleTotal(maps map[string]any) (total int64) {
	db.Model(&Article{}).Where(maps).Count(&total)
	return total
}
func GetArticles(pageNum int, pageSize int, maps any) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}
func GetArticle(id int) (article Article) {
	db.Preload("Tag").Take(&article, id)
	return
}
func EditArticle(id int, data map[string]any) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)
	return true
}
func AddArticle(data map[string]any) bool {
	db.Create(&Article{
		TagID: data["tag_id"].(int),
		Title: data["title"].(string),
		Desc: data["desc"].(string),
		Content: data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State: data["state"].(int),
	})
	return true
}
func DeleteArticle(id int) bool {
	db.Delete(&Article{}, id)
	return true
}
