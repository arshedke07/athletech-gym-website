package services

import (
	"database/sql"

	"github.com/arshedke07/athletech/model"
)

var connectionstring string = "host=localhost port=5432 user=postgres password=1234 dbname=athletech sslmode=disable"

func LoginUserService(emailid string, password string) (*model.AppUser, error) {
	selectstatement := "SELECT user_id, firstname, lastname, password, emailid, mobile FROM app_user WHERE emailid = $1 AND password = $2"
	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var user model.AppUser

	row := db.QueryRow(selectstatement, emailid, password)
	err = row.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Password, &user.Email, &user.Mobile)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func LoginTrainerService(emailid string, password string) (*model.Trainer, error) {
	selectstatement := "SELECT trainer_id, firstname, lastname, age, gender, qualifications, email_id, password, mobile FROM app_trainer WHERE email_id = $1 AND password = $2"
	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var user model.Trainer

	row := db.QueryRow(selectstatement, emailid, password)
	err = row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Age, &user.Gender, &user.Qualifications, &user.Email, &user.Password, &user.Mobile)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
