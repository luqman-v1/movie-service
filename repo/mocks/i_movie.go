// Code generated by mockery 2.9.3. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "movie-service/entity"

	mock "github.com/stretchr/testify/mock"
)

// IMovie is an autogenerated mock type for the IMovie type
type IMovie struct {
	mock.Mock
}

// Upsert provides a mock function with given fields: ctx, req
func (_m *IMovie) Upsert(ctx context.Context, req entity.Search) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Search) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
