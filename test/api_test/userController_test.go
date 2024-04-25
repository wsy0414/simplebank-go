package apitest

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"simplebank/api/model"
	mockdb "simplebank/db/mock"
	"simplebank/db/sqlc"
	"simplebank/util"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetUserInfo(t *testing.T) {

	testCases := []struct {
		Name          string
		setupAuth     func(t *testing.T, req *http.Request)
		mockFunc      func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, response *httptest.ResponseRecorder)
	}{
		{
			Name: "OK",
			setupAuth: func(t *testing.T, req *http.Request) {
				addAuth(t, req, 1)
			},
			mockFunc: func(store *mockdb.MockStore) {
				store.EXPECT().GetUser(gomock.Any(), gomock.Eq(int32(1))).Times(1).Return(
					sqlc.User{
						ID:       1,
						Name:     "ivan",
						Password: "$2a$13$sbz/G4gmGoB5X7Kw.y4KtOMJp286TG2BquEz5QK9z6TBXdV8ULGsu",
					},
					nil,
				)
				store.EXPECT().GetBalanceByUser(gomock.Any(), gomock.Eq(int32(1))).Times(1).Return(
					[]sqlc.Balance{
						sqlc.Balance{
							UserID:   1,
							Currency: "TWD",
							Balance:  "20",
						},
					}, nil,
				)
			},
			checkResponse: func(t *testing.T, response *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, response.Code)
				var user model.GetUserInfoResponse
				err := json.Unmarshal(response.Body.Bytes(), &user)
				require.NoError(t, err)
				require.Equal(t, 1, user.ID)
				require.Equal(t, "ivan", user.Name)
			},
		}, {
			Name: "BadRequest",
			setupAuth: func(t *testing.T, req *http.Request) {
				addAuth(t, req, 0)
			},
			mockFunc: func(store *mockdb.MockStore) {
				store.EXPECT().GetUser(gomock.Any(), gomock.Eq(int32(0))).Times(1).Return(sqlc.User{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, response *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, response.Code)
			},
		},
	}

	ctrl := gomock.NewController(t)
	// 這需要記得關掉
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	server := NewTestServer(store)

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			recoder := httptest.NewRecorder()
			testCase.mockFunc(store)

			request, err := http.NewRequest(http.MethodGet, "/user", nil)
			require.NoError(t, err)
			testCase.setupAuth(t, request)

			server.ServeHTTP(recoder, request)
			testCase.checkResponse(t, recoder)
		})
	}
}

func addAuth(t *testing.T, request *http.Request, userId int) {
	token, err := util.GenerateToken(userId, util.TOKEN_DEFAULT_DURATION)
	require.NoError(t, err)

	request.Header.Set("authorization", fmt.Sprintf("%s %s", "bearer", token))
}
