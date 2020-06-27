package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("hello world", zap.Time("now", time.Now()))
		fmt.Fprintf(w, "hello world")
	})
	log.Println("listening on port *:8080. press ctrl + c to cancel")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
