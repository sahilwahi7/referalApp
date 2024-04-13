package models

import (
	"fmt"

	"github.com/sahilwahi7/referalApp/database"
)

type ConcreteJobs struct {
	Jobs *[]database.Jobs
}

type Jobs interface {
	FetchJobs(company string) (*[]database.Jobs, error)
}

// func(j *Jobs) AddNewJob(job database.Jobs){
// 	//currenlty not optimizing code we can use a map instead of array to make things work
//     j.Jobs=append(j.Jobs,job)
// }

// func(j *Jobs) CheckAllJobs(company string)(*Jobs,error){
// 	jobs,err:=fetchJobs(company)
// 	if err!=nil{
// 		fmt.Printf("Error fetching jobs")
// 		return jobs,err
// 	}else{
// 		fmt.Printf("Successfully fetched jobs")
// 		return jobs,nil

// 	}

// }

func (j *ConcreteJobs) FetchJobs() (*[]database.Jobs, error) {
	//number of current jobs
	jobsArray := []database.Jobs{
		{
			ID:              "1",
			TotalApplicants: 80,
			JobTitle:        "Software engineer",
			PostedBy: &database.Refree{
				ID:            "124",
				RefreeCompany: "Sprinklr",
				RefreeAuth: &database.RefreeAuth{
					ID:       "abc456",
					Email:    "xx@gmail.com",
					Password: "xxxxxx",
					Name:     "sahil wahi",
				},
			},
		},
		{
			ID:              "2",
			TotalApplicants: 10,
			JobTitle:        "Golang developer",
			PostedBy: &database.Refree{
				ID:            "125",
				RefreeCompany: "Phenom",
				RefreeAuth: &database.RefreeAuth{
					ID:       "abc45d6",
					Email:    "xxdf@gmail.com",
					Password: "xxxxxx",
					Name:     "wahi",
				},
			},
		},
		{
			ID:              "3",
			TotalApplicants: 12,
			JobTitle:        "Database engineer",
			PostedBy: &database.Refree{
				ID:            "1245",
				RefreeCompany: "Darwinbox",
				RefreeAuth: &database.RefreeAuth{
					ID:       "abc45dd6",
					Email:    "gmail@gmail.com",
					Password: "xxxxxx",
					Name:     "sahil wahi2",
				},
			},
		},
	}
	//passing jobs array to struct
	fmt.Printf("Feched jobs properly.....")
	j.Jobs = &jobsArray //storing jobsareay in Jobs
	return &jobsArray, nil

}

func (j *ConcreteJobs) StoreInDb(job *database.Jobs) {
	jobs := j.Jobs
	*jobs = append(*jobs, *job)
	fmt.Printf("Added job successfully to DB")
}
