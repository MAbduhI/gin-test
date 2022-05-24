package main

import (
	"github.com/MAbduhI/gin-test/config"
	api "github.com/MAbduhI/gin-test/src"
)

func main() {
	config := config.Init()
	api.InitApi(&config)
}
