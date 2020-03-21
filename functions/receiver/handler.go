package function

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

func Handle(req handler.Request) (handler.Response, error) {

	log.Printf("Received: %q", string(req.Body))

	return handler.Response{
		Body:       []byte(fmt.Sprintf("Received: %q", string(req.Body))),
		StatusCode: http.StatusOK,
	}, nil

}
