package main

import (
	"user-access/config"
	"user-access/router"
)

func main() {
	config.Database()
	router.StartRouter()
}
