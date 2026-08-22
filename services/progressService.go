package services

import (
	"database/sql"

	"github.com/arshedke07/athletech/model"
)

func UserProgressService(progress model.Progress, userId int) error {
	upsertstatement := "INSERT INTO user_goals (user_id, current_weight, weight_goal, cardio_type, current_cardio, cardio_goal, lift_type, current_lift, lift_goal, current_body_fat, body_fat_goal, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) ON CONFLICT (user_id) DO UPDATE SET current_weight = COALESCE(EXCLUDED.current_weight, user_goals.current_weight), weight_goal = COALESCE(EXCLUDED.weight_goal, user_goals.weight_goal), cardio_type = COALESCE(NULLIF(EXCLUDED.cardio_type, ''), user_goals.cardio_type), current_cardio = COALESCE(EXCLUDED.current_cardio, user_goals.current_cardio), cardio_goal = COALESCE(EXCLUDED.cardio_goal, user_goals.cardio_goal), lift_type = COALESCE(NULLIF(EXCLUDED.lift_type, ''), user_goals.lift_type), current_lift = COALESCE(EXCLUDED.current_lift, user_goals.current_lift), lift_goal = COALESCE(EXCLUDED.lift_goal, user_goals.lift_goal), current_body_fat = COALESCE(EXCLUDED.current_body_fat, user_goals.current_body_fat), body_fat_goal = COALESCE(EXCLUDED.body_fat_goal, user_goals.body_fat_goal), updated_at = CURRENT_TIMESTAMP;"
	// in this query COALESCE returns the first non null value. if EXCLUDED.weight is non - null that is the new value i am trying to change the value is updated and if its null it keeps the old value

	db, dbErr := sql.Open("postgres", connectionstring)
	if dbErr != nil {
		return dbErr
	}

	defer db.Close()

	_, err := db.Exec(upsertstatement, userId, &progress.CurrentWeight, &progress.WeightGoal, &progress.CardioType, &progress.CurrentCardio, &progress.CardioGoal, &progress.LiftType, &progress.CurrentLift, &progress.LiftGoal, &progress.CurrentBodyFat, &progress.BodyFatGoal)
	if err != nil {
		return err
	}

	return nil
}

// func UpdateUser(preferences model.Preferences, userId int) error {
// 	// NULLIF($1, '') returns null if $1 is empty and returns $1 if its not empty and COALESCE returns the first non null value in its list of arguments
// 	updatestatement := "UPDATE user_profile SET age = COALESCE(NULLIF($1, ''), age), height = COALESCE(NULLIF($2, ''), height), goal = COALESCE(NULLIF($3, ''), goal), gym_access = COALESCE(NULLIF($4, ''), gym_access), days_available = COALESCE(NULLIF($5, ''), days_available), workout_time_preference = COALESCE(NULLIF($6, ''), workout_time_preference), dietary_restrictions = COALESCE(NULLIF($7, ''), dietary_restrictions), injuries = COALESCE(NULLIF($8, ''), injuries), medical_conditions = COALESCE(NULLIF($9, ''), medical_conditions) WHERE user_id = $10;"
// 	db, err := sql.Open("postgres", connectionstring)
// 	if err != nil {
// 		return err
// 	}

// 	defer db.Close()

// 	_, execErr := db.Exec(updatestatement, preferences.Age, preferences.Height, preferences.Goal, preferences.GymAccess, preferences.DaysAvailable, preferences.WorkoutTimePreference, preferences.DietaryRestrictions, preferences.Injuries, preferences.MedicalConditions, userId)
// 	if execErr != nil {
// 		return execErr
// 	}

// 	return nil
// }

func GetUserProgress(userId int) (*model.Progress, *model.Progress, error) {
	selectstatement := "SELECT current_weight, weight_goal, cardio_type, current_cardio, cardio_goal, lift_type, current_lift, lift_goal, current_body_fat, body_fat_goal FROM user_goals WHERE user_id = $1"
	selectstatement2 := "SELECT current_weight, cardio_type, current_cardio, lift_type, current_lift, current_body_fat FROM user_profile_history WHERE user_id = $1 ORDER BY updated_at ASC LIMIT 1"
	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return nil, nil, err
	}

	defer db.Close()

	progress := model.Progress{}

	row := db.QueryRow(selectstatement, userId)
	scanErr := row.Scan(&progress.CurrentWeight, &progress.WeightGoal, &progress.CardioType, &progress.CurrentCardio, &progress.CardioGoal, &progress.LiftType, &progress.CurrentLift, &progress.LiftGoal, &progress.CurrentBodyFat, &progress.BodyFatGoal)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return &progress, &progress, nil
		}
		return nil, nil, scanErr
	}

	start := model.Progress{}

	row2 := db.QueryRow(selectstatement2, userId)
	scanErr2 := row2.Scan(&start.CurrentWeight, &start.CardioType, &start.CurrentCardio, &start.LiftType, &start.CurrentLift, &start.CurrentBodyFat)
	if scanErr2 != nil {
		if scanErr2 == sql.ErrNoRows {
			return &progress, &progress, nil
		}
		return nil, nil, scanErr
	}

	return &progress, &start, nil
}
