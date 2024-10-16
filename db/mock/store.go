// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/symyzi/financial-helper/db/gen (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/symyzi/financial-helper/db/gen"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateBudget mocks base method.
func (m *MockStore) CreateBudget(arg0 context.Context, arg1 db.CreateBudgetParams) (db.Budget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBudget", arg0, arg1)
	ret0, _ := ret[0].(db.Budget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBudget indicates an expected call of CreateBudget.
func (mr *MockStoreMockRecorder) CreateBudget(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBudget", reflect.TypeOf((*MockStore)(nil).CreateBudget), arg0, arg1)
}

// CreateCategory mocks base method.
func (m *MockStore) CreateCategory(arg0 context.Context, arg1 string) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockStoreMockRecorder) CreateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockStore)(nil).CreateCategory), arg0, arg1)
}

// CreateExpense mocks base method.
func (m *MockStore) CreateExpense(arg0 context.Context, arg1 db.CreateExpenseParams) (db.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExpense", arg0, arg1)
	ret0, _ := ret[0].(db.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateExpense indicates an expected call of CreateExpense.
func (mr *MockStoreMockRecorder) CreateExpense(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExpense", reflect.TypeOf((*MockStore)(nil).CreateExpense), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// CreateWallet mocks base method.
func (m *MockStore) CreateWallet(arg0 context.Context, arg1 db.CreateWalletParams) (db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWallet", arg0, arg1)
	ret0, _ := ret[0].(db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWallet indicates an expected call of CreateWallet.
func (mr *MockStoreMockRecorder) CreateWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockStore)(nil).CreateWallet), arg0, arg1)
}

// DeleteBudget mocks base method.
func (m *MockStore) DeleteBudget(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBudget", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBudget indicates an expected call of DeleteBudget.
func (mr *MockStoreMockRecorder) DeleteBudget(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBudget", reflect.TypeOf((*MockStore)(nil).DeleteBudget), arg0, arg1)
}

// DeleteCategory mocks base method.
func (m *MockStore) DeleteCategory(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategory", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCategory indicates an expected call of DeleteCategory.
func (mr *MockStoreMockRecorder) DeleteCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockStore)(nil).DeleteCategory), arg0, arg1)
}

// DeleteExpense mocks base method.
func (m *MockStore) DeleteExpense(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExpense", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExpense indicates an expected call of DeleteExpense.
func (mr *MockStoreMockRecorder) DeleteExpense(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpense", reflect.TypeOf((*MockStore)(nil).DeleteExpense), arg0, arg1)
}

// DeleteWallet mocks base method.
func (m *MockStore) DeleteWallet(arg0 context.Context, arg1 db.DeleteWalletParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWallet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWallet indicates an expected call of DeleteWallet.
func (mr *MockStoreMockRecorder) DeleteWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWallet", reflect.TypeOf((*MockStore)(nil).DeleteWallet), arg0, arg1)
}

// GetAllCategories mocks base method.
func (m *MockStore) GetAllCategories(arg0 context.Context) ([]db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCategories", arg0)
	ret0, _ := ret[0].([]db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCategories indicates an expected call of GetAllCategories.
func (mr *MockStoreMockRecorder) GetAllCategories(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCategories", reflect.TypeOf((*MockStore)(nil).GetAllCategories), arg0)
}

// GetBudgetByCategoryID mocks base method.
func (m *MockStore) GetBudgetByCategoryID(arg0 context.Context, arg1 int64) (db.Budget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBudgetByCategoryID", arg0, arg1)
	ret0, _ := ret[0].(db.Budget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBudgetByCategoryID indicates an expected call of GetBudgetByCategoryID.
func (mr *MockStoreMockRecorder) GetBudgetByCategoryID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBudgetByCategoryID", reflect.TypeOf((*MockStore)(nil).GetBudgetByCategoryID), arg0, arg1)
}

// GetBudgetByID mocks base method.
func (m *MockStore) GetBudgetByID(arg0 context.Context, arg1 int64) (db.Budget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBudgetByID", arg0, arg1)
	ret0, _ := ret[0].(db.Budget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBudgetByID indicates an expected call of GetBudgetByID.
func (mr *MockStoreMockRecorder) GetBudgetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBudgetByID", reflect.TypeOf((*MockStore)(nil).GetBudgetByID), arg0, arg1)
}

// GetBudgetsByWalletID mocks base method.
func (m *MockStore) GetBudgetsByWalletID(arg0 context.Context, arg1 int64) ([]db.Budget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBudgetsByWalletID", arg0, arg1)
	ret0, _ := ret[0].([]db.Budget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBudgetsByWalletID indicates an expected call of GetBudgetsByWalletID.
func (mr *MockStoreMockRecorder) GetBudgetsByWalletID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBudgetsByWalletID", reflect.TypeOf((*MockStore)(nil).GetBudgetsByWalletID), arg0, arg1)
}

// GetCategoryByID mocks base method.
func (m *MockStore) GetCategoryByID(arg0 context.Context, arg1 int64) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryByID", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoryByID indicates an expected call of GetCategoryByID.
func (mr *MockStoreMockRecorder) GetCategoryByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryByID", reflect.TypeOf((*MockStore)(nil).GetCategoryByID), arg0, arg1)
}

// GetExpense mocks base method.
func (m *MockStore) GetExpense(arg0 context.Context, arg1 db.GetExpenseParams) (db.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExpense", arg0, arg1)
	ret0, _ := ret[0].(db.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExpense indicates an expected call of GetExpense.
func (mr *MockStoreMockRecorder) GetExpense(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExpense", reflect.TypeOf((*MockStore)(nil).GetExpense), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetWallet mocks base method.
func (m *MockStore) GetWallet(arg0 context.Context, arg1 int64) (db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWallet", arg0, arg1)
	ret0, _ := ret[0].(db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWallet indicates an expected call of GetWallet.
func (mr *MockStoreMockRecorder) GetWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWallet", reflect.TypeOf((*MockStore)(nil).GetWallet), arg0, arg1)
}

// ListExpenses mocks base method.
func (m *MockStore) ListExpenses(arg0 context.Context, arg1 db.ListExpensesParams) ([]db.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListExpenses", arg0, arg1)
	ret0, _ := ret[0].([]db.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListExpenses indicates an expected call of ListExpenses.
func (mr *MockStoreMockRecorder) ListExpenses(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListExpenses", reflect.TypeOf((*MockStore)(nil).ListExpenses), arg0, arg1)
}

// ListWallets mocks base method.
func (m *MockStore) ListWallets(arg0 context.Context, arg1 db.ListWalletsParams) ([]db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWallets", arg0, arg1)
	ret0, _ := ret[0].([]db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWallets indicates an expected call of ListWallets.
func (mr *MockStoreMockRecorder) ListWallets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWallets", reflect.TypeOf((*MockStore)(nil).ListWallets), arg0, arg1)
}

// UpdateBudget mocks base method.
func (m *MockStore) UpdateBudget(arg0 context.Context, arg1 db.UpdateBudgetParams) (db.Budget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBudget", arg0, arg1)
	ret0, _ := ret[0].(db.Budget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBudget indicates an expected call of UpdateBudget.
func (mr *MockStoreMockRecorder) UpdateBudget(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBudget", reflect.TypeOf((*MockStore)(nil).UpdateBudget), arg0, arg1)
}

// UpdateCategory mocks base method.
func (m *MockStore) UpdateCategory(arg0 context.Context, arg1 db.UpdateCategoryParams) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCategory indicates an expected call of UpdateCategory.
func (mr *MockStoreMockRecorder) UpdateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategory", reflect.TypeOf((*MockStore)(nil).UpdateCategory), arg0, arg1)
}

// UpdateExpense mocks base method.
func (m *MockStore) UpdateExpense(arg0 context.Context, arg1 db.UpdateExpenseParams) (db.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExpense", arg0, arg1)
	ret0, _ := ret[0].(db.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateExpense indicates an expected call of UpdateExpense.
func (mr *MockStoreMockRecorder) UpdateExpense(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExpense", reflect.TypeOf((*MockStore)(nil).UpdateExpense), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}
