package auth

import (
	"context"
	"ganxue-server/global"
	"ganxue-server/model/md_model"
	"ganxue-server/utils/db/mongodb"
	"github.com/cloudwego/hertz/pkg/app"
	"go.mongodb.org/mongo-driver/bson"
)

// TODO: 更新上次学习章节

func GetDocs() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		id := ctx.Query("id")
		if id == "" {
			ctx.JSON(400, map[string]string{
				"message": "id不能为空",
			})
			return
		}
		var md md_model.Markdown
		// 查询数据库
		if err := mongodb.Find(global.MD, bson.M{"id": id}, &md); err != nil {
			ctx.JSON(500, map[string]string{
				"message": "服务器错误",
			})
			return
		}
		ctx.JSON(200, map[string]interface{}{
			"message": "查询成功",
			"data":    md,
		})
	}
}
