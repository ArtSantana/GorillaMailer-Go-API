package main

import (
	"fmt"
	"gorillaMailer/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":3000"
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hello World")
	})
	router.HandleFunc("/subscriber", controllers.GetSubscribers).Methods("GET")
	router.HandleFunc("/subscriber", controllers.CreateSubscriber).Methods("POST")
	router.HandleFunc("/subscriber/{id}", controllers.GetUniqueSubscriber).Methods("GET")
	// router.HandleFunc("/subscriber/{id}", controllers.DeleteSubscriber).Methods("DELETE")
	// router.HandleFunc("/subscriber/{id}", controllers.EditSubscriber).Methods("PATCH")

	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))

}
