package Models

type PersonDomainModel struct {
	FirstName string
	LastName  string
	SumWage   int
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

func Calculate(persons []PersonDomainModel) float64 {
	var MaxValue float64
	var SumWagePerson int
	for _, value := range persons {
		SumWagePerson += value.SumWage
	}
	MaxValue = float64(SumWagePerson) / float64(len(persons))
	return MaxValue
}


