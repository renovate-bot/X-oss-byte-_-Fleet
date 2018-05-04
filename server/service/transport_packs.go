package service

import (
	"context"
	"encoding/json"
	"net/http"
)

func decodeDeletePackRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	id, err := idFromRequest(r, "id")
	if err != nil {
		return nil, err
	}
	var req deletePackRequest
	req.ID = id
	return req, nil
}

func decodeGetPackRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	id, err := idFromRequest(r, "id")
	if err != nil {
		return nil, err
	}
	var req getPackRequest
	req.ID = id
	return req, nil
}

func decodeListPacksRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	opt, err := listOptionsFromRequest(r)
	if err != nil {
		return nil, err
	}
	return listPacksRequest{ListOptions: opt}, nil
}

func decodeApplyPackSpecsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req applyPackSpecsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil

}
