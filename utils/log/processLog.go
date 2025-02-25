package log

import (
	"context"
	"fmt"
	"ganxue-server/global"
	"gopkg.in/mail.v2"
	"os"
	"path/filepath"
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
		if err != nil {
			return
		}
		switch signal {
		case DebugKey:
			fmt.Println("[debug]", log)
		case WarnKey:
			fmt.Println("[warn]", log)
		case ErrorKey:
			fmt.Println("[error]", log)
			sendEmail("[error] " + log)
			_, err := logFile.WriteString("[error] " + log + "\n")
			if err != nil {
				return
			}
		case PanicKey:
			sendEmail("[panic] " + log)
			_, err := logFile.WriteString("[panic] " + log + "\n")
			if err != nil {
				return
			}
			panic("[panic] " + log)
		case FatalKey:
			fmt.Println(log)
			sendEmail("[fatal] " + log)
			_, err := logFile.WriteString("[fatal] " + log + "\n")
			if err != nil {
				return
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
