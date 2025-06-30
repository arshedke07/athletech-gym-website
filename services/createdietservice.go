package services

import (
	"database/sql"

	"github.com/arshedke07/athletech/model"
)

func GetDayByName(name string) (int, error) {
	selectstatement := "SELECT id FROM days WHERE name = $1"
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

func GetMealType(mealType string) (int, error) {
	selectstatement := "SELECT id from meal_type WHERE meal_type = $1"
	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return -1, err
	}

	defer db.Close()
	var id int

	row := db.QueryRow(selectstatement, mealType)
	scanErr := row.Scan(&id)
	if scanErr != nil {
		return -1, scanErr
	}

	return id, nil
}

func FindMealByName(name string) (int, error) {
	selectstatement := "SELECT meal_id FROM meal WHERE name = $1"
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

func InsertIntoMeal(meal string) error {
	insertstatement := "INSERT INTO meal(name) VALUES ($1) ON CONFLICT(name) DO NOTHING"
	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return err
	}

	_, rowErr := db.Exec(insertstatement, meal)
	if rowErr != nil {
		return rowErr
	}

	return nil
}

func CreateDietService(dietDay model.DietDay, userId int) error {
	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		return err
	}

	defer db.Close()

	insertstatement1 := "INSERT INTO nutrients(protein, fats, carbs, calories) VALUES ($1, $2, $3, $4) returning id"
	insertstatement2 := "INSERT INTO diet(user_id, day_id, meal_type, meal_id, nutrient_id) VALUES($1, $2, $3, $4, $5)"

	var nutrientId int

	row1 := db.QueryRow(insertstatement1, dietDay.Nutrients.Protein, dietDay.Nutrients.Fats, dietDay.Nutrients.Carbs, dietDay.Nutrients.Calories)
	rowErr1 := row1.Scan(&nutrientId)
	if rowErr1 != nil {
		return rowErr1
	}

	for _, diet := range dietDay.Diet {
		dietErr := InsertIntoMeal(diet.Meal) // insert meal into meal table
		if dietErr != nil {
			return dietErr
		}

		mealId, findErr := FindMealByName(diet.Meal) // get meal id from meal table
		if findErr != nil {
			return findErr
		}

		dayId, dayErr := GetDayByName(dietDay.Day)
		if dayErr != nil {
			return dayErr
		}

		mealTypeId, mealErr := GetMealType(diet.MealType)
		if mealErr != nil {
			return mealErr
		}

		_, execErr := db.Exec(insertstatement2, userId, dayId, mealTypeId, mealId, nutrientId)
		if execErr != nil {
			return execErr
		}
	}

	return nil
}
