// Code generated by mockery 2.9.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ICurl is an autogenerated mock type for the ICurl type
type ICurl struct {
	mock.Mock
}

// Get provides a mock function with given fields: endpoint, qParam
func (_m *ICurl) Get(endpoint string, qParam map[string]string) ([]byte, error) {
	ret := _m.Called(endpoint, qParam)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, map[string]string) []byte); ok {
		r0 = rf(endpoint, qParam)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, map[string]string) error); ok {
		r1 = rf(endpoint, qParam)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}