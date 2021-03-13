package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/Windowsfreak/go-mc/domain"
	mhttp "github.com/Windowsfreak/go-mc/http"
	"github.com/Windowsfreak/go-mc/http/contextkeys"
)

// encode errors from business-logic
func MakeEncodeErrorFunc() kithttp.ErrorEncoder {
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		contentType := ctx.Value(contextkeys.AcceptHeader).(string)
		if strings.Contains(contentType, "text/html") {
			err = sendBrowserDoc(w)
			if err == nil {
				return
			}
		}

		w.Header().Set("Content-Type", contentType)

		message := &domain.Error{}
		message.ErrorMessage = toErrorMessage(err)
		w.WriteHeader(toHTTPStatusCode(err))

		if contentType != mhttp.MimeJSON && contentType != mhttp.MimeAll {
			log.Println("contentType unknown. Please use JSON.", contentType, ctx)
			return
		}
		data, err := json.MarshalIndent(message, "", "  ")
		if err != nil {
			log.Println("failed marshalling error message to JSON.", err, ctx)
		}
		_, err = w.Write(data)
		if err != nil {
			log.Println("failed writing error message.", err, ctx)
		}
	}
}

func toErrorMessage(err error) string {
	return err.Error()
}

func toHTTPStatusCode(err error) int {
	switch {
	case errors.Is(err, mhttp.ErrInvalidContentType):
		return http.StatusUnsupportedMediaType
	case errors.Is(err, mhttp.ErrInvalidAcceptHeader):
		return http.StatusNotAcceptable
	case errors.Is(err, domain.ErrMissingArgument):
		return http.StatusBadRequest
	case errors.Is(err, domain.ErrInvalidMessageType):
		return http.StatusBadRequest
	case errors.Is(err, domain.ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, domain.ErrInvalidKey):
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

func sendBrowserDoc(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusUnsupportedMediaType)
	b, err := ioutil.ReadFile("static/browser.htm")
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}
