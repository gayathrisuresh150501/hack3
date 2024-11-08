package middleware

import (
	"github.com/savsgio/atreugo/v11"
)

func CORSMiddleware() atreugo.Middleware {
	return func(ctx *atreugo.RequestCtx) error {
		origin := string(ctx.Request.Header.Peek("Origin"))

		// Allow all origins in development, or specify allowed origins
		if origin != "" {
			ctx.Response.Header.Set("Access-Control-Allow-Origin", origin)
			ctx.Response.Header.Set("Vary", "Origin")
		}

		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if string(ctx.Method()) == "OPTIONS" {
			ctx.Response.Header.Set("Access-Control-Max-Age", "86400")
			ctx.SetStatusCode(200)
			return nil
		}

		return ctx.Next()
	}
}
