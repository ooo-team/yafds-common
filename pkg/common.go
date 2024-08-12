package common

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"

	"github.com/joho/godotenv"
)

func ReadHeaderParam(w http.ResponseWriter, r *http.Request, paramName string, required bool) string {
	paramStr := r.URL.Query().Get(paramName)
	if required && paramStr == "" {
		msg := "Missing required request param" + paramName
		http.Error(w, msg, http.StatusBadRequest)
		return ""
	}
	return paramStr
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func InitEnv() {
	homedir := os.Getenv("HOME")
	if err := godotenv.Load(homedir + "/.config/go/env/.env"); err != nil {
		log.Panic("No .env file found", err.Error())
	}
}

func Valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func LoadEnvVar(varName string) (string, error) {
	var_, exists := os.LookupEnv(varName)

	if !exists {
		InitEnv()
		infMsg := fmt.Sprintf("Env variable %s is not set, calling InitEnv", varName)
		log.Println(infMsg)
		var_, exists = os.LookupEnv(varName)
	}

	errMsg := fmt.Sprintf("Env variable %s is not set", varName)
	if !exists {
		return "", &NotFoundError{Message: errMsg}
	}
	return var_, nil
}
