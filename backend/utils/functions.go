package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

/*
Handling logs
This functions are suitable for logging into the console using some
prefix to identify better the nature of the errors and logs.
*/
type logging struct {
	p *log.Logger
}

var Log *logging
var l logging

func InitLogs(prefix string) *log.Logger {
	l.p = log.New(os.Stdout, prefix, log.LstdFlags)
	l.Info("Initializing logs...")
	return l.p
}

func (lg *logging) Info(x ...interface{}) {
	s := "[INFO] "
	args := append([]interface{}{s}, x...)
	l.p.Println(args...)
}

func (lg *logging) Error(x ...interface{}) {
	s := "[ERROR]"
	args := append([]interface{}{s}, x...)
	l.p.Println(args...)
}

func (lg *logging) Debug(x ...interface{}) {
	s := "[DEBUG]"
	args := append([]interface{}{s}, x...)
	l.p.Println(args...)
}

/*
This function loads the .env file and returns the value of the env variable
*/
func DotEnvGet(key string) string {
	err := godotenv.Load()
	if err != nil {
		l.Error("Error loading .env file.", err)
	}

	return os.Getenv(key)
}


/* HTTP Responders */
type response struct{}
var Respond response

/*
This function responds the information successfully as JSON.
*/
func (r response) JSON(w http.ResponseWriter, code int, payload interface{}) {
	res, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}

/*
This function responds an error message to the frontend as JSON.
*/
func (r response) Error(w http.ResponseWriter, code int, message string) {
	r.JSON(w, code, map[string]string{"error": message})
}
