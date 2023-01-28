package main

import (
	"user-access/pkg/database"
	"user-access/pkg/routes"
)

func main() {
	database.Database()

	routes.HandleRequests()
}
