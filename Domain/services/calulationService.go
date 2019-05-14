package services

import (
	"github.com/apmath-web/expenses/Domain"
	"github.com/apmath-web/expenses/Domain/models"
)

type calculationService struct {
	clientFetcher Domain.ClientFetchInterface
}

func (cs *calculationService) Calculate(ids models.IdsDomainModel) Domain.ExpensesInterface {
	var persons []Domain.PersonDomainModelInterface

	var pdm = cs.clientFetcher.Fetch(ids.ClientId)
	persons = append(persons, &pdm)

	for _, value := range ids.CoborrowersIdSlice {
		var pdm = cs.clientFetcher.Fetch(value)
		persons = append(persons, &pdm)
	}

	var maxValue = models.Calculate(persons)
	var expenses = models.GenExpensesDomainModel(maxValue)

	return expenses
}

func (cs *calculationService) GenCalculationService(clientFetch Domain.ClientFetchInterface) {
	cs.clientFetcher = clientFetch
}
