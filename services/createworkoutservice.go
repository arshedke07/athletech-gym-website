package services

import (
	"database/sql"
)

func FindExerciseByName(name string) (int, error) {
	selectstatement := "SELECT id from exercises WHERE name = $1"
	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return -1, err
	}

	defer db.Close()

	var id int
	row := db.QueryRow(selectstatement, name)
	scanErr := row.Scan(&id)
	if scanErr != nil {
		return -1, scanErr
	}

	return id, nil
}

func CreateWorkoutService(userId int, exercises [][]string, desc [][]string, muscles []string) error {
	insertstatement := "INSERT INTO workout (user_id, day_id, exercise_id, description) VALUES($1, $2, $3, $4)"
	insertstatement2 := "INSERT INTO exercises(name, muscle_group_id) VALUES($1, $2) ON CONFLICT (name) DO NOTHING"
	insertstatement3 := "WITH inserted AS (INSERT INTO muscle_group(muscle_group) VALUES ($1) ON CONFLICT (muscle_group) DO NOTHING RETURNING id) SELECT id FROM inserted UNION ALL SELECT id FROM muscle_group WHERE muscle_group = $1 LIMIT 1;"
	// this query returns muscle_group_id if a new row is inserted and returns just the id if the row already exists in the table

	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return err
	}

	defer db.Close()

	for i := 0; i < 7; i++ { // loop runs till 7 as their are 7 days in the week
		var muscleId int

		row := db.QueryRow(insertstatement3, muscles[i])
		scanErr := row.Scan(&muscleId)
		if scanErr != nil {
			return scanErr
		}

		for j := range exercises[i] {
			_, execErr := db.Exec(insertstatement2, exercises[i][j], muscleId)
			if execErr != nil {
				return execErr
			}

			id, err := FindExerciseByName(exercises[i][j])
			if err != nil {
				return err
			}

			_, insertErr := db.Exec(insertstatement, userId, i+1, id, desc[i][j])
			if insertErr != nil {
				return insertErr
			}
		}
	}

	// updatestatement := "UPDATE app_user SET plan_status = 'not requested' WHERE user_id = $1"

	// _, newErr := db.Exec(updatestatement, userId)
	// if newErr != nil {
	// 	return newErr
	// }

	return nil
}
