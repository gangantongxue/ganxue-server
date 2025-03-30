package auth

import (
	"context"
	"ganxue-server/global"
	"ganxue-server/model/md_model"
	"ganxue-server/utils/db/mongodb"
	"ganxue-server/utils/db/mysql"
	"ganxue-server/utils/log"
	"github.com/cloudwego/hertz/pkg/app"
	"go.mongodb.org/mongo-driver/bson"
)

func GetDocs() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		id := ctx.Query("id")
		group := ctx.Query("group")
		if id == "" {
			ctx.JSON(400, map[string]string{
				"message": "id不能为空",
			})
			return
		}

		// 查询数据库
		md, _err, _ := global.GROUP.Do("md"+id, func() (interface{}, error) {
			var result md_model.Markdown
			if err := mongodb.Find(global.GO_MD, bson.M{"id": id}, &result); err != nil {
				return nil, err.ToError()
			}
			return result, nil
		})
		if _err != nil {
			ctx.JSON(400, map[string]string{
				"message": "获取文档内容失败",
			})
			return
		}
		ctx.JSON(200, map[string]interface{}{
			"message": "查询成功",
			"data":    md,
		})
		userID := c.Value("userID").(uint)

		userInfo, err := mysql.FindInfoByID(uint(userID))
		if err != nil {
			log.Error("用户信息查询失败", err)
		}

		switch group {
		case "golang":
			userInfo.GoLastChapter = id
		case "c":
			userInfo.CLastChapter = id
		case "cpp":
			userInfo.CppLastChapter = id
		}

		if err := mysql.Update(userInfo); err != nil {
			log.Error("用户信息更新失败", err)
		}
	}
}
