package gormdemo

import (
	"fmt"

	"gorm.io/gorm"
)

func MySelect(db *gorm.DB) {
	var product1 Product
	var product2 Product
	var product3 Product

	db.First(&product1)
	fmt.Print(product1, "\n")
	db.Take(&product2)
	fmt.Print(product2, "\n")
	db.Last(&product3)
	fmt.Print(product3, "\n")

	var product4 Product = Product{ID: 3}
	db.First(&product4)
	fmt.Print(product4, "\n")

	var products1 []Product
	db.Where("id > ?", "2").Find(&products1)
	fmt.Print(products1, "\n")

	// var products2 []Product
	// db.Offset(2).Find(&products2)
	// fmt.Print(products2, "\n")

}
