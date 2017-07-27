package models

import (
	"fmt"
	"../consts"
	"strconv"
)

// Provide structure of main app model
type User struct {
	Id int `json:"id"`
	Balance int `json:"balance"`
}

type Users []User

func GetUser(id int) (*User, error) {
	sql := fmt.Sprintf("SELECT * FROM users WHERE users.id = %d", id)
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.Id, &user.Balance)

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, fmt.Errorf(consts.UserNotFound)
}

func DeleteUser(id int) (error) {
	sql := fmt.Sprintf("DELETE FROM users WHERE users.id = %d", id)
	return simpleQuery(sql)
}

func DeleteUsers() (error) {
	sql := "DELETE FROM users"
	return simpleQuery(sql)
}

func CreateUser(balance int) (error) {
	sql := fmt.Sprintf("INSERT INTO users (balance) values (%d)", balance)
	return simpleQuery(sql)
}

func UpdateUser(id, balance int) (error) {
	sql := fmt.Sprintf("UPDATE users SET balance = %d WHERE id = %d", balance, id)
	return simpleQuery(sql)
}

func SetUserPrise(id int, backers []int, prise int) {
	sql := `UPDATE users SET balance = balance - ` + strconv.Itoa(prise) + ` WHERE id IN
		(` + arrayToString(append(backers, id), ", ") + `)`
	db.Query(sql)
}

