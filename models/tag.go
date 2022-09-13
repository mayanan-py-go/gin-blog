package models

import (
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	Model

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}
func (t *Tag) BeforeCreate(tx *gorm.DB) error {
	t.CreatedOn = time.Now().Unix()
	return nil
}
func (t *Tag) AfterUpdate(tx *gorm.DB) error {
	// UpdateColumn会忽略钩子
	tx.Model(&Tag{}).Where("id = ?", t.ID).UpdateColumn("modified_on", time.Now().Unix())
	return nil
}
func GetTags(pageNum, pageSize int, maps any) []Tag {
	tags := make([]Tag, 0)
	db.Where(maps).Limit(pageSize).Offset(pageNum).Find(&tags)
	return tags
}
func GetTagTotal(maps any) (count int64) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}
func ExitTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).Take(&tag)
	return tag.ID > 0
}
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name: name,
		State: state,
		CreatedBy: createdBy,
	})
	return true
}
func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).Take(&tag)
	return tag.ID > 0
}
func EditTag(id int, data any) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}
func DeleteTag(id int) bool {
	db.Delete(&Tag{}, id)
	return true
}
func CleanAllTag() bool {
	db.Unscoped().Where("deleted_at != ?", 0).Delete(&Tag{})
	return true
}
