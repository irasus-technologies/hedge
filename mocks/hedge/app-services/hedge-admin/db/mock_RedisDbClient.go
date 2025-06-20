// Code generated by mockery v2.38.0. DO NOT EDIT.

package db

import (
	"hedge/common/dto"
	errors "hedge/common/errors"

	mock "github.com/stretchr/testify/mock"

	models "hedge/app-services/hedge-admin/models"
)

// MockRedisDbClient is an autogenerated mock type for the RedisDbClient type
type MockRedisDbClient struct {
	mock.Mock
}

type MockRedisDbClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRedisDbClient) EXPECT() *MockRedisDbClient_Expecter {
	return &MockRedisDbClient_Expecter{mock: &_m.Mock}
}

// DeleteNode provides a mock function with given fields: nodeId, keyFieldTuples
func (_m *MockRedisDbClient) DeleteNode(nodeId string, keyFieldTuples []models.KeyFieldTuple) errors.HedgeError {
	ret := _m.Called(nodeId, keyFieldTuples)

	if len(ret) == 0 {
		panic("no return value specified for DeleteNode")
	}

	var r0 errors.HedgeError
	if rf, ok := ret.Get(0).(func(string, []models.KeyFieldTuple) errors.HedgeError); ok {
		r0 = rf(nodeId, keyFieldTuples)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.HedgeError)
		}
	}

	return r0
}

// MockRedisDbClient_DeleteNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteNode'
type MockRedisDbClient_DeleteNode_Call struct {
	*mock.Call
}

// DeleteNode is a helper method to define mock.On call
//   - nodeId string
//   - keyFieldTuples []models.KeyFieldTuple
func (_e *MockRedisDbClient_Expecter) DeleteNode(nodeId interface{}, keyFieldTuples interface{}) *MockRedisDbClient_DeleteNode_Call {
	return &MockRedisDbClient_DeleteNode_Call{Call: _e.mock.On("DeleteNode", nodeId, keyFieldTuples)}
}

func (_c *MockRedisDbClient_DeleteNode_Call) Run(run func(nodeId string, keyFieldTuples []models.KeyFieldTuple)) *MockRedisDbClient_DeleteNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]models.KeyFieldTuple))
	})
	return _c
}

func (_c *MockRedisDbClient_DeleteNode_Call) Return(_a0 errors.HedgeError) *MockRedisDbClient_DeleteNode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRedisDbClient_DeleteNode_Call) RunAndReturn(run func(string, []models.KeyFieldTuple) errors.HedgeError) *MockRedisDbClient_DeleteNode_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteNodeGroup provides a mock function with given fields: parentNodeGroupName, field
func (_m *MockRedisDbClient) DeleteNodeGroup(parentNodeGroupName string, field string) errors.HedgeError {
	ret := _m.Called(parentNodeGroupName, field)

	if len(ret) == 0 {
		panic("no return value specified for DeleteNodeGroup")
	}

	var r0 errors.HedgeError
	if rf, ok := ret.Get(0).(func(string, string) errors.HedgeError); ok {
		r0 = rf(parentNodeGroupName, field)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.HedgeError)
		}
	}

	return r0
}

// MockRedisDbClient_DeleteNodeGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteNodeGroup'
type MockRedisDbClient_DeleteNodeGroup_Call struct {
	*mock.Call
}

// DeleteNodeGroup is a helper method to define mock.On call
//   - parentNodeGroupName string
//   - field string
func (_e *MockRedisDbClient_Expecter) DeleteNodeGroup(parentNodeGroupName interface{}, field interface{}) *MockRedisDbClient_DeleteNodeGroup_Call {
	return &MockRedisDbClient_DeleteNodeGroup_Call{Call: _e.mock.On("DeleteNodeGroup", parentNodeGroupName, field)}
}

func (_c *MockRedisDbClient_DeleteNodeGroup_Call) Run(run func(parentNodeGroupName string, field string)) *MockRedisDbClient_DeleteNodeGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockRedisDbClient_DeleteNodeGroup_Call) Return(_a0 errors.HedgeError) *MockRedisDbClient_DeleteNodeGroup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRedisDbClient_DeleteNodeGroup_Call) RunAndReturn(run func(string, string) errors.HedgeError) *MockRedisDbClient_DeleteNodeGroup_Call {
	_c.Call.Return(run)
	return _c
}

// FindNodeKey provides a mock function with given fields: parentGroup
func (_m *MockRedisDbClient) FindNodeKey(parentGroup string) (string, errors.HedgeError) {
	ret := _m.Called(parentGroup)

	if len(ret) == 0 {
		panic("no return value specified for FindNodeKey")
	}

	var r0 string
	var r1 errors.HedgeError
	if rf, ok := ret.Get(0).(func(string) (string, errors.HedgeError)); ok {
		return rf(parentGroup)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(parentGroup)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) errors.HedgeError); ok {
		r1 = rf(parentGroup)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.HedgeError)
		}
	}

	return r0, r1
}

// MockRedisDbClient_FindNodeKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindNodeKey'
type MockRedisDbClient_FindNodeKey_Call struct {
	*mock.Call
}

// FindNodeKey is a helper method to define mock.On call
//   - parentGroup string
func (_e *MockRedisDbClient_Expecter) FindNodeKey(parentGroup interface{}) *MockRedisDbClient_FindNodeKey_Call {
	return &MockRedisDbClient_FindNodeKey_Call{Call: _e.mock.On("FindNodeKey", parentGroup)}
}

func (_c *MockRedisDbClient_FindNodeKey_Call) Run(run func(parentGroup string)) *MockRedisDbClient_FindNodeKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockRedisDbClient_FindNodeKey_Call) Return(_a0 string, _a1 errors.HedgeError) *MockRedisDbClient_FindNodeKey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRedisDbClient_FindNodeKey_Call) RunAndReturn(run func(string) (string, errors.HedgeError)) *MockRedisDbClient_FindNodeKey_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllNodes provides a mock function with given fields:
func (_m *MockRedisDbClient) GetAllNodes() ([]dto.Node, errors.HedgeError) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllNodes")
	}

	var r0 []dto.Node
	var r1 errors.HedgeError
	if rf, ok := ret.Get(0).(func() ([]dto.Node, errors.HedgeError)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []dto.Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.Node)
		}
	}

	if rf, ok := ret.Get(1).(func() errors.HedgeError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.HedgeError)
		}
	}

	return r0, r1
}

// MockRedisDbClient_GetAllNodes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllNodes'
type MockRedisDbClient_GetAllNodes_Call struct {
	*mock.Call
}

// GetAllNodes is a helper method to define mock.On call
func (_e *MockRedisDbClient_Expecter) GetAllNodes() *MockRedisDbClient_GetAllNodes_Call {
	return &MockRedisDbClient_GetAllNodes_Call{Call: _e.mock.On("GetAllNodes")}
}

func (_c *MockRedisDbClient_GetAllNodes_Call) Run(run func()) *MockRedisDbClient_GetAllNodes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRedisDbClient_GetAllNodes_Call) Return(_a0 []dto.Node, _a1 errors.HedgeError) *MockRedisDbClient_GetAllNodes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRedisDbClient_GetAllNodes_Call) RunAndReturn(run func() ([]dto.Node, errors.HedgeError)) *MockRedisDbClient_GetAllNodes_Call {
	_c.Call.Return(run)
	return _c
}

// GetDBNodeGroupMembers provides a mock function with given fields: nodeHashKey
func (_m *MockRedisDbClient) GetDBNodeGroupMembers(nodeHashKey string) ([]dto.DBNodeGroup, errors.HedgeError) {
	ret := _m.Called(nodeHashKey)

	if len(ret) == 0 {
		panic("no return value specified for GetDBNodeGroupMembers")
	}

	var r0 []dto.DBNodeGroup
	var r1 errors.HedgeError
	if rf, ok := ret.Get(0).(func(string) ([]dto.DBNodeGroup, errors.HedgeError)); ok {
		return rf(nodeHashKey)
	}
	if rf, ok := ret.Get(0).(func(string) []dto.DBNodeGroup); ok {
		r0 = rf(nodeHashKey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.DBNodeGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(string) errors.HedgeError); ok {
		r1 = rf(nodeHashKey)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.HedgeError)
		}
	}

	return r0, r1
}

// MockRedisDbClient_GetDBNodeGroupMembers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDBNodeGroupMembers'
type MockRedisDbClient_GetDBNodeGroupMembers_Call struct {
	*mock.Call
}

// GetDBNodeGroupMembers is a helper method to define mock.On call
//   - nodeHashKey string
func (_e *MockRedisDbClient_Expecter) GetDBNodeGroupMembers(nodeHashKey interface{}) *MockRedisDbClient_GetDBNodeGroupMembers_Call {
	return &MockRedisDbClient_GetDBNodeGroupMembers_Call{Call: _e.mock.On("GetDBNodeGroupMembers", nodeHashKey)}
}

func (_c *MockRedisDbClient_GetDBNodeGroupMembers_Call) Run(run func(nodeHashKey string)) *MockRedisDbClient_GetDBNodeGroupMembers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockRedisDbClient_GetDBNodeGroupMembers_Call) Return(_a0 []dto.DBNodeGroup, _a1 errors.HedgeError) *MockRedisDbClient_GetDBNodeGroupMembers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRedisDbClient_GetDBNodeGroupMembers_Call) RunAndReturn(run func(string) ([]dto.DBNodeGroup, errors.HedgeError)) *MockRedisDbClient_GetDBNodeGroupMembers_Call {
	_c.Call.Return(run)
	return _c
}

// GetNode provides a mock function with given fields: nodeName
func (_m *MockRedisDbClient) GetNode(nodeName string) (*dto.Node, errors.HedgeError) {
	ret := _m.Called(nodeName)

	if len(ret) == 0 {
		panic("no return value specified for GetNode")
	}

	var r0 *dto.Node
	var r1 errors.HedgeError
	if rf, ok := ret.Get(0).(func(string) (*dto.Node, errors.HedgeError)); ok {
		return rf(nodeName)
	}
	if rf, ok := ret.Get(0).(func(string) *dto.Node); ok {
		r0 = rf(nodeName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(string) errors.HedgeError); ok {
		r1 = rf(nodeName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.HedgeError)
		}
	}

	return r0, r1
}

// MockRedisDbClient_GetNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNode'
type MockRedisDbClient_GetNode_Call struct {
	*mock.Call
}

// GetNode is a helper method to define mock.On call
//   - nodeName string
func (_e *MockRedisDbClient_Expecter) GetNode(nodeName interface{}) *MockRedisDbClient_GetNode_Call {
	return &MockRedisDbClient_GetNode_Call{Call: _e.mock.On("GetNode", nodeName)}
}

func (_c *MockRedisDbClient_GetNode_Call) Run(run func(nodeName string)) *MockRedisDbClient_GetNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockRedisDbClient_GetNode_Call) Return(_a0 *dto.Node, _a1 errors.HedgeError) *MockRedisDbClient_GetNode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRedisDbClient_GetNode_Call) RunAndReturn(run func(string) (*dto.Node, errors.HedgeError)) *MockRedisDbClient_GetNode_Call {
	_c.Call.Return(run)
	return _c
}

// GetNodeGroup provides a mock function with given fields: nodeGroupName
func (_m *MockRedisDbClient) GetNodeGroup(nodeGroupName string) (*dto.DBNodeGroup, errors.HedgeError) {
	ret := _m.Called(nodeGroupName)

	if len(ret) == 0 {
		panic("no return value specified for GetNodeGroup")
	}

	var r0 *dto.DBNodeGroup
	var r1 errors.HedgeError
	if rf, ok := ret.Get(0).(func(string) (*dto.DBNodeGroup, errors.HedgeError)); ok {
		return rf(nodeGroupName)
	}
	if rf, ok := ret.Get(0).(func(string) *dto.DBNodeGroup); ok {
		r0 = rf(nodeGroupName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.DBNodeGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(string) errors.HedgeError); ok {
		r1 = rf(nodeGroupName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.HedgeError)
		}
	}

	return r0, r1
}

// MockRedisDbClient_GetNodeGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNodeGroup'
type MockRedisDbClient_GetNodeGroup_Call struct {
	*mock.Call
}

// GetNodeGroup is a helper method to define mock.On call
//   - nodeGroupName string
func (_e *MockRedisDbClient_Expecter) GetNodeGroup(nodeGroupName interface{}) *MockRedisDbClient_GetNodeGroup_Call {
	return &MockRedisDbClient_GetNodeGroup_Call{Call: _e.mock.On("GetNodeGroup", nodeGroupName)}
}

func (_c *MockRedisDbClient_GetNodeGroup_Call) Run(run func(nodeGroupName string)) *MockRedisDbClient_GetNodeGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockRedisDbClient_GetNodeGroup_Call) Return(_a0 *dto.DBNodeGroup, _a1 errors.HedgeError) *MockRedisDbClient_GetNodeGroup_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRedisDbClient_GetNodeGroup_Call) RunAndReturn(run func(string) (*dto.DBNodeGroup, errors.HedgeError)) *MockRedisDbClient_GetNodeGroup_Call {
	_c.Call.Return(run)
	return _c
}

// SaveNode provides a mock function with given fields: node
func (_m *MockRedisDbClient) SaveNode(node *dto.Node) (string, errors.HedgeError) {
	ret := _m.Called(node)

	if len(ret) == 0 {
		panic("no return value specified for SaveNode")
	}

	var r0 string
	var r1 errors.HedgeError
	if rf, ok := ret.Get(0).(func(*dto.Node) (string, errors.HedgeError)); ok {
		return rf(node)
	}
	if rf, ok := ret.Get(0).(func(*dto.Node) string); ok {
		r0 = rf(node)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*dto.Node) errors.HedgeError); ok {
		r1 = rf(node)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.HedgeError)
		}
	}

	return r0, r1
}

// MockRedisDbClient_SaveNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveNode'
type MockRedisDbClient_SaveNode_Call struct {
	*mock.Call
}

// SaveNode is a helper method to define mock.On call
//   - node *commonmodels.Node
func (_e *MockRedisDbClient_Expecter) SaveNode(node interface{}) *MockRedisDbClient_SaveNode_Call {
	return &MockRedisDbClient_SaveNode_Call{Call: _e.mock.On("SaveNode", node)}
}

func (_c *MockRedisDbClient_SaveNode_Call) Run(run func(node *dto.Node)) *MockRedisDbClient_SaveNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*dto.Node))
	})
	return _c
}

func (_c *MockRedisDbClient_SaveNode_Call) Return(_a0 string, _a1 errors.HedgeError) *MockRedisDbClient_SaveNode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRedisDbClient_SaveNode_Call) RunAndReturn(run func(*dto.Node) (string, errors.HedgeError)) *MockRedisDbClient_SaveNode_Call {
	_c.Call.Return(run)
	return _c
}

// SaveNodeGroup provides a mock function with given fields: parentGroup, dbNodeGroup
func (_m *MockRedisDbClient) SaveNodeGroup(parentGroup string, dbNodeGroup *dto.DBNodeGroup) (string, errors.HedgeError) {
	ret := _m.Called(parentGroup, dbNodeGroup)

	if len(ret) == 0 {
		panic("no return value specified for SaveNodeGroup")
	}

	var r0 string
	var r1 errors.HedgeError
	if rf, ok := ret.Get(0).(func(string, *dto.DBNodeGroup) (string, errors.HedgeError)); ok {
		return rf(parentGroup, dbNodeGroup)
	}
	if rf, ok := ret.Get(0).(func(string, *dto.DBNodeGroup) string); ok {
		r0 = rf(parentGroup, dbNodeGroup)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, *dto.DBNodeGroup) errors.HedgeError); ok {
		r1 = rf(parentGroup, dbNodeGroup)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.HedgeError)
		}
	}

	return r0, r1
}

// MockRedisDbClient_SaveNodeGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveNodeGroup'
type MockRedisDbClient_SaveNodeGroup_Call struct {
	*mock.Call
}

// SaveNodeGroup is a helper method to define mock.On call
//   - parentGroup string
//   - dbNodeGroup *commonmodels.DBNodeGroup
func (_e *MockRedisDbClient_Expecter) SaveNodeGroup(parentGroup interface{}, dbNodeGroup interface{}) *MockRedisDbClient_SaveNodeGroup_Call {
	return &MockRedisDbClient_SaveNodeGroup_Call{Call: _e.mock.On("SaveNodeGroup", parentGroup, dbNodeGroup)}
}

func (_c *MockRedisDbClient_SaveNodeGroup_Call) Run(run func(parentGroup string, dbNodeGroup *dto.DBNodeGroup)) *MockRedisDbClient_SaveNodeGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*dto.DBNodeGroup))
	})
	return _c
}

func (_c *MockRedisDbClient_SaveNodeGroup_Call) Return(_a0 string, _a1 errors.HedgeError) *MockRedisDbClient_SaveNodeGroup_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRedisDbClient_SaveNodeGroup_Call) RunAndReturn(run func(string, *dto.DBNodeGroup) (string, errors.HedgeError)) *MockRedisDbClient_SaveNodeGroup_Call {
	_c.Call.Return(run)
	return _c
}

// SaveProtocols provides a mock function with given fields: dsps
func (_m *MockRedisDbClient) SaveProtocols(dsps dto.DeviceServiceProtocols) ([]string, errors.HedgeError) {
	ret := _m.Called(dsps)

	if len(ret) == 0 {
		panic("no return value specified for SaveProtocols")
	}

	var r0 []string
	var r1 errors.HedgeError
	if rf, ok := ret.Get(0).(func(dto.DeviceServiceProtocols) ([]string, errors.HedgeError)); ok {
		return rf(dsps)
	}
	if rf, ok := ret.Get(0).(func(dto.DeviceServiceProtocols) []string); ok {
		r0 = rf(dsps)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(dto.DeviceServiceProtocols) errors.HedgeError); ok {
		r1 = rf(dsps)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.HedgeError)
		}
	}

	return r0, r1
}

// MockRedisDbClient_SaveProtocols_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveProtocols'
type MockRedisDbClient_SaveProtocols_Call struct {
	*mock.Call
}

// SaveProtocols is a helper method to define mock.On call
//   - dsps commonmodels.DeviceServiceProtocols
func (_e *MockRedisDbClient_Expecter) SaveProtocols(dsps interface{}) *MockRedisDbClient_SaveProtocols_Call {
	return &MockRedisDbClient_SaveProtocols_Call{Call: _e.mock.On("SaveProtocols", dsps)}
}

func (_c *MockRedisDbClient_SaveProtocols_Call) Run(run func(dsps dto.DeviceServiceProtocols)) *MockRedisDbClient_SaveProtocols_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(dto.DeviceServiceProtocols))
	})
	return _c
}

func (_c *MockRedisDbClient_SaveProtocols_Call) Return(_a0 []string, _a1 errors.HedgeError) *MockRedisDbClient_SaveProtocols_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRedisDbClient_SaveProtocols_Call) RunAndReturn(run func(dto.DeviceServiceProtocols) ([]string, errors.HedgeError)) *MockRedisDbClient_SaveProtocols_Call {
	_c.Call.Return(run)
	return _c
}

// UpsertChildNodeGroups provides a mock function with given fields: parentNodeName, childNodes
func (_m *MockRedisDbClient) UpsertChildNodeGroups(parentNodeName string, childNodes []string) (string, errors.HedgeError) {
	ret := _m.Called(parentNodeName, childNodes)

	if len(ret) == 0 {
		panic("no return value specified for UpsertChildNodeGroups")
	}

	var r0 string
	var r1 errors.HedgeError
	if rf, ok := ret.Get(0).(func(string, []string) (string, errors.HedgeError)); ok {
		return rf(parentNodeName, childNodes)
	}
	if rf, ok := ret.Get(0).(func(string, []string) string); ok {
		r0 = rf(parentNodeName, childNodes)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, []string) errors.HedgeError); ok {
		r1 = rf(parentNodeName, childNodes)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.HedgeError)
		}
	}

	return r0, r1
}

// MockRedisDbClient_UpsertChildNodeGroups_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpsertChildNodeGroups'
type MockRedisDbClient_UpsertChildNodeGroups_Call struct {
	*mock.Call
}

// UpsertChildNodeGroups is a helper method to define mock.On call
//   - parentNodeName string
//   - childNodes []string
func (_e *MockRedisDbClient_Expecter) UpsertChildNodeGroups(parentNodeName interface{}, childNodes interface{}) *MockRedisDbClient_UpsertChildNodeGroups_Call {
	return &MockRedisDbClient_UpsertChildNodeGroups_Call{Call: _e.mock.On("UpsertChildNodeGroups", parentNodeName, childNodes)}
}

func (_c *MockRedisDbClient_UpsertChildNodeGroups_Call) Run(run func(parentNodeName string, childNodes []string)) *MockRedisDbClient_UpsertChildNodeGroups_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]string))
	})
	return _c
}

func (_c *MockRedisDbClient_UpsertChildNodeGroups_Call) Return(_a0 string, _a1 errors.HedgeError) *MockRedisDbClient_UpsertChildNodeGroups_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRedisDbClient_UpsertChildNodeGroups_Call) RunAndReturn(run func(string, []string) (string, errors.HedgeError)) *MockRedisDbClient_UpsertChildNodeGroups_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRedisDbClient creates a new instance of MockRedisDbClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRedisDbClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRedisDbClient {
	mock := &MockRedisDbClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
