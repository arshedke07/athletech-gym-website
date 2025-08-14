package model

type AppUser struct {
	UserId    int
	FirstName string
	LastName  string
	Password  string
	Email     string
	Mobile    string
	Role      string
}

type UserProfile struct {
	User                  AppUser
	Age                   int
	Height                int
	Weight                float32
	Gender                string
	Experience            string
	Goal                  string
	CurrentBodyType       string
	GymAccess             string
	DaysAvailable         int
	WorkoutTimePreference string
	DietaryRestrictions   string
	Injuries              string
	MedicalConditions     string
	WorkoutCreated        bool // these are flags to check if the user already has a plan created by their trainer or not
	DietCreated           bool
}

type TrainerProfile struct {
	User           AppUser
	Age            int
	Gender         string
	Specialization string
	Experience     string
	Languages      string
	City           string
	State          string
	SocialMedia    string
	Description    string
}
