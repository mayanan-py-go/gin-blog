package models

import (
	"fmt"
	"gin_log/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int64 `json:"created_on" gorm:"autoCreateTime"`
	ModifiedOn int64 `json:"modified_on" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
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

	//if err = db.AutoMigrate(&Model{}); err != nil {
	//	log.Println(err)
	//}
	if err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Tag{}); err != nil {
		log.Println(err)
	}
	// 可以通过Set设置附加参数，下面设置表的存储引擎为InnoDB
	if err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Article{}); err != nil {
		log.Println(err)
	}
	if err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Auth{}); err != nil {
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
