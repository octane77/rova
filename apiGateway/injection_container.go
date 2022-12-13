package main

import (
	"github.com/octane77/rova/apiGateway/controllers"
	"github.com/octane77/rova/apiGateway/routers"
	"github.com/octane77/rova/apiGateway/services"
	"go.uber.org/dig"
)

func BuildApplication() (*ApplicationServer, error) {
	c := dig.New()
	serviceConstructors := []interface{}{

		// CONTROLLERS
		controllers.NewApiServiceController,
		// SERVICES
		services.NewAccountService,
		services.NewIdentityService,
		// REPOSITORIES
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
