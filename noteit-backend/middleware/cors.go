package middleware

import (
	"github.com/savsgio/atreugo/v11"
)

func CORSMiddleware() atreugo.Middleware {
	return func(ctx *atreugo.RequestCtx) error {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if string(ctx.Method()) == "OPTIONS" {
			ctx.SetStatusCode(204)
			return nil
		}

		return ctx.Next()
	}
}
