package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Windowsfreak/go-mc/http/contextkeys"
	"github.com/Windowsfreak/go-mc/http/header"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	ErrInvalidContentType  = errors.New("invalid content-type")
	ErrInvalidAcceptHeader = errors.New("invalid accept header")
)

// DecodeRequest decodes the request based on the content-type
func DecodeRequest(ctx context.Context, r *http.Request, request interface{}) error {
	if r.Body == nil {
		return nil
	}

	return decodeRequest(ctx, r.Body, request)
}

func decodeRequest(ctx context.Context, body io.ReadCloser, request interface{}) error {
	defer func() {
		err := body.Close()
		if err != nil {
			log.Println("failed to close body", ctx, err)
		}
	}()

	contentType := ctx.Value(contextkeys.ContentType).(string)
	if contentType != MimeJSON && contentType != MimeAll {
		return ErrInvalidContentType
	}
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, request); err != nil {
		return err
	}
	return nil
}

// EncodeResponse encodes the response based on the content-type
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	contentType := ctx.Value(contextkeys.AcceptHeader).(string)
	w.Header().Set(header.ContentType, contentType)
	if response == nil {
		return nil
	}
	if contentType != MimeJSON && contentType != MimeAll {
		return ErrInvalidAcceptHeader
	}
	w.Header().Set("Content-Type", fmt.Sprintf("%s; charset=utf-8", MimeJSON))
	data, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}
