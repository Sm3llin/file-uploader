package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sm3llin/file-uploader/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", routes.IndexHandler)
	router.HandleFunc("/upload", routes.UploadHandler)

	loggingRouter := handlers.LoggingHandler(os.Stdout, router)

	fmt.Println("listening on 0.0.0.0:8000 . . . ")
	log.Fatal(http.ListenAndServe(":8000", loggingRouter))
}
