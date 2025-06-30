package model

type AppUser struct {
	UserId      int
	FirstName   string
	LastName    string
	Password    string
	Email       string
	Mobile      string
	Preferences Preferences
}

type Preferences struct {
	Age                   string
	Height                string
	Weight                string
	Gender                string
	Experience            string
	Goal                  string
	CurrentBodyType       string
	GymAccess             string
	DaysAvailable         string
	WorkoutTimePreference string
	DietaryRestrictions   string
	Injuries              string
	MedicalConditions     string
}
