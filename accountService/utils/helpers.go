package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	API_REQ_METHOD_POST = "post"
	API_REQ_METHOD_GET  = "get"
)

func MakeApiRequest(url string, body interface{}, method string) (*interface{}, error) {
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
		break
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
