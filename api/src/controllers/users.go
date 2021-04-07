package controllers

import (
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//docker exec -it mysql-latest mysql -uroot -pstrongpassword
//Create a user
func CreateUser(w http.ResponseWriter, response *http.Request) {

	request, error := ioutil.ReadAll(response.Body)

	if error != nil {
		responses.GetError(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user model.User
	if error = json.Unmarshal(request, &user); error != nil {
		responses.GetError(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.GetError(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	user.ID, error = repository.Create(user)
	if error != nil {
		responses.GetError(w, http.StatusInternalServerError, error)
		return
	}
	responses.GetJsonResponse(w, http.StatusCreated, user)
}

//List users
func ListUsers(writer http.ResponseWriter, response *http.Request) {
	writer.Write([]byte("Listando Usuarios"))
}

//Find user by id
func FindUserById(writer http.ResponseWriter, response *http.Request) {
	writer.Write([]byte("Encontrando Usuario"))
}

//Delete user
func DeleteUserById(writer http.ResponseWriter, response *http.Request) {
	writer.Write([]byte("Deletando Usuario"))
}

//Update user
func UpdateUserById(writer http.ResponseWriter, response *http.Request) {
	writer.Write([]byte("atualizando Usuario"))
}
