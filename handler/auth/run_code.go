package auth

//func Init() {
//	cmd := exec.Command("docker", "run", "-d", "--name", "sandbox", "--network", "none", "--memory", "128m", "--cpus", "0.5", "go_code_executor", "sh", "-c", "trap 'tail -f /dev/null' INT TERM")
//	if err := cmd.Run(); err != nil {
//		log.Fatal("Failed to start sandbox container:", err)
//	}
//}
//
//type PlaygroundRequest struct {
//	Body string `json:"body"` // 要执行的 Go 代码
//}
//
//type PlaygroundResponse struct {
//	Errors string `json:"Errors"`
//	Events []struct {
//		Message string `json:"Message"`
//	} `json:"Events"`
//}
//
//func RunCode() app.HandlerFunc {
//	return func(c context.Context, ctx *app.RequestContext) {
//		data := struct {
//			ID   string `json:"id"`
// 			Group string `json:"group"`
//			Code string `json:"code"`
//		}{}
//		if err := ctx.Bind(&data); err != nil {
//			ctx.JSON(400, map[string]string{"message": "解析请求体失败"})
//			return
//		}
//
//		result, _err, _ := global.GROUP.Do("ans"+data.ID, func() (interface{}, error) {
//			var result answer_model.Answer
//			if err := mongodb.Find(global.ANSWER, bson.M{"id": data.ID}, &result); err != nil {
//				return nil, err.ToError()
//			}
//			return result, nil
//		})
//		if _err != nil {
//			ctx.JSON(500, map[string]string{
//				"message": "获取答案内容失败",
//			})
//			return
//		}
//
//		ans := result.(answer_model.Answer)
//
//		isRight, err := run(c, data.Code, &ans)
//		if err != nil {
//			ctx.JSON(500, map[string]string{
//				"message": err.ToString(),
//			})
//			return
//		}
//		if isRight {
//			ctx.JSON(200, map[string]string{
//				"message": "运行成功",
//				"data":    ans.Output,
//			})
//		} else {
//			ctx.JSON(206, map[string]string{
//				"message": "答案错误",
//				"data":    ans.Output,
//			})
//		}
//	}
//}
//
//func run(c context.Context, code string, ans *answer_model.Answer) (bool, *merr.Error) {
//
//}
