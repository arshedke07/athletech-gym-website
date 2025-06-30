package services

import (
	"database/sql"

	"github.com/arshedke07/athletech/model"
)

func AddTrainerService(trainer *model.Trainer) (*model.Trainer, error) {
	insertstatement := "INSERT INTO app_trainer (firstname, lastname, age, gender, qualifications, email_id, password, mobile) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING trainer_id"
	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var id int

	row := db.QueryRow(insertstatement, trainer.FirstName, trainer.LastName, trainer.Age, trainer.Gender, trainer.Qualifications, trainer.Email, trainer.Password, trainer.Mobile)
	scanErr := row.Scan(&id)
	if scanErr != nil {
		return nil, scanErr
	}

	newUser := model.Trainer{
		Id:             id,
		FirstName:      trainer.FirstName,
		LastName:       trainer.LastName,
		Age:            trainer.Age,
		Gender:         trainer.Gender,
		Qualifications: trainer.Qualifications,
		Email:          trainer.Email,
		Password:       trainer.Password,
		Mobile:         trainer.Mobile,
	}

	return &newUser, nil
}

func GetPendingUsers(id int) (*[]model.AppUser, error) {
	selectstatement := "SELECT user_id, firstname, lastname, age, goal, experience FROM app_user NATURAL JOIN user_profile WHERE trainer = $1 "
	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	row, rowErr := db.Query(selectstatement, id)
	if rowErr != nil {
		return nil, rowErr
	}
	defer row.Close()

	users := []model.AppUser{}

	for row.Next() {
		user := model.AppUser{}
		scanErr := row.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Preferences.Age, &user.Preferences.Goal, &user.Preferences.Experience)
		if scanErr != nil {
			return nil, scanErr
		}
		users = append(users, user)
	}

	return &users, nil
}

func GetAllTrainers() (*[]model.Trainer, error) {
	selectstatement := "SELECT trainer_id, firstname, lastname, age, gender, qualifications, email_id, mobile FROM app_trainer"
	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	trainers := []model.Trainer{}

	rows, rowErr := db.Query(selectstatement)
	if rowErr != nil {
		return nil, rowErr
	}

	defer rows.Close()

	for rows.Next() {
		trainer := model.Trainer{}
		scanErr := rows.Scan(&trainer.Id, &trainer.FirstName, &trainer.LastName, &trainer.Age, &trainer.Gender, &trainer.Qualifications, &trainer.Email, &trainer.Mobile)
		if scanErr != nil {
			return nil, scanErr
		}

		trainers = append(trainers, trainer)
	}
	return &trainers, nil
}
