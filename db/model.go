package db

import "gorm.io/gorm"

type (
	Pajak struct {
		gorm.Model
		ID   uint `gorm:"primaryKey"`
		Nama string
		Rate float64
	}

	Item struct {
		gorm.Model
		ID   uint `gorm:"primaryKey"`
		Nama string
	}

	PajakItem struct {
		gorm.Model
		ID      uint `gorm:"primaryKey"`
		IdItem  int
		IdPajak int
	}
)
