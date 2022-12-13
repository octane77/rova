package main

import (
	"github.com/octane77/rova/transactionService/configs"
	"github.com/octane77/rova/transactionService/controllers"
	"github.com/octane77/rova/transactionService/repositories"
	"github.com/octane77/rova/transactionService/routers"
	"github.com/octane77/rova/transactionService/services"
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
		services.NewTransactionService,
		services.NewGatewayProxyService,
		// REPOSITORIES
		repositories.NewTransactionRepository,
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
