// Code generated by mockery v2.12.0. DO NOT EDIT.

package provisioning

import (
	context "context"

	models "github.com/grafana/grafana/pkg/services/ngalert/models"
	mock "github.com/stretchr/testify/mock"

	store "github.com/grafana/grafana/pkg/services/ngalert/store"

	testing "testing"
)

// MockRuleStore is an autogenerated mock type for the RuleStore type
type MockRuleStore struct {
	mock.Mock
}

type MockRuleStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRuleStore) EXPECT() *MockRuleStore_Expecter {
	return &MockRuleStore_Expecter{mock: &_m.Mock}
}

// DeleteAlertRulesByUID provides a mock function with given fields: ctx, orgID, ruleUID
func (_m *MockRuleStore) DeleteAlertRulesByUID(ctx context.Context, orgID int64, ruleUID ...string) error {
	_va := make([]interface{}, len(ruleUID))
	for _i := range ruleUID {
		_va[_i] = ruleUID[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, orgID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...string) error); ok {
		r0 = rf(ctx, orgID, ruleUID...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRuleStore_DeleteAlertRulesByUID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAlertRulesByUID'
type MockRuleStore_DeleteAlertRulesByUID_Call struct {
	*mock.Call
}

// DeleteAlertRulesByUID is a helper method to define mock.On call
//  - ctx context.Context
//  - orgID int64
//  - ruleUID ...string
func (_e *MockRuleStore_Expecter) DeleteAlertRulesByUID(ctx interface{}, orgID interface{}, ruleUID ...interface{}) *MockRuleStore_DeleteAlertRulesByUID_Call {
	return &MockRuleStore_DeleteAlertRulesByUID_Call{Call: _e.mock.On("DeleteAlertRulesByUID",
		append([]interface{}{ctx, orgID}, ruleUID...)...)}
}

func (_c *MockRuleStore_DeleteAlertRulesByUID_Call) Run(run func(ctx context.Context, orgID int64, ruleUID ...string)) *MockRuleStore_DeleteAlertRulesByUID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), args[1].(int64), variadicArgs...)
	})
	return _c
}

func (_c *MockRuleStore_DeleteAlertRulesByUID_Call) Return(_a0 error) *MockRuleStore_DeleteAlertRulesByUID_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetAlertRuleByUID provides a mock function with given fields: ctx, query
func (_m *MockRuleStore) GetAlertRuleByUID(ctx context.Context, query *models.GetAlertRuleByUIDQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.GetAlertRuleByUIDQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRuleStore_GetAlertRuleByUID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAlertRuleByUID'
type MockRuleStore_GetAlertRuleByUID_Call struct {
	*mock.Call
}

// GetAlertRuleByUID is a helper method to define mock.On call
//  - ctx context.Context
//  - query *models.GetAlertRuleByUIDQuery
func (_e *MockRuleStore_Expecter) GetAlertRuleByUID(ctx interface{}, query interface{}) *MockRuleStore_GetAlertRuleByUID_Call {
	return &MockRuleStore_GetAlertRuleByUID_Call{Call: _e.mock.On("GetAlertRuleByUID", ctx, query)}
}

func (_c *MockRuleStore_GetAlertRuleByUID_Call) Run(run func(ctx context.Context, query *models.GetAlertRuleByUIDQuery)) *MockRuleStore_GetAlertRuleByUID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.GetAlertRuleByUIDQuery))
	})
	return _c
}

func (_c *MockRuleStore_GetAlertRuleByUID_Call) Return(_a0 error) *MockRuleStore_GetAlertRuleByUID_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetAlertRulesGroupByRuleUID provides a mock function with given fields: ctx, query
func (_m *MockRuleStore) GetAlertRulesGroupByRuleUID(ctx context.Context, query *models.GetAlertRulesGroupByRuleUIDQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.GetAlertRulesGroupByRuleUIDQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRuleStore_GetAlertRulesGroupByRuleUID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAlertRulesGroupByRuleUID'
type MockRuleStore_GetAlertRulesGroupByRuleUID_Call struct {
	*mock.Call
}

// GetAlertRulesGroupByRuleUID is a helper method to define mock.On call
//  - ctx context.Context
//  - query *models.GetAlertRulesGroupByRuleUIDQuery
func (_e *MockRuleStore_Expecter) GetAlertRulesGroupByRuleUID(ctx interface{}, query interface{}) *MockRuleStore_GetAlertRulesGroupByRuleUID_Call {
	return &MockRuleStore_GetAlertRulesGroupByRuleUID_Call{Call: _e.mock.On("GetAlertRulesGroupByRuleUID", ctx, query)}
}

func (_c *MockRuleStore_GetAlertRulesGroupByRuleUID_Call) Run(run func(ctx context.Context, query *models.GetAlertRulesGroupByRuleUIDQuery)) *MockRuleStore_GetAlertRulesGroupByRuleUID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.GetAlertRulesGroupByRuleUIDQuery))
	})
	return _c
}

func (_c *MockRuleStore_GetAlertRulesGroupByRuleUID_Call) Return(_a0 error) *MockRuleStore_GetAlertRulesGroupByRuleUID_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetRuleGroupInterval provides a mock function with given fields: ctx, orgID, namespaceUID, ruleGroup
func (_m *MockRuleStore) GetRuleGroupInterval(ctx context.Context, orgID int64, namespaceUID string, ruleGroup string) (int64, error) {
	ret := _m.Called(ctx, orgID, namespaceUID, ruleGroup)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string) int64); ok {
		r0 = rf(ctx, orgID, namespaceUID, ruleGroup)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, string, string) error); ok {
		r1 = rf(ctx, orgID, namespaceUID, ruleGroup)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRuleStore_GetRuleGroupInterval_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRuleGroupInterval'
type MockRuleStore_GetRuleGroupInterval_Call struct {
	*mock.Call
}

// GetRuleGroupInterval is a helper method to define mock.On call
//  - ctx context.Context
//  - orgID int64
//  - namespaceUID string
//  - ruleGroup string
func (_e *MockRuleStore_Expecter) GetRuleGroupInterval(ctx interface{}, orgID interface{}, namespaceUID interface{}, ruleGroup interface{}) *MockRuleStore_GetRuleGroupInterval_Call {
	return &MockRuleStore_GetRuleGroupInterval_Call{Call: _e.mock.On("GetRuleGroupInterval", ctx, orgID, namespaceUID, ruleGroup)}
}

func (_c *MockRuleStore_GetRuleGroupInterval_Call) Run(run func(ctx context.Context, orgID int64, namespaceUID string, ruleGroup string)) *MockRuleStore_GetRuleGroupInterval_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockRuleStore_GetRuleGroupInterval_Call) Return(_a0 int64, _a1 error) *MockRuleStore_GetRuleGroupInterval_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// InsertAlertRules provides a mock function with given fields: ctx, rule
func (_m *MockRuleStore) InsertAlertRules(ctx context.Context, rule []models.AlertRule) (map[string]int64, error) {
	ret := _m.Called(ctx, rule)

	var r0 map[string]int64
	if rf, ok := ret.Get(0).(func(context.Context, []models.AlertRule) map[string]int64); ok {
		r0 = rf(ctx, rule)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []models.AlertRule) error); ok {
		r1 = rf(ctx, rule)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRuleStore_InsertAlertRules_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertAlertRules'
type MockRuleStore_InsertAlertRules_Call struct {
	*mock.Call
}

// InsertAlertRules is a helper method to define mock.On call
//  - ctx context.Context
//  - rule []models.AlertRule
func (_e *MockRuleStore_Expecter) InsertAlertRules(ctx interface{}, rule interface{}) *MockRuleStore_InsertAlertRules_Call {
	return &MockRuleStore_InsertAlertRules_Call{Call: _e.mock.On("InsertAlertRules", ctx, rule)}
}

func (_c *MockRuleStore_InsertAlertRules_Call) Run(run func(ctx context.Context, rule []models.AlertRule)) *MockRuleStore_InsertAlertRules_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]models.AlertRule))
	})
	return _c
}

func (_c *MockRuleStore_InsertAlertRules_Call) Return(_a0 map[string]int64, _a1 error) *MockRuleStore_InsertAlertRules_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListAlertRules provides a mock function with given fields: ctx, query
func (_m *MockRuleStore) ListAlertRules(ctx context.Context, query *models.ListAlertRulesQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.ListAlertRulesQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRuleStore_ListAlertRules_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAlertRules'
type MockRuleStore_ListAlertRules_Call struct {
	*mock.Call
}

// ListAlertRules is a helper method to define mock.On call
//  - ctx context.Context
//  - query *models.ListAlertRulesQuery
func (_e *MockRuleStore_Expecter) ListAlertRules(ctx interface{}, query interface{}) *MockRuleStore_ListAlertRules_Call {
	return &MockRuleStore_ListAlertRules_Call{Call: _e.mock.On("ListAlertRules", ctx, query)}
}

func (_c *MockRuleStore_ListAlertRules_Call) Run(run func(ctx context.Context, query *models.ListAlertRulesQuery)) *MockRuleStore_ListAlertRules_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.ListAlertRulesQuery))
	})
	return _c
}

func (_c *MockRuleStore_ListAlertRules_Call) Return(_a0 error) *MockRuleStore_ListAlertRules_Call {
	_c.Call.Return(_a0)
	return _c
}

// UpdateAlertRules provides a mock function with given fields: ctx, rule
func (_m *MockRuleStore) UpdateAlertRules(ctx context.Context, rule []store.UpdateRule) error {
	ret := _m.Called(ctx, rule)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []store.UpdateRule) error); ok {
		r0 = rf(ctx, rule)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRuleStore_UpdateAlertRules_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAlertRules'
type MockRuleStore_UpdateAlertRules_Call struct {
	*mock.Call
}

// UpdateAlertRules is a helper method to define mock.On call
//  - ctx context.Context
//  - rule []store.UpdateRule
func (_e *MockRuleStore_Expecter) UpdateAlertRules(ctx interface{}, rule interface{}) *MockRuleStore_UpdateAlertRules_Call {
	return &MockRuleStore_UpdateAlertRules_Call{Call: _e.mock.On("UpdateAlertRules", ctx, rule)}
}

func (_c *MockRuleStore_UpdateAlertRules_Call) Run(run func(ctx context.Context, rule []store.UpdateRule)) *MockRuleStore_UpdateAlertRules_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]store.UpdateRule))
	})
	return _c
}

func (_c *MockRuleStore_UpdateAlertRules_Call) Return(_a0 error) *MockRuleStore_UpdateAlertRules_Call {
	_c.Call.Return(_a0)
	return _c
}

// NewMockRuleStore creates a new instance of MockRuleStore. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRuleStore(t testing.TB) *MockRuleStore {
	mock := &MockRuleStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
