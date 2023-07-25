package gormdemo

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123@tcp(139.9.61.90:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print("db connect err")
	}

	db.AutoMigrate(&Product{})
	//db.Create(&Product{Code: "101", Price: 1000})
	//MyCreateBySQL(db)
}
