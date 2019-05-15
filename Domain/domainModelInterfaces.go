package Domain

type HelloWorldApplicationModel interface {
	GetMessage() string
	SetMessage(message string)
}

type PersonDomainModelInterface interface {
	GetFirstName() string
	GetLastName() string
	GetSumWage() int
}

type IdsDomainModelInterface interface {
	GetClientId() int
	GetCoborrowersIdSlice() []int
}

type ExpensesInterface interface {
	GetAmount() float64
	SetAmount(amount float64)
}
