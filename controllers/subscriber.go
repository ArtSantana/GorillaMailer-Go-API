package controllers

import (
	"encoding/json"
	"fmt"
	"gorillaMailer/models"
	"gorillaMailer/repository"
	"gorillaMailer/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetSubscribers controller
func GetSubscribers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	subscribers := repository.GetSubscribers()

	result, err := json.Marshal(subscribers)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling the sensors"}`))
	}

	res.WriteHeader(http.StatusOK)
	res.Write(result)
	return
}

// CreateSubscriber controller
func CreateSubscriber(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var subscriber models.Subscriber

	err := decoder.Decode(&subscriber)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	created, result := repository.CreateSubscriber(&subscriber)

	if !created {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	location := fmt.Sprintf("http://localhost:300/subscriber/%d", result.ID)

	res.Header().Set("location", location)
	res.WriteHeader(http.StatusCreated)

}

// GetUniqueSubscriber controller
func GetUniqueSubscriber(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id := params["id"]

	idParsed, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	if idParsed < 1 {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	subscriber := repository.GetUniqueSubscriber(idParsed)

	exist, result := services.ValidationExist(subscriber)

	if !exist {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(result)
	return,
}
