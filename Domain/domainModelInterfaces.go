package Domain

type HelloWorldApplicationModel interface {
	GetMessage() string
	SetMessage(message string)
}

type PersonDomainModelInterface interface {
	GetFirstName() string
	GetLastName() string
	GetJobs() []JobDomainModelInterface
}

type JobDomainModelInterface interface {
	GetName() string
	GetWage() int
}

type IdsModelInterface interface {
	GetClientId() int
	GetCoborrowersIdSlice() []int
}