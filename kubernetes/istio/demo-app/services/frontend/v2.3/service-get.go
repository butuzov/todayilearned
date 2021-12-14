package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

func getDashboardEndpoint(svc DashboardService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.Get(request.([]string))
	}
}

func decodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	data := []string{}
	for k, v := range req.Header {
		data = append(data, fmt.Sprintf("%s: %s", k, strings.Join(v, "; ")))
	}

	sort.Strings(data)

	return data, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
