package Domain

type PersonDomainModelInterface interface {
	GetFirstName() string
	GetLastName() string
	GetJobs() []JobdomainModelInterface
}

type JobDomainModelInterface interface {
	GetName() string
	GetWage() int
}
