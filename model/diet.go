package model

type Nutrients struct {
	Protein  int
	Fats     int
	Carbs    int
	Calories int
}

type Diet struct {
	Meal     string
	MealType string
}

type DietDay struct {
	Day       string
	Nutrients Nutrients
	Diet      []Diet
}
