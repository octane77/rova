package services

import (
	"errors"
	"fmt"
	"github.com/octane77/rova/apiGateway/dtos"
	"github.com/octane77/rova/apiGateway/utils"
	"os"
)

type AccountService interface {
	HandleRequest(callServiceApiDto dtos.CallServiceApiDto) (interface{}, error)
}

type accountService struct {
	allowedRoutes []string
}

func (a accountService) HandleRequest(callServiceApiDto dtos.CallServiceApiDto) (interface{}, error) {
	url := os.Getenv("ACCOUNT_SERVICE_URL")
	isAllowedRoute := false
	if url == "" {
		return nil, errors.New("account service url not found")
	}
	for _, route := range a.allowedRoutes {
		if route == callServiceApiDto.Route {
			isAllowedRoute = true
		}
	}
	if isAllowedRoute {
		fmt.Println(callServiceApiDto.Path, "path", url+callServiceApiDto.Path, callServiceApiDto.Method)
		return utils.MakeApiRequest(url+callServiceApiDto.Path, callServiceApiDto.Body, callServiceApiDto.Method, nil)
	}
	return nil, errors.New("route Not Found")
}

func NewAccountService() AccountService {
	return accountService{
		allowedRoutes: []string{"create"},
	}
}
