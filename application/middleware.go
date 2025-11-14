package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type middleware struct {
	logger  *os.File
	handler http.Handler
}

func (mdlwr middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format("2006-01-02 15:04:05")
	mdlwr.logger.WriteString(fmt.Sprintf("%v : %s - %s\n", now, r.Method, r.URL))
	mdlwr.handler.ServeHTTP(w, r)
}

func NewMiddleWare(mx *http.ServeMux) *middleware {

	file, err := os.Create("logger.log")
	if err != nil {
		log.Println("error forming a connection with log file: ", err.Error())
	}

	return &middleware{logger: file,
		handler: mx}
}
