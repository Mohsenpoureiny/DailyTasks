// Package model provides ...
package model

import (
	"github.com/taigacute/DailyTasks/database"
	"golang.org/x/crypto/bcrypt"
)

//User struct
type User struct {
	id       int
	Username string
	Password string
}

//RegisterUser add  user
func (user *User) RegisterUser(uname string, pwd string, email string) error {
	sql := "insert into user(username,password,email)values(?,?,?)"
	err := database.TaskExec(sql, uname, pwd, email)
	return err
}

//IsExist will return true when user isExist in databse
func (user *User) UserIsExist(uname string) bool {
	var username string
	sql := "select username from user where username = ?"
	rows := database.TaskQueryRows(sql, uname)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&username)
		if err != nil {
			return false
		}
		return true
	}
	return false
}

//ValidUser will return true when user password valided
func (user *User) ValidUser(uname, pwd string) bool {
	var passwordformdb string
	sql := "select password from user where username = ?"
	rows := database.TaskQueryRows(sql, uname)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&passwordformdb)
		if err != nil {
			return false
		}
	}
	err := bcrypt.CompareHashAndPassword([]byte(passwordformdb), []byte(pwd))
	if err != nil {
		return false
	}
	return true
}