package auth

import (
	"context"
	"ganxue-server/utils/db/mysql"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

func UserInfo() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		userID := c.Value("userID")
		if userID == nil {
			ctx.AbortWithStatus(401)
			return
		}
		user, err := mysql.FindUserAndInfoByID(userID.(uint))
		if err != nil {
			ctx.AbortWithStatus(401)
			return
		}

		// 判断上次更新时间是否为昨天
		now := time.Now()
		todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		yesterdayStart := todayStart.AddDate(0, 0, -1)

		if user.UserInfo.UpdatedAt.After(todayStart) || user.UserInfo.UpdatedAt.Equal(todayStart) {
			// today
		} else if user.UserInfo.UpdatedAt.After(yesterdayStart) || user.UserInfo.UpdatedAt.Equal(yesterdayStart) {
			// yesterday
			user.UserInfo.StreakDays++
			user.UserInfo.TotalDays++
		} else {
			// other
			user.UserInfo.StreakDays = 1
			user.UserInfo.TotalDays++
		}

		if err := mysql.Update(user.UserInfo); err != nil {
			ctx.JSON(500, map[string]string{"message": "服务器错误"})
			return
		}

		ctx.JSON(200, map[string]interface{}{
			"message": "success",
			"data": map[string]interface{}{
				"user_info": map[string]interface{}{
					"email":     user.Email,
					"user_name": user.UserName,
				},
				"study_stats": map[string]interface{}{
					"streak_days": user.UserInfo.StreakDays,
					"total_days":  user.UserInfo.TotalDays,
					"last_time":   user.UserInfo.LastTime,
				},
			},
		})
	}
}
