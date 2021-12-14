package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func getAllAuthorsEndpoint(svc AuthorsService) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return GetAllAuthorsResponse(svc.GetAllAuthors()), nil
	}
}

func decodeGetAllAuthorsRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeGetAllAuthorsResponse(_ context.Context, w http.ResponseWriter, res interface{}) error {
	return json.NewEncoder(w).Encode(res)
}

// For each method, we define request and response structs
type GetAllAuthorsRequest struct{}

type GetAllAuthorsResponse []Author
