package database

// this is basic job interface
type Jobs struct {
	ID              string
	TotalApplicants int64
	PostedBy        *Refree
	JobTitle        string
}

//now we want to make a concrete implementation of jobs for refree and cadidate

type JobsRefree struct { //jobs  for refree's pane
	Jobs
	candidates []Candidate
}

type JobsCandidate struct { //jobs  for  cadidate's pane
	Jobs
	noOfCandidates int
}
