package world

import (
	"context"
	"fmt"
	"github.com/Windowsfreak/go-mc/domain"
	"github.com/go-kit/kit/endpoint"
	"log"
)

func makeGetEventsAfterValidationMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			msg, ok := req.(domain.GetEventsAfterRequest)
			if !ok {
				return nil, domain.ErrInvalidMessageType
			}
			if err := validateGetEventsAfterRequest(msg); err != nil {
				return nil, fmt.Errorf("%w: %s", domain.ErrMissingArgument, err)
			}
			return next(ctx, req)
		}
	}
}

func validateGetEventsAfterRequest(msg domain.GetEventsAfterRequest) error {
	return validatePreSharedKeyRequest(msg.PreSharedKeyRequest)
}

func validatePreSharedKeyRequest(msg domain.PreSharedKeyRequest) error {
	return validatePreSharedKey(msg.PreSharedKey)
}

func validatePreSharedKey(key string) error {
	if len(domain.Config.PreSharedKey) < 1 {
		log.Fatal("pre-shared key missing in config")
	}
	if key != domain.Config.PreSharedKey {
		return domain.ErrInvalidKey
	}
	return nil
}
