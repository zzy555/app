package gormdemo

import (
	"fmt"
        "time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    uint         `gorm:"primaryKey;autoIncrement"`
	Code  string       `gorm:"type:varchar(8)"`
	Price uint         `gorm:"column:prices"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdateTime  time.Time `gorm:"autoUpdateTime"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Print("before create:",p)
	return nil

}

func Gorm() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123@tcp(139.9.61.90:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print("db connect err")
	}

	db.AutoMigrate(&Product{})
	//db.Create(&Product{Code: "101", Price: 1000})

	myCreate(db)
}

func myCreate(db *gorm.DB) {
	//var product = []Product{{Code:"110", Price: 666}, {Code:"111",Price: 777}}
	//db.CreateInBatches(product,5)

	//db.Model(&Product{}).Create(map[string]interface{}{
	//	"Code": "121", "Price": 11,
	//})
	
	db.Select("Code","Price").Create(&Product{Code:"131", Price: 12})

}
