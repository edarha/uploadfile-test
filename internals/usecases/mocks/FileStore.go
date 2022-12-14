// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// FileStore is an autogenerated mock type for the FileStore type
type FileStore struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *FileStore) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UploadFile provides a mock function with given fields: ctx, fileName, path
func (_m *FileStore) UploadFile(ctx context.Context, fileName string, path string) error {
	ret := _m.Called(ctx, fileName, path)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, fileName, path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewFileStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewFileStore creates a new instance of FileStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFileStore(t mockConstructorTestingTNewFileStore) *FileStore {
	mock := &FileStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
