package services

import (
	"database/sql"

	"github.com/arshedke07/athletech/model"
)

func AddTrainerService(profile *model.TrainerProfile) (*model.AppUser, error) {
	insertstatement1 := "INSERT INTO app_user (firstname, lastname, emailid, password, mobile, role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id"
	insertstatement2 := "INSERT INTO trainer_profile (user_id, age, gender, specialization, experience, languages, city, state, social_media, description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var id int

	row := db.QueryRow(insertstatement1, profile.User.FirstName, profile.User.LastName, profile.User.Email, profile.User.Password, profile.User.Mobile, profile.User.Role)
	scanErr := row.Scan(&id)
	if scanErr != nil {
		return nil, scanErr
	}

	_, execErr := db.Exec(insertstatement2, id, profile.Age, profile.Gender, profile.Specialization, profile.Experience, profile.Languages, profile.City, profile.State, profile.SocialMedia, profile.Description)
	if execErr != nil {
		return nil, execErr
	}

	newUser := model.AppUser{
		UserId:    id,
		FirstName: profile.User.FirstName,
		LastName:  profile.User.LastName,
		Email:     profile.User.Email,
		Password:  profile.User.Password,
		Mobile:    profile.User.Mobile,
	}

	return &newUser, nil
}

func GetPendingUsers(id int) (*[]model.UserProfile, error) {
	selectstatement := "SELECT user_id, firstname, lastname, age, goal FROM app_user NATURAL JOIN user_profile up WHERE up.trainer_id = $1 "
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

	users := []model.UserProfile{}

	for row.Next() {
		user := model.UserProfile{}
		scanErr := row.Scan(&user.User.UserId, &user.User.FirstName, &user.User.LastName, &user.Age, &user.Goal)
		if scanErr != nil {
			return nil, scanErr
		}
		users = append(users, user)
	}

	return &users, nil
}

func GetAllTrainers() (*[]model.TrainerProfile, error) {
	selectstatement := "SELECT user_id, firstname, lastname, emailid, mobile, age, gender, specialization, experience, languages, city, state, social_media, description FROM app_user au NATURAL JOIN trainer_profile WHERE au.role = 'trainer' "
	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	trainers := []model.TrainerProfile{}

	rows, rowErr := db.Query(selectstatement)
	if rowErr != nil {
		return nil, rowErr
	}

	defer rows.Close()

	for rows.Next() {
		trainer := model.TrainerProfile{}
		scanErr := rows.Scan(&trainer.User.UserId, &trainer.User.FirstName, &trainer.User.LastName, &trainer.User.Email, &trainer.User.Mobile, &trainer.Age, &trainer.Gender, &trainer.Specialization, &trainer.Experience, &trainer.Languages, &trainer.City, &trainer.State, &trainer.SocialMedia, &trainer.Description)
		if scanErr != nil {
			return nil, scanErr
		}

		trainers = append(trainers, trainer)
	}
	return &trainers, nil
}
