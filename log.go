package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

// var debugLogger *log.Logger
// var infoLogger *log.Logger
// var warnLogger *log.Logger
// var errorLogger *log.Logger

// func x() {
// 	_, f, l, _ := runtime.Caller(1)
// 	s := strings.Split(f, "/")

// 	debugLogger = log.New(os.Stdout, time.Now().Format("2006-01-02 15:04:05.999Z")+" [DEBUG] "+fmt.Sprintf("[%s:%v] ", s[len(s)-1], l), 0)
// 	infoLogger = log.New(os.Stdout, time.Now().Format("2006-01-02 15:04:05.999Z")+" [INFO] "+fmt.Sprintf("[%s:%v] ", s[len(s)-1], l), 0)
// 	warnLogger = log.New(os.Stdout, time.Now().Format("2006-01-02 15:04:05.999Z")+" [WARN] "+fmt.Sprintf("[%s:%v] ", s[len(s)-1], l), 0)
// 	errorLogger = log.New(os.Stdout, time.Now().Format("2006-01-02 15:04:05.999Z")+" [ERROR] "+fmt.Sprintf("[%s:%v] ", s[len(s)-1], l), 0)
// }
// x()

var debugLogger *log.Logger = log.New(os.Stdout, time.Now().Format("2006-01-02 15:04:05.999Z")+" [DEBUG] ", 0)
var infoLogger *log.Logger = log.New(os.Stdout, time.Now().Format("2006-01-02 15:04:05.999Z")+" [INFO] ", 0)
var warnLogger *log.Logger = log.New(os.Stdout, time.Now().Format("2006-01-02 15:04:05.999Z")+" [WARN] ", 0)
var errorLogger *log.Logger = log.New(os.Stdout, time.Now().Format("2006-01-02 15:04:05.999Z")+" [ERROR] ", 0)

func Debug(x ...interface{}) {
	_, f, l, _ := runtime.Caller(1)
	s := strings.Split(f, "/")
	var y []interface{}
	y = append(y, fmt.Sprintf("[%s:%v]", s[len(s)-1], l))
	y = append(y, x...)
	debugLogger.Println(y...)
}

func Info(x ...interface{}) {
	_, f, l, _ := runtime.Caller(1)
	s := strings.Split(f, "/")
	var y []interface{}
	y = append(y, fmt.Sprintf("[%s:%v]", s[len(s)-1], l))
	y = append(y, x...)
	infoLogger.Println(y...)
}

func Warn(x ...interface{}) {
	_, f, l, _ := runtime.Caller(1)
	s := strings.Split(f, "/")
	var y []interface{}
	y = append(y, fmt.Sprintf("[%s:%v]", s[len(s)-1], l))
	y = append(y, x...)
	warnLogger.Println(y...)
}

func Error(x ...interface{}) {
	_, f, l, _ := runtime.Caller(1)
	s := strings.Split(f, "/")
	var y []interface{}
	y = append(y, fmt.Sprintf("[%s:%v]", s[len(s)-1], l))
	y = append(y, x...)
	errorLogger.Println(y...)
}

// WriterConfig: The following config options are available to be set
type WriterConfig struct {
	TimeFormat string
}

// Init: used to change the configuration of the logger
func Init(w io.Writer, wc WriterConfig) {
	if wc.TimeFormat == "" {
		wc.TimeFormat = "2006-01-02 15:04:05.999Z"
	}

	// fmt.Println(fmt.Sprintf("%s:%v", s[len(s)-1], l))
	debugLogger = log.New(w, time.Now().Format(wc.TimeFormat)+" [DEBUG] ", 0)
	infoLogger = log.New(w, time.Now().Format(wc.TimeFormat)+" [INFO] ", 0)
	warnLogger = log.New(w, time.Now().Format(wc.TimeFormat)+" [WARN] ", 0)
	errorLogger = log.New(w, time.Now().Format(wc.TimeFormat)+" [ERROR] ", 0)
}
