package world

import (
	"context"
	"github.com/Windowsfreak/go-mc/domain"
	"github.com/go-kit/kit/endpoint"
	"time"
)

func makeGetEventsAfterEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.GetEventsAfterRequest)
		req.PreSharedKey = ""
		events, err := s.GetEventsAfter(req.Rowid, req.Limit)
		if err != nil {
			return nil, err
		}
		return events, err
	}
}

func makeGetEventsByUserAfterEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.GetEventsByUserAfterRequest)
		req.PreSharedKey = ""
		events, err := s.GetEventsByUserAfter(req.Rowid, req.Username, req.Limit)
		if err != nil {
			return nil, err
		}
		return events, err
	}
}

func makeGetEventsByRegionAfterEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.GetEventsByRegionAfterRequest)
		req.PreSharedKey = ""
		events, err := s.GetEventsByRegionAfter(req.Rowid, req.X1, req.Z1, req.X2, req.Z2, req.Limit)
		if err != nil {
			return nil, err
		}
		return events, err
	}
}

func makeGetLastEventRowIdBeforeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(int64)
		rowid, err := s.GetLastEventRowIdBefore(req)
		if err != nil {
			return nil, err
		}
		return rowid, err
	}
}

func makeGetLastEventRowIdByUserBeforeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.GetByTimeAndUsernameRequest)
		rowid, err := s.GetLastEventRowIdByUserBefore(req.Time, req.Username)
		if err != nil {
			return nil, err
		}
		return rowid, err
	}
}

func makeGetLastEventRowIdByRegionBeforeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.GetByTimeAndRegionRequest)
		rowid, err := s.GetLastEventRowIdByRegionBefore(req.Time, req.X1, req.Z1, req.X2, req.Z2)
		if err != nil {
			return nil, err
		}
		return rowid, err
	}
}

func makeGetChatsAfterEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.GetEventsAfterRequest)
		req.PreSharedKey = ""
		events, err := s.GetChatsAfter(req.Rowid, req.Limit)
		if err != nil {
			return nil, err
		}
		return events, err
	}
}

func makeGetLastChatRowIdBeforeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(int64)
		rowid, err := s.GetLastChatRowIdBefore(req)
		if err != nil {
			return nil, err
		}
		return rowid, err
	}
}

func makeGetUsersSeenAfterEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		now := time.Now().Unix()
		req := request.(int64)
		users, err := s.GetUsersSeenAfter(req)
		if err != nil {
			return nil, err
		}
		return domain.UserResponse{Time: now, Users: users}, err
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(string)
		user, err := s.GetUser(req)
		if err != nil {
			return nil, err
		}
		return user, err
	}
}

func makeGetSnitchesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		snitches, err := s.GetSnitches()
		if err != nil {
			return nil, err
		}
		return snitches, err
	}
}

func makeSetSnitchAlertByRegionEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.SetSnitchAlertByRegionRequest)
		err := s.SetSnitchAlert(req.X1, req.Z1, req.X2, req.Z2, req.Alert)
		if err != nil {
			return nil, err
		}
		return struct{}{}, err
	}
}
