package services

import (
	"database/sql"

	"github.com/arshedke07/athletech/model"
)

func AddUserService(user *model.AppUser, profile *model.Preferences) (*model.AppUser, error) {
	insertstatement := "INSERT INTO app_user (firstname, lastname, password, emailid, mobile) VALUES ($1, $2, $3, $4, $5) RETURNING user_id"
	insertstatement2 := "INSERT INTO user_profile(user_id, age, height, weight, gender, experience, goal, current_body_type, gym_access, days_available, workout_time_preference, dietary_restrictions, injuries, medical_conditions) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)"

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var id int

	row := db.QueryRow(insertstatement, user.FirstName, user.LastName, user.Password, user.Email, user.Mobile)
	scanErr := row.Scan(&id)
	if scanErr != nil {
		return nil, scanErr
	}

	_, execErr := db.Exec(insertstatement2, id, profile.Age, profile.Height, profile.Weight, profile.Gender, profile.Experience, profile.Goal, profile.CurrentBodyType, profile.GymAccess, profile.DaysAvailable, profile.WorkoutTimePreference, profile.DietaryRestrictions, profile.Injuries, profile.MedicalConditions)
	if execErr != nil {
		return nil, execErr
	}

	newUser := model.AppUser{
		UserId:    id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Mobile:    user.Mobile,
	}

	return &newUser, nil
}

func UpdateUserService(userId int, trainerId int) error {
	updatestatement := "UPDATE app_user SET trainer = $1 WHERE user_id = $2"
	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return err
	}

	defer db.Close()

	_, execErr := db.Exec(updatestatement, trainerId, userId)
	if execErr != nil {
		return execErr
	}

	return nil
}

func GetUserById(userId int) (*model.AppUser, error) {
	selectstatement := "SELECT firstname, lastname, emailid, mobile, age, height, weight, gender, experience, goal, current_body_type, gym_access, days_available, workout_time_preference, dietary_restrictions, injuries, medical_conditions FROM app_user NATURAL JOIN user_profile WHERE user_id = $1"
	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	row := db.QueryRow(selectstatement, userId)

	user := model.AppUser{}

	scanErr := row.Scan(
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Mobile,
		&user.Preferences.Age,
		&user.Preferences.Height,
		&user.Preferences.Weight,
		&user.Preferences.Gender,
		&user.Preferences.Experience,
		&user.Preferences.Goal,
		&user.Preferences.CurrentBodyType,
		&user.Preferences.GymAccess,
		&user.Preferences.DaysAvailable,
		&user.Preferences.WorkoutTimePreference,
		&user.Preferences.DietaryRestrictions,
		&user.Preferences.Injuries,
		&user.Preferences.MedicalConditions,
	)
	if scanErr != nil {
		return nil, scanErr
	}

	return &user, nil
}
