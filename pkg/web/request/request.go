package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gmalheiro/playground-middleware/pkg/web/response"
)

const (
	ContentTypeTextPlain             = "text/plain"
	ContentTypeApplicationJSON       = "application/json"
	ContentTypeMultipartFormData     = "multipart/form-data"
	ContentTypeApplicationUrlEncoded = "application/x-www-form-urlencoded"
)

type Request struct {
	request *http.Request
}

func From(httpr *http.Request) *Request {
	return &Request{
		request: httpr,
	}
}

func (r *Request) ToValidContentType(contentType string) error {
	if r.request.Header.Get("Content-Type") != contentType {
		s := fmt.Sprintf("request content-type is not %s", contentType)
		return errors.New(s)
	}

	return nil
}

func (r *Request) ParseJSON(data any) error {
	if err := r.ToValidContentType(ContentTypeApplicationJSON); err != nil {
		return err
	}

	if err := json.NewDecoder(r.request.Body).Decode(data); err != nil {
		return err
	}

	return nil
}

func (r *Request) ParseValidJSON(data any) ([]response.Validation, error) {
	if err := r.ParseJSON(data); err != nil {
		return nil, err
	}

	if err := validateJSON(data); err != nil {
		return err, nil
	}

	return nil, nil
}
