package log

import (
	"context"
	"errors"
	"fmt"
	"ganxue-server/global"
	"github.com/go-redis/redis/v8"
	"gopkg.in/mail.v2"
	"os"
	"path/filepath"
	"regexp"
)

func processLog() {
	// 打开logFile文件
	wd, _err := os.Getwd()
	if _err != nil {
		return
	}
	filePath := filepath.Join(wd, "utils", "log", "logFile")
	logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			return
		}
	}(logFile)
	for {
		signal := <-Signal
		// 从redis中获取日志
		log, err := global.RDB.LPop(context.Background(), logKey[signal]).Result()
		if errors.Is(err, redis.Nil) {
			fmt.Println("Redis 列表为空或键不存在:", logKey[signal])
			continue
		} else if err != nil {
			fmt.Println("redis获取日志失败:", err)
			return
		}

		switch signal {
		case DebugKey:
			fmt.Println(Green+"[debug]"+Reset, log)
		case WarnKey:
			fmt.Println(Yellow+"[warn]"+Reset, log)
		case ErrorKey:
			fmt.Println(Red+"[error]"+Reset, log)
			_log := removeANSIEscapeCodes(log)
			sendEmail("[error] " + _log)
			_, err := logFile.WriteString("[error] " + _log + "\n")
			if err != nil {
				return
			}
		case PanicKey:
			_log := removeANSIEscapeCodes(log)
			sendEmail("[panic] " + _log)
			_, err := logFile.WriteString("[panic] " + _log + "\n")
			if err != nil {
				return
			}
			panic("[panic] " + log)
		case FatalKey:
			fmt.Println(log)
			_log := removeANSIEscapeCodes(log)
			sendEmail("[fatal] " + log)
			_, err := logFile.WriteString("[fatal] " + _log + "\n")
			if err != nil {
				fmt.Println("写入日志失败", err)
			}
			os.Exit(1)
		}
	}
}

func sendEmail(log string) {
	m := mail.NewMessage()
	m.SetHeader("From", global.CONFIG.Email.From)
	m.SetHeader("To", "gangantongxue@outlook.com")
	m.SetHeader("Subject", "日志报警")
	m.SetBody("text/html", log)
	d := mail.NewDialer(global.CONFIG.Email.Host, global.CONFIG.Email.Port, global.CONFIG.Email.Username, global.CONFIG.Email.Password)
	d.SSL = true
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("发送邮件失败：", err)
	}
}

// removeANSIEscapeCodes 去除 ANSI 转义码
func removeANSIEscapeCodes(s string) string {
	ansiEscape := `[\x1b][[()#;?]*(?:[0-9]{1,4}(?:;[0-9]{0,4})*)?[0-9A-ORZcf-nqry=><~]`
	re := regexp.MustCompile(ansiEscape)
	return re.ReplaceAllString(s, "")
}
