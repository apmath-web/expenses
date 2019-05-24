package services

import (
	"github.com/apmath-web/expenses/Domain"
	"github.com/apmath-web/expenses/Domain/models"
	"testing"
)

func TestCalculate (t *testing.T) {
	var calc = CalculateService{}
	var maxValue float64

	persons := []models.PersonDomainModel{
		models.PersonDomainModel {
			FirstName: "Ivan",
			LastName:  "Ivanov",
			SumWage:   1000,
		},
		models.PersonDomainModel {
			FirstName: "Petr",
			LastName: "Petrov",
			SumWage: 2000,
		},
		models.PersonDomainModel {
			FirstName: "Artem",
			LastName: "Artemov",
			SumWage: 3000,
		},
		models.PersonDomainModel {
			FirstName: "Timur",
			LastName: "Timurov",
			SumWage: 4000,
		},
		models.PersonDomainModel {
			FirstName: "Pasha",
			LastName: "Pashov",
			SumWage: 5000,
		},
	}

	personInf := make([]Domain.PersonDomainModelInterface, len(persons))

	for i, v := range persons{
		personInf[i] = v
	}

	_ = calc.Calculate(personInf)
	maxValue = calc.GetMaxValue()

	if maxValue != 3000 {
		t.Error("Incorrect answer for simple case", maxValue, calc.len, calc.sumWagePerson)
	}

	persons[0].SumWage = -1

	err := calc.Calculate(personInf)

	if err != nil  {
		t.Error("Negative value")
	}
}
