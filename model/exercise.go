package model

type Exercise struct {
	Exercise    string
	Description string
}

type Day struct {
	Name         string
	MuscleGroups string
}

type Workout struct {
	Day       Day
	Exercises []Exercise
}
