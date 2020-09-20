package main

import (
	pagination "github.com/evrintobing17/go-postgres-pagination-example/"

	"github.com/evrintobing17/go-postgres-pagination-example/offsetlimit"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	db := pagination.InitDB()
	fetchHandler := offsetlimit.FetchHandler(db)
	e := echo.New()
	e.GET("/payments", fetchHandler)
	logrus.Error(e.Start(":9090"))
}
