package frame

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var UserDB *gorm.DB

func Initialize() {

	UserDB = initPostgreSQL()

}

func initPostgreSQL() *gorm.DB {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=lxy666217 dbname=user port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("数据库链接失败:" + err.Error())
	}

	return db
}
