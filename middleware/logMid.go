package middleware

import (
	"context"
	"ganxue-server/utils/log"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

func LogMid() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		start := time.Now()
		path := c.Request.URI().PathOriginal()

		c.Next(ctx)
		duration := time.Since(start)

		log.Debug("Request: ", string(path), " , Duration: ", duration)
	}
}
