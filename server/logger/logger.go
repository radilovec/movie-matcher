package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var Log *log.Logger

func InitLogger() {
	file, err := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}

	Log = log.New(file, "APP_LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func getCallerInfo() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "unknown file:unknown line"
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func LogInfo(message string) {
	Log.Println("INFO:", getCallerInfo(), message)
}

func LogError(message string) {
	Log.Println("ERROR:", getCallerInfo(), message)
}

func LogFatal(message string) {
	Log.Fatalln("FATAL:", getCallerInfo(), message)
}
