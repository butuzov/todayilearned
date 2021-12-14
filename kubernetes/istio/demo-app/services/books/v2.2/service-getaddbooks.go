package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

type GetAllBooksRequest struct {
	Sorted bool
}

func getAllBooksEndpoint(svc BooksService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllBooksRequest)

		return svc.GetAllBooks(req.Sorted), nil
	}
}

func decodeGetAllBooks(_ context.Context, req *http.Request) (interface{}, error) {
	var res GetAllBooksRequest
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

func encodeGetAllBooks(_ context.Context, w http.ResponseWriter, res interface{}) error {
	return json.NewEncoder(w).Encode(res)
}
