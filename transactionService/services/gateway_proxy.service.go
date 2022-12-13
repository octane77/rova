package services

type GatewayProxyService interface {
}

type gatewayProxyService struct {
}

func NewGatewayProxyService() GatewayProxyService {
	return gatewayProxyService{}
}
