package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type getBookRequest struct {
	Id string
}

func getDashboardEndpoint(svc DashboardService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.Get()
	}
}

func decodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
