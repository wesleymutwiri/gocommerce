package models

import "github.com/jinzhu/gorm"

type Product struct {
		ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
		Name string `gorm:"size:255;not null;" json:"name"`
		Description string `gorm:"size:600;not null;" json:"description"`
		Price uint `json:"price""` 
}

func (product *Product) saveProduct(db *gorm.DB) (*Product, error){
	var err error
	err = db.Debug().Create(&product).Error
	if err != nil {
		return &Product{}, err
	}
	return &Product{}, nil
}