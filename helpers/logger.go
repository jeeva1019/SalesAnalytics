package helpers

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type HelperStruct struct {
	Sid       string
	Reference string
	LogLevel  int
	LogFilter string
}

func (h *HelperStruct) Init() {
	h.Sid = strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}

func (h *HelperStruct) log(level string, msg ...any) {
	if h.Sid == "" {
		return
	}

	pc, file, line, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	fileShort := filepath.Join(filepath.Base(filepath.Dir(file)), filepath.Base(file))
	message := strings.Trim(fmt.Sprint(msg...), "[]")

	log.Printf("[%s] @@ %s @@ (%s) @@ %s @@ %s @@ ln %d @@ %s\n", level, h.Sid, h.Reference, fileShort, funcName, line, message)
}

func (h *HelperStruct) Info(msg ...any) {
	h.log("INFO", msg...)
}

func (h *HelperStruct) Warn(msg ...any) {
	h.log("WARN", msg...)
}

func (h *HelperStruct) Error(msg ...any) {
	h.log("ERROR", msg...)
}
