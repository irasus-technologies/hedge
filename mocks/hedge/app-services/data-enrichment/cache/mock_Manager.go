// Code generated by mockery v2.38.0. DO NOT EDIT.

package cache

import mock "github.com/stretchr/testify/mock"

// MockManager is an autogenerated mock type for the Manager type
type MockManager struct {
	mock.Mock
}

type MockManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockManager) EXPECT() *MockManager_Expecter {
	return &MockManager_Expecter{mock: &_m.Mock}
}

// FetchDeviceTags provides a mock function with given fields: device
func (_m *MockManager) FetchDeviceTags(device string) (interface{}, bool) {
	ret := _m.Called(device)

	if len(ret) == 0 {
		panic("no return value specified for FetchDeviceTags")
	}

	var r0 interface{}
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (interface{}, bool)); ok {
		return rf(device)
	}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(device)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(device)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// MockManager_FetchDeviceTags_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchDeviceTags'
type MockManager_FetchDeviceTags_Call struct {
	*mock.Call
}

// FetchDeviceTags is a helper method to define mock.On call
//   - device string
func (_e *MockManager_Expecter) FetchDeviceTags(device interface{}) *MockManager_FetchDeviceTags_Call {
	return &MockManager_FetchDeviceTags_Call{Call: _e.mock.On("FetchDeviceTags", device)}
}

func (_c *MockManager_FetchDeviceTags_Call) Run(run func(device string)) *MockManager_FetchDeviceTags_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockManager_FetchDeviceTags_Call) Return(_a0 interface{}, _a1 bool) *MockManager_FetchDeviceTags_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockManager_FetchDeviceTags_Call) RunAndReturn(run func(string) (interface{}, bool)) *MockManager_FetchDeviceTags_Call {
	_c.Call.Return(run)
	return _c
}

// FetchDownSamplingConfiguration provides a mock function with given fields: profile
func (_m *MockManager) FetchDownSamplingConfiguration(profile string) (interface{}, bool) {
	ret := _m.Called(profile)

	if len(ret) == 0 {
		panic("no return value specified for FetchDownSamplingConfiguration")
	}

	var r0 interface{}
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (interface{}, bool)); ok {
		return rf(profile)
	}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(profile)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(profile)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// MockManager_FetchDownSamplingConfiguration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchDownSamplingConfiguration'
type MockManager_FetchDownSamplingConfiguration_Call struct {
	*mock.Call
}

// FetchDownSamplingConfiguration is a helper method to define mock.On call
//   - profile string
func (_e *MockManager_Expecter) FetchDownSamplingConfiguration(profile interface{}) *MockManager_FetchDownSamplingConfiguration_Call {
	return &MockManager_FetchDownSamplingConfiguration_Call{Call: _e.mock.On("FetchDownSamplingConfiguration", profile)}
}

func (_c *MockManager_FetchDownSamplingConfiguration_Call) Run(run func(profile string)) *MockManager_FetchDownSamplingConfiguration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockManager_FetchDownSamplingConfiguration_Call) Return(_a0 interface{}, _a1 bool) *MockManager_FetchDownSamplingConfiguration_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockManager_FetchDownSamplingConfiguration_Call) RunAndReturn(run func(string) (interface{}, bool)) *MockManager_FetchDownSamplingConfiguration_Call {
	_c.Call.Return(run)
	return _c
}

// FetchProfileTags provides a mock function with given fields: profile
func (_m *MockManager) FetchProfileTags(profile string) (interface{}, bool) {
	ret := _m.Called(profile)

	if len(ret) == 0 {
		panic("no return value specified for FetchProfileTags")
	}

	var r0 interface{}
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (interface{}, bool)); ok {
		return rf(profile)
	}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(profile)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(profile)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// MockManager_FetchProfileTags_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchProfileTags'
type MockManager_FetchProfileTags_Call struct {
	*mock.Call
}

// FetchProfileTags is a helper method to define mock.On call
//   - profile string
func (_e *MockManager_Expecter) FetchProfileTags(profile interface{}) *MockManager_FetchProfileTags_Call {
	return &MockManager_FetchProfileTags_Call{Call: _e.mock.On("FetchProfileTags", profile)}
}

func (_c *MockManager_FetchProfileTags_Call) Run(run func(profile string)) *MockManager_FetchProfileTags_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockManager_FetchProfileTags_Call) Return(_a0 interface{}, _a1 bool) *MockManager_FetchProfileTags_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockManager_FetchProfileTags_Call) RunAndReturn(run func(string) (interface{}, bool)) *MockManager_FetchProfileTags_Call {
	_c.Call.Return(run)
	return _c
}

// FetchRawDataConfiguration provides a mock function with given fields:
func (_m *MockManager) FetchRawDataConfiguration() (interface{}, bool) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FetchRawDataConfiguration")
	}

	var r0 interface{}
	var r1 bool
	if rf, ok := ret.Get(0).(func() (interface{}, bool)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// MockManager_FetchRawDataConfiguration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchRawDataConfiguration'
type MockManager_FetchRawDataConfiguration_Call struct {
	*mock.Call
}

// FetchRawDataConfiguration is a helper method to define mock.On call
func (_e *MockManager_Expecter) FetchRawDataConfiguration() *MockManager_FetchRawDataConfiguration_Call {
	return &MockManager_FetchRawDataConfiguration_Call{Call: _e.mock.On("FetchRawDataConfiguration")}
}

func (_c *MockManager_FetchRawDataConfiguration_Call) Run(run func()) *MockManager_FetchRawDataConfiguration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockManager_FetchRawDataConfiguration_Call) Return(_a0 interface{}, _a1 bool) *MockManager_FetchRawDataConfiguration_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockManager_FetchRawDataConfiguration_Call) RunAndReturn(run func() (interface{}, bool)) *MockManager_FetchRawDataConfiguration_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateDeviceTags provides a mock function with given fields: deviceName, forceFetch
func (_m *MockManager) UpdateDeviceTags(deviceName string, forceFetch bool) {
	_m.Called(deviceName, forceFetch)
}

// MockManager_UpdateDeviceTags_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDeviceTags'
type MockManager_UpdateDeviceTags_Call struct {
	*mock.Call
}

// UpdateDeviceTags is a helper method to define mock.On call
//   - deviceName string
//   - forceFetch bool
func (_e *MockManager_Expecter) UpdateDeviceTags(deviceName interface{}, forceFetch interface{}) *MockManager_UpdateDeviceTags_Call {
	return &MockManager_UpdateDeviceTags_Call{Call: _e.mock.On("UpdateDeviceTags", deviceName, forceFetch)}
}

func (_c *MockManager_UpdateDeviceTags_Call) Run(run func(deviceName string, forceFetch bool)) *MockManager_UpdateDeviceTags_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(bool))
	})
	return _c
}

func (_c *MockManager_UpdateDeviceTags_Call) Return() *MockManager_UpdateDeviceTags_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockManager_UpdateDeviceTags_Call) RunAndReturn(run func(string, bool)) *MockManager_UpdateDeviceTags_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateNodeRawData provides a mock function with given fields: forceFetch
func (_m *MockManager) UpdateNodeRawData(forceFetch bool) {
	_m.Called(forceFetch)
}

// MockManager_UpdateNodeRawData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateNodeRawData'
type MockManager_UpdateNodeRawData_Call struct {
	*mock.Call
}

// UpdateNodeRawData is a helper method to define mock.On call
//   - forceFetch bool
func (_e *MockManager_Expecter) UpdateNodeRawData(forceFetch interface{}) *MockManager_UpdateNodeRawData_Call {
	return &MockManager_UpdateNodeRawData_Call{Call: _e.mock.On("UpdateNodeRawData", forceFetch)}
}

func (_c *MockManager_UpdateNodeRawData_Call) Run(run func(forceFetch bool)) *MockManager_UpdateNodeRawData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *MockManager_UpdateNodeRawData_Call) Return() *MockManager_UpdateNodeRawData_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockManager_UpdateNodeRawData_Call) RunAndReturn(run func(bool)) *MockManager_UpdateNodeRawData_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateProfileDownSampling provides a mock function with given fields: profileName, forceFetch
func (_m *MockManager) UpdateProfileDownSampling(profileName string, forceFetch bool) {
	_m.Called(profileName, forceFetch)
}

// MockManager_UpdateProfileDownSampling_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateProfileDownSampling'
type MockManager_UpdateProfileDownSampling_Call struct {
	*mock.Call
}

// UpdateProfileDownSampling is a helper method to define mock.On call
//   - profileName string
//   - forceFetch bool
func (_e *MockManager_Expecter) UpdateProfileDownSampling(profileName interface{}, forceFetch interface{}) *MockManager_UpdateProfileDownSampling_Call {
	return &MockManager_UpdateProfileDownSampling_Call{Call: _e.mock.On("UpdateProfileDownSampling", profileName, forceFetch)}
}

func (_c *MockManager_UpdateProfileDownSampling_Call) Run(run func(profileName string, forceFetch bool)) *MockManager_UpdateProfileDownSampling_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(bool))
	})
	return _c
}

func (_c *MockManager_UpdateProfileDownSampling_Call) Return() *MockManager_UpdateProfileDownSampling_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockManager_UpdateProfileDownSampling_Call) RunAndReturn(run func(string, bool)) *MockManager_UpdateProfileDownSampling_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateProfileTags provides a mock function with given fields: profileName, forceFetch
func (_m *MockManager) UpdateProfileTags(profileName string, forceFetch bool) {
	_m.Called(profileName, forceFetch)
}

// MockManager_UpdateProfileTags_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateProfileTags'
type MockManager_UpdateProfileTags_Call struct {
	*mock.Call
}

// UpdateProfileTags is a helper method to define mock.On call
//   - profileName string
//   - forceFetch bool
func (_e *MockManager_Expecter) UpdateProfileTags(profileName interface{}, forceFetch interface{}) *MockManager_UpdateProfileTags_Call {
	return &MockManager_UpdateProfileTags_Call{Call: _e.mock.On("UpdateProfileTags", profileName, forceFetch)}
}

func (_c *MockManager_UpdateProfileTags_Call) Run(run func(profileName string, forceFetch bool)) *MockManager_UpdateProfileTags_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(bool))
	})
	return _c
}

func (_c *MockManager_UpdateProfileTags_Call) Return() *MockManager_UpdateProfileTags_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockManager_UpdateProfileTags_Call) RunAndReturn(run func(string, bool)) *MockManager_UpdateProfileTags_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockManager creates a new instance of MockManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockManager {
	mock := &MockManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
