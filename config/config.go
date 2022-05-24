package config

import (
	"fmt"

	"github.com/MAbduhI/gin-test/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB *gorm.DB
}

func Init() Config {
	c := Config{}
	c.initGorm()
	return c
}
func (c *Config) initDB() {

	return
}

func (c *Config) initGorm() {
	var err error
	dsn := "phpmyadmin:Almarkaz~123@tcp(127.0.0.1:3306)/tesgin?charset=utf8mb4&parseTime=True&loc=Local"
	c.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to connect database")
		return
	}
	c.DB.AutoMigrate(&db.Pajak{}, &db.Item{}, &db.PajakItem{})
	return
}
