package repo

import (
	"fmt"

	"github.com/sahilwahi7/referalApp/models"
)

type UpdatedUser struct {
	UserName       string `json:"userName"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	LinkedProfile  string `json:"linkedinProfile`
	CurrentCompany string `json:"company"`
	Name           string `json:"name"`
	IsRefree       bool   `json:"isRefree`
}

type Data struct {
	JobsTitle       string `json:"jobTitle"`
	PostedBy        string `json:"postedBy"`
	TotalApplicants int64  `json:"totalApplicants"`
	Compnay         string `json:"company"`
}

type Authentication struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type User struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsRefree bool   `json:"isRefree"`
}

type Concreterepo struct {
}

type Repo interface {
	// AuthenticateCandidate(email string, password string) *database.CandidateAuth
	// AuthenticateRefree(email string, password string) *database.RefreeAuth
	// ApplyJob(joblink string, status bool) error
	FindOpenJobs(company string) []Data
	UpdateUserProfile(username string, linkedProfile string, currentCompany string, description string, title string) bool
	Login(username string, password string) bool
	Signup(name string, username string, password string, id string, isrefree bool) bool
	//ViewAllUsers() []Authentication
}

func (c *Concreterepo) UpdateUserProfile(username string, linkedProfile string, currentCompany string, description string, title string) bool {
	user := &models.ConcreteUser{}
	resp := user.CheckDetails(username, linkedProfile, currentCompany, description, title)
	if resp == true {
		fmt.Printf("User updated successfully....")
		return true
	} else {
		fmt.Printf("Failed to update user details...")
		return false
	}

}

func (r *Concreterepo) FindOpenJobs(company string) []Data {
	//there will be a connection to databse to fect jobs
	jobsModel := &models.ConcreteJobs{}
	jobsArray, error := jobsModel.FetchJobs()
	if error != nil {
		fmt.Printf("Error in fetching jobs")
		return make([]Data, 0)
	}
	arr := *jobsArray
	response := make([]Data, 0)
	for x := 0; x < len(arr); x++ {
		currentJob := arr[x]
		//finding what to display
		finalData := Data{
			JobsTitle:       currentJob.JobTitle,
			PostedBy:        currentJob.PostedBy.RefreeAuth.Name,
			TotalApplicants: currentJob.TotalApplicants,
			Compnay:         currentJob.PostedBy.RefreeCompany,
		}
		response = append(response, finalData)
	}
	companyJobs := make([]Data, 0)
	if company == "" {
		//in actual endpoint we have to add a database connection here, since we are using a dummy
		//system we can just give sample jobs

		//jobsPreprocessing, this can be done a sepearte folderr of utils for time being doimg here

		return response

	} else {

		for x := 0; x < len(response); x++ {
			currentJob := response[x]
			if currentJob.Compnay == company {
				companyJobs = append(companyJobs, response[x])
			}

		}
		fmt.Printf("Fetching company jobs")
		return companyJobs
	}

}

func (r *Concreterepo) Login(username string, password string) bool {
	//
	user := &models.ConcreteUser{}
	checkIfExisting := user.CheckUser(username, password)
	if checkIfExisting == false {
		fmt.Printf("Your user does not exist or password is wrong please signup...")
		return false
	} else {
		fmt.Printf("Successfully Autheticated")

		return true
	}
}

func (r *Concreterepo) Signup(name string, username string, password string, id string, isrefree bool) bool {
	user := &models.ConcreteUser{}
	resp := user.Authenticate(name, username, password, id, isrefree)
	if resp == false {
		fmt.Printf("Failed to signup either user already exist or lenght of password is less than 8")
		return false
	} else {
		return true
	}
}
