package controllers

import (
	"SalesAnalytics/helpers"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

type RspStruct struct {
	ErrMsg string `json:"msg,omitempty"`
	Status string `json:"status,omitempty"`
	Result any    `json:"result,omitempty"`
}

func ResponseWriter(debug *helpers.HelperStruct, resp http.ResponseWriter, result any) {
	debug.Info("ResponseWriter(+)")

	respStr, err := json.Marshal(result)
	if err != nil {
		debug.Error("error occured at marshal", err)
		return
	}

	fmt.Fprintln(resp, string(respStr))
	debug.Info("ResponseWriter(-)")
}

func ResponseConstructor(status, msg string, val any) string {
	var finalRes RspStruct

	finalRes.Status = status
	finalRes.ErrMsg = msg
	finalRes.Result = val

	bodyStr, err := json.Marshal(finalRes)
	if err != nil {
		return ""
	}

	return string(bodyStr)
}

func CSVFileReader(debug *helpers.HelperStruct, file *os.File, Struct interface{}) error {
	debug.Info("CSVFileReader(+)")
	if err := gocsv.UnmarshalFile(file, Struct); err != nil {
		debug.Error("CRCFR:001", err)
		return err
	}
	debug.Info("CSVFileReader(-)")
	return nil
}

func Validator(fields map[string]string) error {
	for name, val := range fields {
		if strings.TrimSpace(val) == "" {
			return fmt.Errorf("field '%s' is required", name)
		}
	}
	return nil
}
