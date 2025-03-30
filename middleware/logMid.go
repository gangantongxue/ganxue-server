package middleware

import (
	"context"
	"ganxue-server/utils/log"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

func LogMid() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		start := time.Now()
		path := ctx.Request.URI().PathOriginal()

		ctx.Next(c)
		duration := time.Since(start)

		log.Debug("Request: ", string(path), " , Duration: ", duration)
	}
}
