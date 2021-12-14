package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func getAllBooksEndpoint(svc BooksService) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return svc.GetAllBooks(), nil
	}
}

func decodeGetAllBooks(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeGetAllBooks(_ context.Context, w http.ResponseWriter, res interface{}) error {
	return json.NewEncoder(w).Encode(res)
}
