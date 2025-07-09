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
	Age                   int
	Height                int
	Weight                int
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
}
