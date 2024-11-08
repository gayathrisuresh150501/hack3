package middleware

import (
	"net/http"

	"github.com/savsgio/atreugo/v11"
)

var userID string = "user_1"

func AuthMiddleware() atreugo.Middleware {
	return func(ctx *atreugo.RequestCtx) error {
		// userID = ctx.UserValue("user_id")
		if userID == "" {
			return ctx.JSONResponse(map[string]string{"error": "Unauthorized"}, http.StatusUnauthorized)
		}
		return ctx.Next()
	}
}
