package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func JwtMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// TODO: jwt middleware
	}
}
