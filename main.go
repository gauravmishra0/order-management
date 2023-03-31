package main

import (
	_ "github.com/go-sql-driver/mysql"
	"order-management-service/api"
)

func main() {
	api.Run()
}
