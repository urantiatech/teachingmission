package views

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	api "github.com/urantiatech/teachingmission/apigateways/webgw/api"
)

func encodeRequest(ctx context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func decodePageResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var buf api.Page
	if err := json.NewDecoder(r.Body).Decode(&buf); err != nil {
		return nil, err
	}
	return buf, nil
}
