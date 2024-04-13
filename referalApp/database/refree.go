package database

type Refree struct {
	ID            string //passing the authentication details here
	RefreeCompany string
	RefreeAuth    *RefreeAuth
}

type RefreeAuth struct {
	ID       string
	Email    string
	Password string
	Name     string
}
