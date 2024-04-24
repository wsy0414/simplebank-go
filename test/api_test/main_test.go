package apitest

import (
	"os"
	"simplebank/api"
	"simplebank/db/sqlc"
	"testing"

	"github.com/gin-gonic/gin"
)

func NewTestServer(store sqlc.Store) *gin.Engine {
	return api.NewServer(store)
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
