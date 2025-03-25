package logger

import (
	"fmt"
	"log"
	"os"
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

func LogInfo(message string) {
	Log.Println("INFO: " + message)
}

func LogError(message string) {
	Log.Println("ERROR: " + message)
}

func LogFatal(message string) {
	Log.Fatalln("FATAL: " + message)
}
