package postgresdb

import (
	"context"
	"crud/internal/service"
)

func (db *DB) CreateUser(name string, surname string, id uint) error {
	query := `INSERT INTO users (id, name, surname) VALUES ($1, $2, $3)`
	_, err := db.db.Exec(context.Background(), query, id, name, surname)
	return err
}

func (db *DB) GetAllUsers() ([]service.User, error) {
	rows, err := db.db.Query(context.Background(), `SELECT id, name, surname FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []service.User
	for rows.Next() {
		var user service.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Surname); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (db *DB) GetUserById(id string) ([]service.User, error) {
	query := `SELECT id, name, surname FROM users WHERE id = $1`
	row := db.db.QueryRow(context.Background(), query, id)

	var user service.User
	err := row.Scan(&user.Id, &user.Name, &user.Surname)
	if err != nil {
		return nil, err
	}

	return []service.User{user}, nil
}

func (db *DB) UpdateUser(id string, user service.User) ([]service.User, error) {
	query := `UPDATE users SET name=$1, surname=$2 WHERE id=$3`
	_, err := db.db.Exec(context.Background(), query, user.Name, user.Surname, id)
	if err != nil {
		return nil, err
	}

	selectQuery := `SELECT id, name, surname FROM users WHERE id = $1`
	row := db.db.QueryRow(context.Background(), selectQuery, id)

	var updatedUser service.User
	err = row.Scan(&updatedUser.Id, &updatedUser.Name, &updatedUser.Surname)
	if err != nil {
		return nil, err
	}

	return []service.User{updatedUser}, nil
}

func (db *DB) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := db.db.Exec(context.Background(), query, id)
	return err
}
