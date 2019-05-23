package services

import (
	"github.com/apmath-web/expenses/Domain"
	"github.com/apmath-web/expenses/Domain/models"
)

type CalculationService struct {
	clientFetcher Domain.ClientFetchInterface
}

func (cs *CalculationService) Calculate(ids Domain.IdsDomainModelInterface) (Domain.ExpensesInterface, error) {
	var persons []Domain.PersonDomainModelInterface
	var calc = CalculateService{}
	var pdm, err = cs.clientFetcher.Fetch(ids.GetClientId())

	if err != nil {
		return nil, err
	}

	persons = append(persons, pdm)

	for _, value := range ids.GetCoborrowersIdSlice() {
		var pdm, err = cs.clientFetcher.Fetch(value)
		if err != nil {
			return nil, err
		}
		persons = append(persons, pdm)
	}

	err = calc.Calculate(persons)

	if err != nil {
		return nil, err
	}

	var expenses = models.GenExpensesDomainModel(calc.GetMaxValue())

	return expenses, nil
}

func (cs *CalculationService) GenCalculationService(clientFetch Domain.ClientFetchInterface) {
	cs.clientFetcher = clientFetch
}
