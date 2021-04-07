package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

//GetJsonResponse return a response in json format for the request
func GetJsonResponse(writer http.ResponseWriter, statusCode int, data interface{}) {

	writer.WriteHeader(statusCode)

	if error := json.NewEncoder(writer).Encode(data); error != nil {
		log.Fatal(error)
	}
}

//Return an error
func GetError(writer http.ResponseWriter, statusCode int, e error) {

	GetJsonResponse(writer, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: e.Error(),
	})
}
