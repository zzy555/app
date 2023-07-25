package gormdemo

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func MyCreate(db *gorm.DB) {
	var product = []Product{{Code: "110", Price: 666}, {Code: "111", Price: 777}}
	db.CreateInBatches(product, 5)

	db.Model(&Product{}).Create(map[string]interface{}{
		"Code": "121", "Price": 11,
	})

	db.Select("Code", "Price").Create(&Product{Code: "131", Price: 12})

}

func MyCreateBySQL(db *gorm.DB) {
	db.Model(Product{}).Create(map[string]interface{}{
		"Code":  "501",
		"Price": clause.Expr{SQL: "?", Vars: []interface{}{555}},
	})
}
