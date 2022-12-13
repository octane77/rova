package main

import (
	"github.com/octane77/rova/identityService/configs"
	"github.com/octane77/rova/identityService/controllers"
	"github.com/octane77/rova/identityService/repositories"
	"github.com/octane77/rova/identityService/routers"
	"github.com/octane77/rova/identityService/services"
	"go.uber.org/dig"
)

func BuildApplication() (*ApplicationServer, error) {
	c := dig.New()
	serviceConstructors := []interface{}{
		// DATABASE
		configs.SetUpDatabaseConnection,
		// CONTROLLERS
		controllers.NewUserController,
		// SERVICES
		services.NewUserService,
		services.NewAPIService,
		services.NewJWTService,
		// REPOSITORIES
		repositories.NewUserRepository,
		// ROUTER
		routers.NewRouter,
		// APPLICATION SERVER
		NewApplicationServer,
	}

	for _, service := range serviceConstructors {
		if err := c.Provide(service); err != nil {
			return nil, err
		}
	}

	//GET SERVER INSTANCE FROM CONTAINER
	var app *ApplicationServer
	err := c.Invoke(func(a *ApplicationServer) {
		app = a
	})
	return app, err
}
