package main

import (
	"os"
	"unisun/api/advisor-listener/src"
	"unisun/api/advisor-listener/src/config"
	"unisun/api/advisor-listener/src/constants"
)

func main() {
	if os.Getenv(constants.NODE) != constants.PRODUCTION {
		config.ConfigENV()
	}
	r := src.App()
	port := os.Getenv(constants.PORT)
	if port == "" {
		r.Run(":8080")
	} else {
		r.Run(":" + port)
	}
}
