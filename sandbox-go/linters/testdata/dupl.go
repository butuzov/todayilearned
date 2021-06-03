package testdata

// PROJECT URL: https://github.com/mibk/dupl

import (
	"context"
)

type (
	permissionUser struct{}
	filterResult   struct{}
	modelTours     []modelTour
	modelTour      struct{}
	iamOverview    struct{}
)

type Tour interface {
	Find(ctx context.Context, user permissionUser, query filterResult) (modelTours, error)
	Get(ctx context.Context, user permissionUser, id int) (modelTour, error)
	Create(ctx context.Context, user permissionUser, t modelTour) (int, error)
	Update(ctx context.Context, user permissionUser, id int, t modelTour) (modelTour, error)
	InUse(ctx context.Context, vehicleID int) (bool, error)
	TourStatuses() []string
	JoinChat(ctx context.Context, token string, tourID int, user iamOverview) (int, error)
}

type TourServiceMock struct {
	MockFind         func(ctx context.Context, user permissionUser, query filterResult) (modelTours, error)
	MockGet          func(ctx context.Context, user permissionUser, id int) (modelTour, error)
	MockCreate       func(ctx context.Context, user permissionUser, t modelTour) (int, error)
	MockUpdate       func(ctx context.Context, user permissionUser, id int, t modelTour) (modelTour, error)
	MockInUse        func(ctx context.Context, vehicleID int) (bool, error)
	MockTourStatuses func() []string
	MockJoinChat     func(ctx context.Context, token string, tourID int, user iamOverview) (int, error)
}

var globalvar = "no ok"
