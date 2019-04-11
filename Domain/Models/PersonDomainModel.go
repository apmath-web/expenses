package Models

type PersonDomainModel struct {
	FirstName string
	LastName  string
	SumWage   int
	MaxValue  int
}

func (person *PersonDomainModel) GetFirstName() string {
	return person.FirstName
}

func (person *PersonDomainModel) GetLastName() string {
	return person.FirstName
}

func (person *PersonDomainModel) GetSumWage() int {
	return person.SumWage
}


