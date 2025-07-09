package model

type Progress struct {
	UserId         int
	CurrentWeight  *string
	WeightGoal     *string
	CardioType     *string
	CurrentCardio  *string
	CardioGoal     *string
	LiftType       *string
	CurrentLift    *string
	LiftGoal       *string
	CurrentBodyFat *string
	BodyFatGoal    *string
}
