package main

import (
	"github.com/dionisiusst2/clean-architecture-go/config"
	"github.com/dionisiusst2/clean-architecture-go/infrastructure/database/psql"
	"github.com/dionisiusst2/clean-architecture-go/infrastructure/router"
	"github.com/dionisiusst2/clean-architecture-go/registry"
)

func main() {
	config.LoadEnv()

	db := psql.InitDB()
	psql.Migrate(db)

	registry := registry.NewRegistry(db)
	router := router.InitRouter(registry.NewAppController(), registry.NewAppMiddleware())

	router.Run(":8000")
}
