package database

type User struct {
	ID       string
	UserName string
	Password string
	Name     string
	IsRefree bool //to differentiate the flows
}

//we can use a user factory here
