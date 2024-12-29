package main

import (
	"fmt"
	"io"
	"net/http"
)

const (
	InfoMessage    = "INFO"
	WarningMessage = "WARNING"
	ErrorMessage   = "ERROR"
)

func main() {
	ParseFlags()
	fmt.Println(GetPort())

	http.HandleFunc("/Info", PrintInfo)
	http.HandleFunc("/Warn", PrintWarning)
	http.HandleFunc("/Err", PrintError)
	http.ListenAndServe(":6969", nil)
}

func printMessage(errorType string, message string) {
	fmt.Printf("[")
	switch errorType {
	case InfoMessage:
		fmt.Printf(string("\033[32m"))
	case WarningMessage:
		fmt.Printf(string("\033[33m"))
	case ErrorMessage:
		fmt.Printf(string("\033[31m"))
	}
	fmt.Printf(errorType)
	fmt.Printf(string("\033[0m"))
	fmt.Printf("] %s \n", message)
}

func GenericController(writer http.ResponseWriter, request *http.Request, MsgType string) {
	writer.Header().Set("Content-Type", "application/json")
	switch request.Method {
	case "POST":
		bodyBytes, err := io.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
		bodyString := string(bodyBytes)
		printMessage(MsgType, bodyString)
		writer.WriteHeader(http.StatusOK)
	default:
		printMessage(ErrorMessage, fmt.Sprintf("Wrong request type: use POST-request instead of %s", request.Method))
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func PrintInfo(writer http.ResponseWriter, request *http.Request) {
	GenericController(writer, request, InfoMessage)
}

func PrintWarning(writer http.ResponseWriter, request *http.Request) {
	GenericController(writer, request, WarningMessage)
}

func PrintError(writer http.ResponseWriter, request *http.Request) {
	GenericController(writer, request, ErrorMessage)
}
