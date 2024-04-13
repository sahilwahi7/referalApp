package database

type Candidate struct {
	ID              string
	NoOfJobsApplied []JobsCandidate
}

type CandidateAuth struct {
	id       string
	email    string
	password string
}
