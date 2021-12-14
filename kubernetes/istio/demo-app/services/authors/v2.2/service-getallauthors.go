package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

// For each method, we define request and response structs
type GetAllAuthorsRequest struct {
	Sorted bool
}

type GetAllAuthorsResponse []Author

func getAllAuthorsEndpoint(svc AuthorsService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllAuthorsRequest)

		return GetAllAuthorsResponse(svc.GetAllAuthors(req.Sorted)), nil
	}
}

func decodeGetAllAuthorsRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	var res GetAllAuthorsRequest
	for k := range req.Header {
		if !strings.EqualFold(k, "ff") {
			continue
		}

		if res.Sorted {
			continue
		}

		for i := range req.Header[k] {
			if req.Header[k][i] == "sort" {
				res.Sorted = true
			}
		}

	}

	return res, nil
}

func encodeGetAllAuthorsResponse(_ context.Context, w http.ResponseWriter, res interface{}) error {
	return json.NewEncoder(w).Encode(res)
}
