// @Author LiuYong
// @Created at 2021-01-31
// @Modified at 2021-01-31
package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Level int

var (
	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG   Level = iota // debug信息
	INFO                 // 普通信息
	WARNING              // 警告
	ERROR                // 错误
	FATAL                // 失败信息
)

// TODO logger的多线陈单例
func Setup() {
	// 获得日志记录器，第一个参数是IO句柄，第二个是每行日志的开头，第三个定义了日志记录属性
	logger = log.New(os.Stdout, DefaultPrefix, log.LstdFlags)

	Info(map[string]interface{}{
		"时间": time.Now(),
	})
}

// 设置前缀
func setPrefix(level Level) {
	// 返回当前程序栈中的文件和执行的行数
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}

func Debug(v ...interface{}) {
	if logger == nil {
		Setup()
	}
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	if logger == nil {
		Setup()
	}
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	if logger == nil {
		Setup()
	}
	setPrefix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	if logger == nil {
		Setup()
	}
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	if logger == nil {
		Setup()
	}
	setPrefix(FATAL)
	logger.Fatalln(v)
}
