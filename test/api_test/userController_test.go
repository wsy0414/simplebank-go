package apitest

import (
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

	ctrl := gomock.NewController(t)

	// 這需要記得關掉
	defer ctrl.Finish()
	store := mockdb.NewMockStore(ctrl)

	server := NewTestServer(store)
	recoder := httptest.NewRecorder()

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

	request, err := http.NewRequest(http.MethodGet, "/user", nil)
	require.NoError(t, err)
	addAuth(t, request, 1)

	server.ServeHTTP(recoder, request)
	var user model.GetUserInfoResponse
	err = json.Unmarshal(recoder.Body.Bytes(), &user)
	require.NoError(t, err)

	fmt.Printf("%v", user)
}

func addAuth(t *testing.T, request *http.Request, userId int) {
	token, err := util.GenerateToken(userId, util.TOKEN_DEFAULT_DURATION)
	require.NoError(t, err)

	request.Header.Set("authorization", fmt.Sprintf("%s %s", "bearer", token))
}
