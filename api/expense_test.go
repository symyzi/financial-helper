package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	db "github.com/symyzi/financial-helper/db/gen"
	mockdb "github.com/symyzi/financial-helper/db/mock"
	"github.com/symyzi/financial-helper/token"
	"github.com/symyzi/financial-helper/util"
)

func TestCreateExpenseAPI(t *testing.T) {
	user, _ := randomUser(t)
	wallet := RandomWallet(user.Username)
	category := RandomCategory(user.Username)
	expense := RandomExpense(wallet.ID, category.ID)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"wallet_id":           expense.WalletID,
				"amount":              expense.Amount,
				"expense_description": expense.ExpenseDescription,
				"category_id":         expense.CategoryID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(expense.WalletID)).
					Times(1).
					Return(wallet, nil)

				arg := db.CreateExpenseParams{
					WalletID:           expense.WalletID,
					Amount:             expense.Amount,
					ExpenseDescription: expense.ExpenseDescription,
					CategoryID:         expense.CategoryID,
				}
				store.EXPECT().
					CreateExpense(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(expense, nil)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"wallet_id":           expense.WalletID,
				"amount":              expense.Amount,
				"expense_description": expense.ExpenseDescription,
				"category_id":         expense.CategoryID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(expense.WalletID)).
					Times(1).
					Return(wallet, nil)
				store.EXPECT().
					CreateExpense(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Expense{}, sql.ErrConnDone)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name: "UnauthorizedUser",
			body: gin.H{
				"wallet_id":           expense.WalletID,
				"amount":              expense.Amount,
				"expense_description": expense.ExpenseDescription,
				"category_id":         expense.CategoryID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, "unauthorized_user1", time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(expense.WalletID)).
					AnyTimes().
					Return(wallet, nil)

				arg := db.CreateExpenseParams{
					WalletID:           expense.WalletID,
					Amount:             expense.Amount,
					ExpenseDescription: expense.ExpenseDescription,
					CategoryID:         expense.CategoryID,
				}

				store.EXPECT().
					CreateExpense(gomock.Any(), gomock.Eq(arg)).
					AnyTimes().
					Return(db.Expense{}, sql.ErrConnDone)
			},

			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recoder.Code)
			},
		},
		{
			name: "InvalidWalletID",
			body: gin.H{
				"wallet_id":           0,
				"amount":              expense.Amount,
				"expense_description": expense.ExpenseDescription,
				"category_id":         expense.CategoryID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Any()).
					AnyTimes().
					Return(db.Wallet{}, sql.ErrNoRows)

			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name: "InvalidAmount",
			body: gin.H{
				"wallet_id":           expense.WalletID,
				"amount":              "dfs",
				"expense_description": expense.ExpenseDescription,
				"category_id":         expense.CategoryID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(expense.WalletID)).
					AnyTimes().
					Return(wallet, nil)

				arg := db.CreateExpenseParams{
					WalletID:           expense.WalletID,
					Amount:             expense.Amount,
					ExpenseDescription: expense.ExpenseDescription,
					CategoryID:         expense.CategoryID,
				}

				store.EXPECT().
					CreateExpense(gomock.Any(), gomock.Eq(arg)).
					AnyTimes().
					Return(db.Expense{}, sql.ErrConnDone)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := fmt.Sprintf("/wallets/%d/expenses", wallet.ID)
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}

}

func TestGetExpenseByID(t *testing.T) {
	user, _ := randomUser(t)
	wallet := RandomWallet(user.Username)
	category := RandomCategory(user.Username)
	expense := RandomExpense(wallet.ID, category.ID)

	testCases := []struct {
		name          string
		expenseID     int64
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			expenseID: expense.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
					Times(1).
					Return(wallet, nil)

				store.EXPECT().
					GetExpense(gomock.Any(), gomock.Eq(expense.ID)).
					Times(1).
					Return(expense, nil)
			},

			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
				requireBodyMatchExpense(t, recoder.Body, expense)
			},
		},
		{
			name:      "Unauthorized",
			expenseID: expense.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, "unauthorized_user", time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
					Times(1).
					Return(wallet, nil)

				store.EXPECT().
					GetExpense(gomock.Any(), gomock.Eq(expense.ID)).
					Times(1).
					Return(expense, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:      "NotFound",
			expenseID: expense.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
					AnyTimes().
					Return(db.Wallet{}, sql.ErrNoRows)

				store.EXPECT().
					GetExpense(gomock.Any(), gomock.Eq(expense.ID)).
					Times(1).
					Return(db.Expense{}, sql.ErrNoRows)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "InternalError",
			expenseID: expense.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
					Return(wallet, nil).
					AnyTimes()

				store.EXPECT().
					GetExpense(gomock.Any(), gomock.Eq(expense.ID)).
					Times(1).
					Return(db.Expense{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "InvalidID",
			expenseID: 0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/wallets/%d/expenses/%d", wallet.ID, tc.expenseID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestListExpensesAPI(t *testing.T) {
	user, _ := randomUser(t)
	wallet := RandomWallet(user.Username)
	category := RandomCategory(user.Username)
	n := 5

	expenses := make([]db.Expense, n)

	for i := 0; i < n; i++ {
		expenses[i] = RandomExpense(wallet.ID, category.ID)
	}

	type Query struct {
		pageID   int
		pageSize int
	}

	testCases := []struct {
		name          string
		query         Query
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
					AnyTimes().
					Return(wallet, nil)

				arg := db.ListExpensesParams{
					Limit:  int32(n),
					Offset: 0,
				}

				store.EXPECT().
					ListExpenses(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(expenses, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchExpenses(t, recorder.Body, expenses)
			},
		},
		{
			name: "NoAuthorization",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalError",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
					AnyTimes().
					Return(wallet, nil)

				store.EXPECT().
					ListExpenses(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.Expense{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidPageID",
			query: Query{
				pageID:   -1,
				pageSize: n,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
					AnyTimes().
					Return(wallet, nil)

				store.EXPECT().
					ListExpenses(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidPageSize",
			query: Query{
				pageID:   1,
				pageSize: 100000,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
					AnyTimes().
					Return(wallet, nil)

				store.EXPECT().
					ListExpenses(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/wallets/%d/expenses", wallet.ID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			q := request.URL.Query()
			q.Add("page_id", fmt.Sprintf("%d", tc.query.pageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.query.pageSize))
			request.URL.RawQuery = q.Encode()

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestDeleteExpenseAPI(t *testing.T) {
	user, _ := randomUser(t)
	wallet := RandomWallet(user.Username)
	expense := RandomExpense(wallet.ID, RandomCategory(user.Username).ID)

	testCases := []struct {
		name          string
		expenseID     int64
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			expenseID: expense.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
					Times(1).
					Return(wallet, nil)

				store.EXPECT().
					GetExpense(gomock.Any(), gomock.Eq(expense.ID)).
					Times(1).
					Return(expense, nil)

				store.EXPECT().
					DeleteExpense(gomock.Any(), gomock.Eq(expense.ID)).
					Times(1).
					Return(nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:      "Unauthorized",
			expenseID: expense.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, "unauthorized_user", time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
					Times(1).
					Return(wallet, nil)

				store.EXPECT().
					GetExpense(gomock.Any(), gomock.Eq(expense.ID)).
					Times(1).
					Return(expense, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:      "NotFound",
			expenseID: expense.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
					AnyTimes().
					Return(wallet, nil)

				store.EXPECT().
					GetExpense(gomock.Any(), gomock.Eq(expense.ID)).
					Times(1).
					Return(db.Expense{}, sql.ErrNoRows)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "InternalError",
			expenseID: expense.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
					Times(1).
					Return(wallet, nil)

				store.EXPECT().
					GetExpense(gomock.Any(), gomock.Eq(expense.ID)).
					Times(1).
					Return(expense, nil)

				store.EXPECT().
					DeleteExpense(gomock.Any(), gomock.Eq(expense.ID)).
					Times(1).
					Return(sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "InvalidID",
			expenseID: 0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/wallets/%d/expenses/%d", wallet.ID, tc.expenseID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func requireBodyMatchExpenses(t *testing.T, body *bytes.Buffer, expenses []db.Expense) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotExpenses []db.Expense
	err = json.Unmarshal(data, &gotExpenses)
	require.NoError(t, err)
	require.Equal(t, expenses, gotExpenses)
}

func requireBodyMatchExpense(t *testing.T, body *bytes.Buffer, expense db.Expense) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotExpense db.Expense
	err = json.Unmarshal(data, &gotExpense)
	require.NoError(t, err)
	require.Equal(t, expense, gotExpense)
}

func RandomExpense(walletID int64, CategoryID int64) db.Expense {
	return db.Expense{
		ID:                 util.RandomInt(1, 1000),
		WalletID:           walletID,
		Amount:             util.RandomInt(1, 1000),
		ExpenseDescription: util.RandomString(12),
		CategoryID:         CategoryID,
	}
}
