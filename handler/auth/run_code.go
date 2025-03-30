package auth

import (
	"bytes"
	"context"
	"ganxue-server/global"
	"ganxue-server/model/answer_model"
	"ganxue-server/utils/db/mongodb"
	"github.com/cloudwego/hertz/pkg/app"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"os/exec"
)

func RunCode() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// 获取请求体
		userCode := struct {
			ID   string `json:"id"`
			Code string `json:"code"`
		}{}
		if err := ctx.Bind(&userCode); err != nil {
			ctx.JSON(400, map[string]string{"message": "解析请求体失败"})
			return
		}
		var ans answer_model.Answer
		if err := mongodb.Find(global.GO_ANSWER, bson.M{"id": userCode.ID}, &ans); err != nil {
			ctx.JSON(400, map[string]string{"message": "查询数据失败"})
			return
		}
		ctx.JSON(100, map[string]string{
			"message": "等待运行中...",
		})

		// 将用户代码写入文件中
		if err := os.WriteFile("/home/ganxue-server/utils/run_code/user_code.go", []byte(userCode.Code), 0644); err != nil {
			ctx.JSON(500, map[string]string{"message": "写入文件失败"})
			return
		}

		// 执行用户代码
		cmd := exec.Command("go", "run", "/home/ganxue-server/utils/run_code/.")
		// 交互式输入
		stdin, err := cmd.StdinPipe()
		if err != nil {
			ctx.JSON(500, map[string]string{"message": "创建stdin失败"})
			return
		}

		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		if err := cmd.Start(); err != nil {
			ctx.JSON(500, map[string]string{"message": "执行用户代码失败"})
			return
		}
		if _, err := stdin.Write([]byte(ans.Input)); err != nil {
			ctx.JSON(500, map[string]string{"message": "写入标准输入失败"})
			return
		}
		if err := stdin.Close(); err != nil {
			ctx.JSON(500, map[string]string{"message": "关闭标准输入失败"})
			return
		}

		if err := cmd.Wait(); err != nil {
			ctx.JSON(206, map[string]string{
				"message": "执行出错",
				"data":    stderr.String() + "\n" + err.Error(),
			})
			return
		}

		actualOut := stdout.String()
		if stderr.String() != "" {
			ctx.JSON(206, map[string]string{
				"message": "执行出错",
				"error":   stderr.String(),
			})
			return
		}
		if actualOut != ans.Output {
			ctx.JSON(206, map[string]string{
				"message": "结果错误",
				"data":    actualOut,
			})
			return
		} else {
			ctx.JSON(200, map[string]string{
				"message": "运行成功",
				"data":    actualOut,
			})
		}
	}
}
