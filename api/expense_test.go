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
	category := RandomCategory()
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
				"expense_date":        expense.ExpenseDate.Truncate(24 * time.Hour),
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
					ExpenseDate:        expense.ExpenseDate.Truncate(24 * time.Hour),
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
				"expense_date":        expense.ExpenseDate.Truncate(24 * time.Hour),
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
				"expense_date":        expense.ExpenseDate.Truncate(24 * time.Hour),
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
					ExpenseDate:        expense.ExpenseDate.Truncate(24 * time.Hour),
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
				"expense_date":        expense.ExpenseDate.Truncate(24 * time.Hour),
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
				"expense_date":        expense.ExpenseDate.Truncate(24 * time.Hour),
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
					ExpenseDate:        expense.ExpenseDate.Truncate(24 * time.Hour),
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

func TestGetExpenseAPI(t *testing.T) {
	user, _ := randomUser(t)
	wallet := RandomWallet(user.Username)
	category := RandomCategory()
	expense := RandomExpense(wallet.ID, category.ID)

	testCases := []struct {
		name          string
		walletID      int64
		expenseID     int64
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			walletID:  wallet.ID,
			expenseID: expense.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
					AnyTimes().
					Return(wallet, nil)

				arg := db.GetExpenseParams{
					ID:       expense.ID,
					WalletID: wallet.ID,
				}

				store.EXPECT().
					GetExpense(gomock.Any(), gomock.Eq(arg)).
					AnyTimes().
					Return(expense, nil) // Убедитесь, что возвращается корректное значение
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchExpense(t, recorder.Body, expense)
			},
		},
		{
			name:      "Unauthorized",
			walletID:  wallet.ID,
			expenseID: expense.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, "unauthorized_user", time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
					Times(1).
					Return(wallet, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:      "NotFound",
			walletID:  wallet.ID,
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
					GetExpense(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Expense{}, sql.ErrNoRows) // Simulating not found
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "InternalError",
			walletID:  wallet.ID,
			expenseID: expense.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
					Return(wallet, nil).
					Times(1)

				store.EXPECT().
					GetExpense(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Expense{}, sql.ErrConnDone) // Simulating an internal error
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "InvalidID",
			walletID:  wallet.ID,
			expenseID: 0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Any()).
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

			url := fmt.Sprintf("/wallets/%d/expenses/%d", tc.walletID, tc.expenseID)
			fmt.Println(url)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			fmt.Println(recorder.Body)
			tc.checkResponse(t, recorder)
		})
	}
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
		ID:       util.RandomInt(1, 1000),
		WalletID: walletID,
		Amount:   util.RandomInt(1, 1000),
		ExpenseDescription: sql.NullString{
			String: util.RandomString(6),
			Valid:  true,
		},
		CategoryID:  CategoryID,
		ExpenseDate: time.Now(),
	}
}
