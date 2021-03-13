package world

import (
	"context"
	"github.com/Windowsfreak/go-mc/domain"
	mhttp "github.com/Windowsfreak/go-mc/http"
	"github.com/Windowsfreak/go-mc/http/middleware"
	"github.com/Windowsfreak/go-mc/http/serveroption"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
	"strconv"
)

func MakeGetEventsAfterHandler(
	s Service,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(serveroption.ExtractContentTypeIntoContext),
		kithttp.ServerBefore(serveroption.ExtractAcceptHeaderIntoContext),
		kithttp.ServerErrorEncoder(middleware.MakeEncodeErrorFunc()),
	}

	mw := endpoint.Chain(
		middleware.MakeAcceptHeaderValidationMiddleware(),
		makeGetEventsAfterValidationMiddleware(),
	)

	endpointHandler := kithttp.NewServer(
		mw(makeGetEventsAfterEndpoint(s)),
		decodeGetEventsAfterRequest,
		mhttp.EncodeResponse,
		opts...,
	)

	return endpointHandler
}

func MakeGetEventsByUserAfterHandler(
	s Service,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(serveroption.ExtractContentTypeIntoContext),
		kithttp.ServerBefore(serveroption.ExtractAcceptHeaderIntoContext),
		kithttp.ServerErrorEncoder(middleware.MakeEncodeErrorFunc()),
	}

	mw := endpoint.Chain(
		middleware.MakeAcceptHeaderValidationMiddleware(),
	)

	endpointHandler := kithttp.NewServer(
		mw(makeGetEventsByUserAfterEndpoint(s)),
		decodeGetEventsByUserAfterRequest,
		mhttp.EncodeResponse,
		opts...,
	)

	return endpointHandler
}

func MakeGetEventsByRegionAfterHandler(
	s Service,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(serveroption.ExtractContentTypeIntoContext),
		kithttp.ServerBefore(serveroption.ExtractAcceptHeaderIntoContext),
		kithttp.ServerErrorEncoder(middleware.MakeEncodeErrorFunc()),
	}

	mw := endpoint.Chain(
		middleware.MakeAcceptHeaderValidationMiddleware(),
	)

	endpointHandler := kithttp.NewServer(
		mw(makeGetEventsByRegionAfterEndpoint(s)),
		decodeGetEventsByRegionAfterRequest,
		mhttp.EncodeResponse,
		opts...,
	)

	return endpointHandler
}

func MakeGetLastEventRowIdBeforeHandler(
	s Service,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(serveroption.ExtractContentTypeIntoContext),
		kithttp.ServerBefore(serveroption.ExtractAcceptHeaderIntoContext),
		kithttp.ServerErrorEncoder(middleware.MakeEncodeErrorFunc()),
	}

	mw := endpoint.Chain(
		middleware.MakeAcceptHeaderValidationMiddleware(),
	)

	endpointHandler := kithttp.NewServer(
		mw(makeGetLastEventRowIdBeforeEndpoint(s)),
		decodeGetLastEventRowIdBeforeRequest,
		mhttp.EncodeResponse,
		opts...,
	)

	return endpointHandler
}

func MakeGetLastEventRowIdByUserBeforeHandler(
	s Service,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(serveroption.ExtractContentTypeIntoContext),
		kithttp.ServerBefore(serveroption.ExtractAcceptHeaderIntoContext),
		kithttp.ServerErrorEncoder(middleware.MakeEncodeErrorFunc()),
	}

	mw := endpoint.Chain(
		middleware.MakeAcceptHeaderValidationMiddleware(),
	)

	endpointHandler := kithttp.NewServer(
		mw(makeGetLastEventRowIdByUserBeforeEndpoint(s)),
		decodeGetLastEventRowIdByUserBeforeRequest,
		mhttp.EncodeResponse,
		opts...,
	)

	return endpointHandler
}

func MakeGetLastEventRowIdByRegionBeforeHandler(
	s Service,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(serveroption.ExtractContentTypeIntoContext),
		kithttp.ServerBefore(serveroption.ExtractAcceptHeaderIntoContext),
		kithttp.ServerErrorEncoder(middleware.MakeEncodeErrorFunc()),
	}

	mw := endpoint.Chain(
		middleware.MakeAcceptHeaderValidationMiddleware(),
	)

	endpointHandler := kithttp.NewServer(
		mw(makeGetLastEventRowIdByRegionBeforeEndpoint(s)),
		decodeGetLastEventRowIdByRegionBeforeRequest,
		mhttp.EncodeResponse,
		opts...,
	)

	return endpointHandler
}

func MakeGetChatsAfterHandler(
	s Service,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(serveroption.ExtractContentTypeIntoContext),
		kithttp.ServerBefore(serveroption.ExtractAcceptHeaderIntoContext),
		kithttp.ServerErrorEncoder(middleware.MakeEncodeErrorFunc()),
	}

	mw := endpoint.Chain(
		middleware.MakeAcceptHeaderValidationMiddleware(),
		makeGetEventsAfterValidationMiddleware(),
	)

	endpointHandler := kithttp.NewServer(
		mw(makeGetChatsAfterEndpoint(s)),
		decodeGetEventsAfterRequest,
		mhttp.EncodeResponse,
		opts...,
	)

	return endpointHandler
}

func MakeGetLastChatRowIdBeforeHandler(
	s Service,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(serveroption.ExtractContentTypeIntoContext),
		kithttp.ServerBefore(serveroption.ExtractAcceptHeaderIntoContext),
		kithttp.ServerErrorEncoder(middleware.MakeEncodeErrorFunc()),
	}

	mw := endpoint.Chain(
		middleware.MakeAcceptHeaderValidationMiddleware(),
	)

	endpointHandler := kithttp.NewServer(
		mw(makeGetLastChatRowIdBeforeEndpoint(s)),
		decodeGetLastEventRowIdBeforeRequest,
		mhttp.EncodeResponse,
		opts...,
	)

	return endpointHandler
}

func MakeGetUsersSeenAfterHandler(
	s Service,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(serveroption.ExtractContentTypeIntoContext),
		kithttp.ServerBefore(serveroption.ExtractAcceptHeaderIntoContext),
		kithttp.ServerErrorEncoder(middleware.MakeEncodeErrorFunc()),
	}

	mw := endpoint.Chain(
		middleware.MakeAcceptHeaderValidationMiddleware(),
	)

	endpointHandler := kithttp.NewServer(
		mw(makeGetUsersSeenAfterEndpoint(s)),
		decodeGetUsersSeenAfterRequest,
		mhttp.EncodeResponse,
		opts...,
	)

	return endpointHandler
}

func MakeGetUserHandler(
	s Service,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(serveroption.ExtractContentTypeIntoContext),
		kithttp.ServerBefore(serveroption.ExtractAcceptHeaderIntoContext),
		kithttp.ServerErrorEncoder(middleware.MakeEncodeErrorFunc()),
	}

	mw := endpoint.Chain(
		middleware.MakeAcceptHeaderValidationMiddleware(),
	)

	endpointHandler := kithttp.NewServer(
		mw(makeGetUserEndpoint(s)),
		decodeGetUserRequest,
		mhttp.EncodeResponse,
		opts...,
	)

	return endpointHandler
}

func MakeGetSnitchesHandler(
	s Service,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(serveroption.ExtractContentTypeIntoContext),
		kithttp.ServerBefore(serveroption.ExtractAcceptHeaderIntoContext),
		kithttp.ServerErrorEncoder(middleware.MakeEncodeErrorFunc()),
	}

	mw := endpoint.Chain(
		middleware.MakeAcceptHeaderValidationMiddleware(),
	)

	endpointHandler := kithttp.NewServer(
		mw(makeGetSnitchesEndpoint(s)),
		DecodeGetRequest,
		mhttp.EncodeResponse,
		opts...,
	)

	return endpointHandler
}

func MakeSetSnitchAlertByRegionHandler(
	s Service,
) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(serveroption.ExtractContentTypeIntoContext),
		kithttp.ServerBefore(serveroption.ExtractAcceptHeaderIntoContext),
		kithttp.ServerErrorEncoder(middleware.MakeEncodeErrorFunc()),
	}

	mw := endpoint.Chain(
		middleware.MakeAcceptHeaderValidationMiddleware(),
	)

	endpointHandler := kithttp.NewServer(
		mw(makeSetSnitchAlertByRegionEndpoint(s)),
		decodeSetSnitchAlertByRegionRequest,
		mhttp.EncodeResponse,
		opts...,
	)

	return endpointHandler
}

func decodeGetEventsAfterRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request domain.GetEventsAfterRequest
	err := mhttp.DecodeRequest(ctx, r, &request)
	if err != nil {
		query := r.URL.Query()
		request.PreSharedKey = query.Get("token")
		request.Rowid, err = strconv.ParseInt(query.Get("rowid"), 10, 64)
		if err != nil {
			request.Rowid = 0
		}
		request.Limit, err = strconv.ParseInt(query.Get("limit"), 10, 64)
		if err != nil {
			return request, err
		}
		err = nil
	}
	return request, err
}

func decodeGetEventsByUserAfterRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request domain.GetEventsByUserAfterRequest
	err := mhttp.DecodeRequest(ctx, r, &request)
	if err != nil {
		query := r.URL.Query()
		request.PreSharedKey = query.Get("token")
		request.Rowid, err = strconv.ParseInt(query.Get("rowid"), 10, 64)
		if err != nil {
			request.Rowid = 0
		}
		request.Username = query.Get("username")
		request.Limit, err = strconv.ParseInt(query.Get("limit"), 10, 64)
		if err != nil {
			return request, err
		}
		err = nil
	}
	if err := validatePreSharedKey(request.PreSharedKey); err != nil {
		return request, err
	}
	return request, err
}

func decodeGetEventsByRegionAfterRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request domain.GetEventsByRegionAfterRequest
	err := mhttp.DecodeRequest(ctx, r, &request)
	if err != nil {
		query := r.URL.Query()
		request.PreSharedKey = query.Get("token")
		request.Rowid, err = strconv.ParseInt(query.Get("rowid"), 10, 64)
		if err != nil {
			request.Rowid = 0
		}
		request.X1, err = strconv.ParseInt(query.Get("x1"), 10, 64)
		if err != nil {
			return nil, err
		}
		request.Z1, err = strconv.ParseInt(query.Get("z1"), 10, 64)
		if err != nil {
			return nil, err
		}
		request.X2, err = strconv.ParseInt(query.Get("x2"), 10, 64)
		if err != nil {
			return nil, err
		}
		request.Z2, err = strconv.ParseInt(query.Get("z2"), 10, 64)
		if err != nil {
			return nil, err
		}
		request.Limit, err = strconv.ParseInt(query.Get("limit"), 10, 64)
		if err != nil {
			return request, err
		}
		err = nil
	}
	if err := validatePreSharedKey(request.PreSharedKey); err != nil {
		return request, err
	}
	return request, err
}

func decodeSetSnitchAlertByRegionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request domain.SetSnitchAlertByRegionRequest
	err := mhttp.DecodeRequest(ctx, r, &request)
	if err != nil {
		query := r.URL.Query()
		request.PreSharedKey = query.Get("token")
		request.Alert, err = strconv.ParseBool(query.Get("alert"))
		if err != nil {
			return nil, err
		}
		request.X1, err = strconv.ParseInt(query.Get("x1"), 10, 64)
		if err != nil {
			return nil, err
		}
		request.Z1, err = strconv.ParseInt(query.Get("z1"), 10, 64)
		if err != nil {
			return nil, err
		}
		request.X2, err = strconv.ParseInt(query.Get("x2"), 10, 64)
		if err != nil {
			return nil, err
		}
		request.Z2, err = strconv.ParseInt(query.Get("z2"), 10, 64)
		if err != nil {
			return nil, err
		}
		err = nil
	}
	if err := validatePreSharedKey(request.PreSharedKey); err != nil {
		return request, err
	}
	return request, err
}

func DecodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request domain.PreSharedKeyRequest
	err := mhttp.DecodeRequest(ctx, r, &request)
	if err != nil {
		query := r.URL.Query()
		if err = validatePreSharedKey(query.Get("token")); err != nil {
			return nil, err
		}
		err = nil
	} else {
		err = validatePreSharedKeyRequest(request)
	}
	return nil, err
}

func decodeGetLastEventRowIdBeforeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request domain.GetByTimeRequest
	err := mhttp.DecodeRequest(ctx, r, &request)
	var val int64
	if err != nil {
		query := r.URL.Query()
		if err = validatePreSharedKey(query.Get("token")); err != nil {
			return request, err
		}
		val, err = strconv.ParseInt(query.Get("time"), 10, 64)
		if err != nil {
			val = 0
		}
		err = nil
	} else {
		val = request.Time
		err = validatePreSharedKeyRequest(request.PreSharedKeyRequest)
	}
	return val, err
}

func decodeGetLastEventRowIdByUserBeforeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request domain.GetByTimeAndUsernameRequest
	err := mhttp.DecodeRequest(ctx, r, &request)
	if err != nil {
		query := r.URL.Query()
		if err = validatePreSharedKey(query.Get("token")); err != nil {
			return request, err
		}
		request.Time, err = strconv.ParseInt(query.Get("time"), 10, 64)
		if err != nil {
			request.Time = 0
		}
		request.Username = query.Get("username")
		err = nil
	} else {
		err = validatePreSharedKeyRequest(request.PreSharedKeyRequest)
	}
	return request, err
}

func decodeGetLastEventRowIdByRegionBeforeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request domain.GetByTimeAndRegionRequest
	err := mhttp.DecodeRequest(ctx, r, &request)
	if err != nil {
		query := r.URL.Query()
		if err = validatePreSharedKey(query.Get("token")); err != nil {
			return request, err
		}
		request.Time, err = strconv.ParseInt(query.Get("time"), 10, 64)
		if err != nil {
			request.Time = 0
		}
		request.X1, err = strconv.ParseInt(query.Get("x1"), 10, 64)
		if err != nil {
			return nil, err
		}
		request.Z1, err = strconv.ParseInt(query.Get("z1"), 10, 64)
		if err != nil {
			return nil, err
		}
		request.X2, err = strconv.ParseInt(query.Get("x2"), 10, 64)
		if err != nil {
			return nil, err
		}
		request.Z2, err = strconv.ParseInt(query.Get("z2"), 10, 64)
		if err != nil {
			return nil, err
		}
		err = nil
	} else {
		err = validatePreSharedKeyRequest(request.PreSharedKeyRequest)
	}
	return request, err
}

func decodeGetUsersSeenAfterRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request domain.GetByTimeRequest
	err := mhttp.DecodeRequest(ctx, r, &request)
	var val int64
	if err != nil {
		query := r.URL.Query()
		if err = validatePreSharedKey(query.Get("token")); err != nil {
			return request, err
		}
		val, err = strconv.ParseInt(query.Get("time"), 10, 64)
		if err != nil {
			val = 0
		}
		err = nil
	} else {
		val = request.Time
		err = validatePreSharedKeyRequest(request.PreSharedKeyRequest)
	}
	return val, err
}

func decodeGetUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request domain.GetByUsernameRequest
	err := mhttp.DecodeRequest(ctx, r, &request)
	var val string
	if err != nil {
		query := r.URL.Query()
		if err = validatePreSharedKey(query.Get("token")); err != nil {
			return request, err
		}
		val = query.Get("username")
		err = nil
	} else {
		val = request.Username
		err = validatePreSharedKeyRequest(request.PreSharedKeyRequest)
	}
	return val, err
}
