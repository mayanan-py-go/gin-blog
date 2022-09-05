package models

import (
	"fmt"
	"gin_log/pkg/setting"
	"gorm.io/gorm"
	"log"

	"gorm.io/driver/mysql"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}
func init() {
	var (
		err error
		dbName, user, password, host string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "fail to GetSection 'database': %v", err)
	}

	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, dbName,
		)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	if err = db.AutoMigrate(&Model{}); err != nil {
		log.Println(err)
	}
}
func CloseDB() {
	sqlDb, err := db.DB()
	if err != nil {
		log.Println(err)
		return
	}
	if err = sqlDb.Close(); err != nil {
		fmt.Println(err)
	}
}
