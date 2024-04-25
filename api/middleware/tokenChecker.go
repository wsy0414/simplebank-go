package middleware

import (
	"errors"
	"log"
	"net/http"
	"simplebank/util"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("enter in middleware")
		jwtToken := ctx.GetHeader("authorization")
		tokens := strings.Fields(jwtToken)
		if len(tokens) == 0 {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("must have authorization header"))
			return
		}
		if strings.ToLower(tokens[0]) != "bearer" {
			ctx.AbortWithError(http.StatusForbidden, errors.New("authorization type not bearer"))
			return
		}

		claims, err := util.ParseToken(tokens[1])
		if err != nil {
			ctx.AbortWithError(http.StatusForbidden, err)
			return
		}

		now := time.Now()
		if claims.StandardClaims.ExpiresAt < now.Unix() {
			ctx.AbortWithError(http.StatusForbidden, errors.New("token is over date"))
			return
		}
		log.Println("claims.ID: ", claims.ID)
		ctx.Set("userId", claims.ID)
		token, err := util.GenerateToken(claims.ID, util.TOKEN_DEFAULT_DURATION)
		if err != nil {
			log.Println("here???")
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.Header("authorization", token)

		ctx.Next()
	}
}
