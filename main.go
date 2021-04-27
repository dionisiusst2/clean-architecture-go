package main

import (
	"github.com/dionisiusst2/bakery-id/config"
	"github.com/dionisiusst2/bakery-id/infrastructure/database/psql"
	"github.com/dionisiusst2/bakery-id/infrastructure/router"
	"github.com/dionisiusst2/bakery-id/registry"
)

func main() {
	config.LoadEnv()

	db := psql.InitDB()
	psql.Migrate(db)

	registry := registry.NewRegistry(db)
	router := router.InitRouter(registry.NewAppController(), registry.NewAppMiddleware())

	router.Run(":8000")
}
