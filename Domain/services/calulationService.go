package services

import (
	"github.com/apmath-web/expenses/Domain"
	"github.com/apmath-web/expenses/Domain/models"
)

type CalculationService struct {
	clientFetcher Domain.ClientFetchInterface
}

func (cs *CalculationService) Calculate(ids models.IdsDomainModel) (Domain.ExpensesInterface, error) {
	var persons []Domain.PersonDomainModelInterface

	var pdm, err = cs.clientFetcher.Fetch(ids.ClientId)
	if err != nil {
		return nil, err
	}
	persons = append(persons, pdm)

	for _, value := range ids.CoborrowersIdSlice {
		var pdm, err = cs.clientFetcher.Fetch(value)
		if err != nil {
			return nil, err
		}
		persons = append(persons, pdm)
	}

	var maxValue = models.Calculate(persons)
	var expenses = models.GenExpensesDomainModel(maxValue)

	return expenses, nil
}

func (cs *CalculationService) GenCalculationService(clientFetch Domain.ClientFetchInterface) {
	cs.clientFetcher = clientFetch
}
