package main

import (
	"fmt"
	"project/e-commerce/config"
	"project/e-commerce/factory"
	"project/e-commerce/migration"
	"project/e-commerce/utils/database/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDBmySql(cfg)

	migration.InitMigrate(db)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}
