package services

import (
	"database/sql"

	"github.com/arshedke07/athletech/model"
)

func GetWokoutService(userId int) (*[]*model.Workout, error) {
	selectstatement := "SELECT d.day AS day_name, e.name AS exercise_name, mg.muscle_group AS muscle_group_name, w.description FROM workout w JOIN days d ON w.day_id = d.id JOIN exercises e ON w.exercise_id = e.id JOIN muscle_group mg ON e.muscle_group_id = mg.id WHERE w.user_id = $1 ORDER BY d.id"
	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, rowErr := db.Query(selectstatement, userId)
	if rowErr != nil {
		return nil, rowErr
	}

	defer rows.Close()

	workoutMap := make(map[string]*model.Workout) // using map to order exercises by day and we have mapped to pointer so that we are able to modify the orginal value stored in the map
	workouts := []*model.Workout{}                // use slice to maintain order of days as map does not maintain any order

	for rows.Next() {
		ex := model.Exercise{}
		day := model.Day{}

		scanErr := rows.Scan(&day.Name, &ex.Exercise, &day.MuscleGroups, &ex.Description)
		if scanErr != nil {
			return nil, scanErr
		}

		w, exists := workoutMap[day.Name]
		if !exists { // this day does not already exist in map so make new entry
			w = &model.Workout{
				Day:       day,
				Exercises: []model.Exercise{ex},
			}

			workoutMap[day.Name] = w
			workouts = append(workouts, w)
		} else {
			w.Exercises = append(w.Exercises, ex)
		}
	}

	return &workouts, nil
}

func GetDietService(userId int) (*[]*model.DietDay, error) {
	selectstatement := "SELECT d.day, mt.meal_type, m.name, n.protein, n.fats, n.carbs, n.calories FROM diet di JOIN days d ON di.day_id = d.id JOIN meal_type mt ON mt.id = di.meal_type JOIN meal m ON m.meal_id = di.meal_id JOIN nutrients n ON n.id = di.nutrient_id WHERE di.user_id = $1"
	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, rowErr := db.Query(selectstatement, userId)
	if rowErr != nil {
		return nil, err
	}

	dietMap := make(map[string]*model.DietDay)
	diets := []*model.DietDay{}

	for rows.Next() {
		nutrients := model.Nutrients{}
		diet := model.Diet{}
		var dayName string

		scanErr := rows.Scan(&dayName, &diet.MealType, &diet.Meal, &nutrients.Protein, &nutrients.Fats, &nutrients.Carbs, &nutrients.Calories)
		if scanErr != nil {
			return nil, scanErr
		}

		d, exists := dietMap[dayName]
		if !exists {
			d = &model.DietDay{
				Day:       dayName,
				Nutrients: nutrients,
				Diet:      []model.Diet{diet},
			}
			dietMap[dayName] = d
			diets = append(diets, dietMap[dayName])

		} else {
			d.Diet = append(d.Diet, diet)
		}
	}

	return &diets, nil
}
