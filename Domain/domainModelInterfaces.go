package Domain

type HelloWorldApplicationModel interface {
	GetMessage() string
	SetMessage(message string)
}

type PersonDomainModelInterface interface {
	GetFirstName() string
	GetLastName()  string
	GetSumWage()   int
}
