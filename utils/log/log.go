package log

//TODO: config

import (
	"context"
	"fmt"
	"ganxue-server/global"
	err "ganxue-server/utils/error"
	"strconv"
	"time"
)

var ctx = context.Background()
var Signal = make(chan uint8, 100)

const (
	DebugKey uint8 = 0
	WarnKey  uint8 = 1
	ErrorKey uint8 = 2
	PanicKey uint8 = 3
	FatalKey uint8 = 4
)

var logKey = map[uint8]string{
	DebugKey: "debug",
	WarnKey:  "warn",
	ErrorKey: "error",
	PanicKey: "panic",
	FatalKey: "fatal",
}

const (
	Reset  = "\033[0m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Red    = "\033[31m"
	Blue   = "\033[34m"
)

func Init() {
	go processLog()
}

// Parse 解析 interface{}
func Parse(v []interface{}) string {
	var str string
	now := time.Now().Format("2006-01-02 15:04:05")
	str += Blue + now + Reset + " "
	for _, v := range v {
		switch v.(type) {
		case string:
			str += v.(string)
		case int:
			str += strconv.Itoa(v.(int))
		case float64:
			str += fmt.Sprintf("%f", v.(float64))
		case bool:
			str += strconv.FormatBool(v.(bool))
		case err.Error:
			d := v.(err.Error)
			str += d.ToString()
		default:
			str += fmt.Sprintf("%v", v)
		}
	}
	return str
}

// Debug 打印日志
func Debug(v ...interface{}) {
	str := Parse(v)
	global.RDB.RPush(ctx, "debug", str)
	Signal <- DebugKey
}

// Warn 警告日志
func Warn(v ...interface{}) {
	str := Parse(v)
	global.RDB.RPush(ctx, "warn", str)
	Signal <- WarnKey
}

// Error 失败日志
func Error(v ...interface{}) {
	str := Parse(v)
	global.RDB.RPush(ctx, "error", str)
	Signal <- ErrorKey
}

// Panic 严重错误日志
func Panic(v ...interface{}) {
	str := Parse(v)
	global.RDB.RPush(ctx, "panic", str)
	Signal <- PanicKey
}

// Fatal 严重错误日志
func Fatal(v ...interface{}) {
	str := Parse(v)
	global.RDB.RPush(ctx, "fatal", str)
	Signal <- FatalKey
}
