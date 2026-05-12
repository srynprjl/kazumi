package misc

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

func Log(msg string, typ string) {
	datetime := time.Now().Format("2006-01-02 15:04:05")
	date := time.Now().Format("2006-01-02")
	a := "NORMAL"
	switch typ {
	case "e":
		a = "ERROR"
	case "f":
		a = "FATAL"
	}
	log_msg := fmt.Sprintf("[%s] [%s] > %s\n", a, datetime, msg)
	p := GetLogDir()
	current_log := path.Join(p, strings.Join([]string{date, "txt"}, "."))

	f, err := os.OpenFile(current_log, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if _, err := f.WriteString(log_msg); err != nil {
		log.Fatal(err)
	}
}
