package main

import (
	"github.com/octane77/rova/accountService/configs"
	"github.com/octane77/rova/accountService/controllers"
	"github.com/octane77/rova/accountService/repositories"
	"github.com/octane77/rova/accountService/routers"
	"github.com/octane77/rova/accountService/services"
	"go.uber.org/dig"
)

func BuildApplication() (*ApplicationServer, error) {
	c := dig.New()
	serviceConstructors := []interface{}{
		// DATABASE
		configs.SetUpDatabaseConnection,
		// CONTROLLERS
		controllers.NewAccountController,
		// SERVICES
		services.NewAccountService,
		services.NewGatewayProxyService,
		// REPOSITORIES
		repositories.NewAccountRepository,
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
