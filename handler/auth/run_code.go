package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"ganxue-server/global"
	"ganxue-server/model/answer_model"
	"ganxue-server/utils/db/mongodb"
	"ganxue-server/utils/error"
	"github.com/cloudwego/hertz/pkg/app"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"net/http"
)

type PlaygroundRequest struct {
	Code  string `json:"code"`            // 要执行的 Go 代码
	Input string `json:"input,omitempty"` // 控制台输入（可选）
}

type PlaygroundResponse struct {
	Output string `json:"output"` // 执行结果的输出
	Error  string `json:"error"`  // 错误信息（如果有）
}

func RunCode() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		data := struct {
			ID   string `json:"id"`
			Code string `json:"code"`
		}{}
		if err := ctx.Bind(&data); err != nil {
			ctx.JSON(400, map[string]string{"message": "解析请求体失败"})
			return
		}
		ans := answer_model.Answer{}
		if err := mongodb.Find(global.ANSWER, bson.M{"id": data.ID}, &ans); err != nil {
			ctx.JSON(400, map[string]string{"message": "查询答案失败"})
			return
		}
		isRight, err := run(data.Code, ans.Input, ans.Output)
		if err != nil {
			if err.Code == error.CodeError {
				ctx.JSON(206, map[string]string{
					"message": "答案错误",
					"data":    err.Message,
				})
				return
			} else {
				ctx.JSON(500, map[string]string{"message": "服务器错误"})
				return
			}
		}
		if isRight {
			ctx.JSON(200, map[string]string{"message": "答案正确"})
			return
		} else {
			ctx.JSON(206, map[string]string{"message": "答案错误"})
			return
		}
	}
}

func run(code string, input string, output string) (bool, *error.Error) {
	reqData := PlaygroundRequest{
		Code:  code,
		Input: input,
	}
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return false, error.New(error.ParseJsonError, err, "")
	}
	resp, err := http.Post(
		"https://play.golang.org/api/playground/run",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return false, error.New(error.RunCodeError, err, "")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, error.New(error.RunCodeError, err, "")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, error.New(error.RunCodeError, err, "")
	}

	var respData PlaygroundResponse
	if err := json.Unmarshal(body, &respData); err != nil {
		return false, error.New(error.RunCodeError, err, "")
	}
	if respData.Error != "" {
		return false, error.New(error.CodeError, err, respData.Error)
	}
	if respData.Output != output {
		return false, error.New(error.CodeError, err, respData.Output)
	}
	return true, nil
}
