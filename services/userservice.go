package services

import (
	"database/sql"

	"github.com/arshedke07/athletech/model"
)

func AddUserService(profile *model.UserProfile) (*model.AppUser, error) {
	insertstatement := "INSERT INTO app_user (firstname, lastname, password, emailid, mobile, role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id"
	insertstatement2 := "INSERT INTO user_profile(user_id, age, height, weight, gender, experience, goal, current_body_type, gym_access, days_available, workout_time_preference, dietary_restrictions, injuries, medical_conditions, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var id int

	row := db.QueryRow(insertstatement, profile.User.FirstName, profile.User.LastName, profile.User.Password, profile.User.Email, profile.User.Mobile, profile.User.Role)
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
		FirstName: profile.User.FirstName,
		LastName:  profile.User.LastName,
		Email:     profile.User.Email,
		Mobile:    profile.User.Mobile,
		Role:      profile.User.Role,
	}

	return &newUser, nil
}

func UpdateUserService(userId int, trainerId int) error {
	updatestatement := "UPDATE user_profile SET trainer_id = $1 WHERE user_id = $2"
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

func GetUserById(userId int) (*model.UserProfile, error) {
	selectstatement := "SELECT firstname, lastname, emailid, mobile, age, height, weight, gender, experience, goal, current_body_type, gym_access, days_available, workout_time_preference, dietary_restrictions, injuries, medical_conditions FROM app_user NATURAL JOIN user_profile WHERE user_id = $1"
	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	row := db.QueryRow(selectstatement, userId)

	user := model.UserProfile{}

	scanErr := row.Scan(
		&user.User.FirstName,
		&user.User.LastName,
		&user.User.Email,
		&user.User.Mobile,
		&user.Age,
		&user.Height,
		&user.Weight,
		&user.Gender,
		&user.Experience,
		&user.Goal,
		&user.CurrentBodyType,
		&user.GymAccess,
		&user.DaysAvailable,
		&user.WorkoutTimePreference,
		&user.DietaryRestrictions,
		&user.Injuries,
		&user.MedicalConditions,
	)
	if scanErr != nil {
		return nil, scanErr
	}

	return &user, nil
}
