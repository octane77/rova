package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/octane77/rova/identityService/dtos"
	"github.com/octane77/rova/identityService/entities"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	API_REQ_METHOD_POST = "post"
	API_REQ_METHOD_GET  = "get"
)

type APIService interface {
	GetAccountDetails(account *entities.User) (*interface{}, error)
}

type apiService struct {
}

func (a apiService) GetAccountDetails(user *entities.User) (*interface{}, error) {
	url := os.Getenv("ACCOUNT_SERVICE_URL")
	if url == "" {
		return nil, errors.New("account service url not found")
	}
	body := dtos.GetAccountDetailsDto{
		CustomerId: user.ID,
	}
	return a.makeApiRequest(url+"/get-all-by-customer-id", body, API_REQ_METHOD_POST)
}

func (a apiService) makeApiRequest(url string, body interface{}, method string) (*interface{}, error) {
	var (
		reqBody  []byte
		err      error
		response *http.Response
		data     *interface{}
	)
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}
	switch method {
	case API_REQ_METHOD_POST:
		response, err = http.Post(url, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			return nil, err
		}
		break
	case API_REQ_METHOD_GET:
		response, err = http.Post(url, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			return nil, err
		}
		break

	default:
		return nil, errors.New("method not found")
	}
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}

	response.Body.Close()
	return data, nil
}

func NewAPIService() APIService {
	return apiService{}
}
