package controller

import (
	"context"
	"net/http"
	"simplebank/db"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	conn, err := db.GetQuery()
	if err != nil {
		c.String(http.StatusInternalServerError, "database error")
	}

	defer conn.Conn.Close()

	account, err := conn.Q.GetAccount(context.Background(), 1)
	if err != nil {
		c.String(http.StatusInternalServerError, "database error")
	}

	c.JSON(http.StatusOK, account)
}
