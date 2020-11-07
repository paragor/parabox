package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var validationsRegexps = []*regexp.Regexp{
	regexp.MustCompile("^(\\d{10})$"),
	regexp.MustCompile("^(\\d{4}) (\\d{6})$"),
	regexp.MustCompile("^\\((\\d{4})\\) (\\d{6})$"),
	regexp.MustCompile("^\\((\\d{4})\\)(\\d{6})$"),
	regexp.MustCompile("^(\\d{4})-(\\d{6})$"),
	regexp.MustCompile("^PC-(\\d{4})-(\\d{6})$"),
}

func main() {
	shutdownChan := make(chan bool)
	go func() {
		<-shutdownChan
		os.Exit(0)
	}()

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/validatePassportCode", validatePassportCodeHandler)
	http.HandleFunc("/shutdown", createShutdownHandler(shutdownChan))
	log.Fatal(http.ListenAndServe(":7777", nil))
}
func createShutdownHandler(shutdownChan chan bool) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		writer.Write(nil)
		fmt.Println("shutdown cmd")
		shutdownChan <- true
	}
}

func pingHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
}
func validatePassportCodeHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		writer.WriteHeader(400)
		fmt.Printf("[ERR] error parse form:%s\n", err.Error())
		return
	}

	passportCode := request.Form.Get("passport_code")
	if passportCode == "" {
		writer.WriteHeader(400)
		fmt.Printf("[ERR] invalid passport code\n")
		return
	}

	result, err := ValidatePassportCode(passportCode)
	if err != nil {
		response := struct {
			Status bool `json:"status"`
		}{
			Status: false,
		}
		bytes, err := json.Marshal(response)
		if err != nil {
			writer.WriteHeader(400)
			fmt.Printf("[ERR] cant marshal bad response\n")
			return
		}
		writer.WriteHeader(200)
		writer.Header().Set("content-type", "application/json")
		writer.Write(bytes)
		fmt.Printf("[OK] fail validate passport code\n")
		return
	}

	response := struct {
		Status     bool   `json:"status"`
		Normalized string `json:"normalized"`
	}{
		Status:     true,
		Normalized: result,
	}
	bytes, err := json.Marshal(response)
	if err != nil {
		writer.WriteHeader(400)
		fmt.Printf("[ERR] cant marshal good response\n")
		return
	}
	writer.Header().Set("content-type", "application/json")
	writer.WriteHeader(200)
	writer.Write(bytes)
	fmt.Printf("[OK] valide passport code!\n")
	return
}

func ValidatePassportCode(code string) (string, error) {
	code = strings.TrimSpace(code)
	result := ""
	for _, r := range validationsRegexps {
		if r.MatchString(code) {
			submatchs := r.FindStringSubmatch(code)
			result = strings.Join(submatchs[1:], "")
			result = result[:4] + "-" + result[4:]
		}
	}

	if result == "" {
		return result, errors.New("nope")
	}
	return result, nil
}
