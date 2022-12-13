package services

import (
	"errors"
	"fmt"
	"github.com/octane77/rova/apiGateway/dtos"
	"github.com/octane77/rova/apiGateway/utils"
	"os"
)

type IdentityService interface {
	HandleRequest(callServiceApiDto dtos.CallServiceApiDto) (interface{}, error)
}

type identityService struct {
	allowedRoutes []string
}

func (i identityService) HandleRequest(callServiceApiDto dtos.CallServiceApiDto) (interface{}, error) {
	url := os.Getenv("IDENTITY_SERVICE_URL")
	isAllowedRoute := false
	if url == "" {
		return nil, errors.New("account service url not found")
	}
	for _, route := range i.allowedRoutes {
		if route == callServiceApiDto.Route {
			isAllowedRoute = true
		}
	}
	if isAllowedRoute {
		fmt.Println(callServiceApiDto.Path, "path", url+callServiceApiDto.Path, callServiceApiDto.Method)
		return utils.MakeApiRequest(url+callServiceApiDto.Path, callServiceApiDto.Body, callServiceApiDto.Method, &callServiceApiDto.Token)
	}
	return nil, errors.New("route Not Found")
}

func NewIdentityService() IdentityService {
	return identityService{
		allowedRoutes: []string{"create", "login", "profile"},
	}
}
