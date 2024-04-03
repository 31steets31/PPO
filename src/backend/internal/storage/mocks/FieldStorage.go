// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"
	dto "course/internal/service/dto"

	mock "github.com/stretchr/testify/mock"

	model "course/internal/model"
)

// FieldStorage is an autogenerated mock type for the FieldStorage type
type FieldStorage struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, request
func (_m *FieldStorage) Create(ctx context.Context, request *dto.CreateDocumentFieldRequest) (*model.Field, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *model.Field
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.CreateDocumentFieldRequest) (*model.Field, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.CreateDocumentFieldRequest) *model.Field); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Field)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.CreateDocumentFieldRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, request
func (_m *FieldStorage) Delete(ctx context.Context, request *dto.DeleteDocumentFieldRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.DeleteDocumentFieldRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, request
func (_m *FieldStorage) Get(ctx context.Context, request *dto.GetDocumentFieldRequest) (*model.Field, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *model.Field
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetDocumentFieldRequest) (*model.Field, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetDocumentFieldRequest) *model.Field); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Field)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.GetDocumentFieldRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCardFields provides a mock function with given fields: ctx, request
func (_m *FieldStorage) ListCardFields(ctx context.Context, request *dto.ListDocumentFieldsRequest) ([]*model.Field, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListCardFields")
	}

	var r0 []*model.Field
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ListDocumentFieldsRequest) ([]*model.Field, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ListDocumentFieldsRequest) []*model.Field); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Field)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.ListDocumentFieldsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewFieldStorage creates a new instance of FieldStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFieldStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *FieldStorage {
	mock := &FieldStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
