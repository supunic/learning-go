package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

func main() {
	// doc
	// https://pkg.go.dev/log

	LoggingSettings("test.log")
	_, err := os.Open("fafafa")
	if err != nil {
		log.Fatalln("Exit", err) // プログラム終了
	}

	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	log.Fatalf("%T %v", "test", "test")
	log.Fatalln("error!!") // プログラム終了

	fmt.Println("ok!")
}
