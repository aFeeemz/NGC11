package main

import (
	"NGC11/initializers"

	"github.com/labstack/echo/v4"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {
	e := echo.New()
}
