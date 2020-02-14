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

var debugLogger *log.Logger = log.New(os.Stdout, time.Now().Format("2006-01-02 15:04:05.999Z")+" [DEBUG] ", 0)
var infoLogger *log.Logger = log.New(os.Stdout, time.Now().Format("2006-01-02 15:04:05.999Z")+" [INFO] ", 0)
var warnLogger *log.Logger = log.New(os.Stdout, time.Now().Format("2006-01-02 15:04:05.999Z")+" [WARN] ", 0)
var errorLogger *log.Logger = log.New(os.Stdout, time.Now().Format("2006-01-02 15:04:05.999Z")+" [ERROR] ", 0)

// The following format:
// [randomApp] 2006-01-02 15:04:05.999Z [DEBUG] [log_test.go:32] bug me not!
func Debug(x ...interface{}) {
	_, f, l, _ := runtime.Caller(1)
	s := strings.Split(f, "/")
	var y []interface{}
	y = append(y, fmt.Sprintf("[%s:%v]", s[len(s)-1], l))
	y = append(y, x...)
	debugLogger.Println(y...)
}

// Same as Debug, but uses INFO
func Info(x ...interface{}) {
	_, f, l, _ := runtime.Caller(1)
	s := strings.Split(f, "/")
	var y []interface{}
	y = append(y, fmt.Sprintf("[%s:%v]", s[len(s)-1], l))
	y = append(y, x...)
	infoLogger.Println(y...)
}

// Same as Debug, but uses WARN
func Warn(x ...interface{}) {
	_, f, l, _ := runtime.Caller(1)
	s := strings.Split(f, "/")
	var y []interface{}
	y = append(y, fmt.Sprintf("[%s:%v]", s[len(s)-1], l))
	y = append(y, x...)
	warnLogger.Println(y...)
}

// Same as Debug, but uses ERROR
func Error(x ...interface{}) {
	_, f, l, _ := runtime.Caller(1)
	s := strings.Split(f, "/")
	var y []interface{}
	y = append(y, fmt.Sprintf("[%s:%v]", s[len(s)-1], l))
	y = append(y, x...)
	errorLogger.Println(y...)
}

// For backward compatability
// Same as INFO
func Println(x ...interface{}) {
	_, f, l, _ := runtime.Caller(1)
	s := strings.Split(f, "/")
	var y []interface{}
	y = append(y, fmt.Sprintf("[%s:%v]", s[len(s)-1], l))
	y = append(y, x...)
	infoLogger.Println(y...)
}

// For backward compatability
// Same as INFO
func Printf(str string, x ...interface{}) {
	_, f, l, _ := runtime.Caller(1)
	s := strings.Split(f, "/")
	var y []interface{}
	y = append(y, fmt.Sprintf("[%s:%v]", s[len(s)-1], l))
	y = append(y, fmt.Sprintf(str, x...))
	infoLogger.Println(y...)
	// Info(fmt.Sprintf(s, x...))
}

// WriterConfig: Set config for what to write onto the log
type WriterConfig struct {
	AppName    string
	TimeFormat string
	Prefix     string
}

// Init: Used to change the configuration of the logger
// See source code for default values of config
func Init(w io.Writer, wc WriterConfig) {
	if wc.TimeFormat == "" {
		wc.TimeFormat = "2006-01-02 15:04:05.999Z"
	}
	if wc.AppName == "" {
		wc.AppName = "[APPX] "
	}

	debugLogger = log.New(w, fmt.Sprintf("[%s] ", wc.AppName)+time.Now().Format(wc.TimeFormat)+" [DEBUG] "+fmt.Sprintf("%s", wc.Prefix), 0)
	infoLogger = log.New(w, fmt.Sprintf("[%s] ", wc.AppName)+time.Now().Format(wc.TimeFormat)+" [INFO] "+fmt.Sprintf("%s", wc.Prefix), 0)
	warnLogger = log.New(w, fmt.Sprintf("[%s] ", wc.AppName)+time.Now().Format(wc.TimeFormat)+" [WARN] "+fmt.Sprintf("%s", wc.Prefix), 0)
	errorLogger = log.New(w, fmt.Sprintf("[%s] ", wc.AppName)+time.Now().Format(wc.TimeFormat)+" [ERROR] "+fmt.Sprintf("%s", wc.Prefix), 0)
}
