package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

type getBookRequest struct {
	Id string
}

func getAuthorEndpoint(svc BooksService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getBookRequest)

		v, err := svc.GetBook(req.Id)
		if err != nil {
			var a Book
			return a, err
		}

		return v, nil
	}
}

func decodeGetBook(ctx context.Context, req *http.Request) (interface{}, error) {
	var request getBookRequest
	vars := mux.Vars(req)
	request.Id = vars["id"]

	return request, nil
}

func encodeGetBook(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
