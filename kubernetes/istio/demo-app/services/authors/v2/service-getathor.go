package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

type getAuthorRequest struct {
	Id string
}

func getAuthorEndpoint(svc AuthorsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAuthorRequest)

		v, err := svc.GetAuthor(req.Id)
		if err != nil {
			var a Author
			return a, err
		}
		return v, nil
	}
}

func decodeGetAuthor(ctx context.Context, req *http.Request) (interface{}, error) {
	var request getAuthorRequest
	vars := mux.Vars(req)
	request.Id = vars["id"]

	return request, nil
}

func encodeGetAuthor(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
