package repository

import (
	"api/src/model"
	"database/sql"
)

type users struct {
	db *sql.DB
}

// Cria um repositorios de usuarios
func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

//Insere um usuario na base de dados.
func (userRepository *users) Create(user model.User) (uint64, error) {
	statement, error := userRepository.db.Prepare(
		"insert into user (nome, nick, email, senha) values(?, ?, ?, ?)",
	)

	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if error != nil {
		return 0, error
	}

	lastID, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastID), nil

}

// func NewUserRepository(db *sql.DB) *users {
// 	return &users{db}
// }
