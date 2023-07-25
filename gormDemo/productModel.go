package gormdemo

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Code        string    `gorm:"type:varchar(8)"`
	Price       uint      `gorm:"column:prices"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdateTime  time.Time `gorm:"autoUpdateTime"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Print("before create:", p, "\n")
	return nil
}
