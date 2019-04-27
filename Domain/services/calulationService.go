package services

import (
	"github.com/apmath-web/expenses/Domain"
	"github.com/apmath-web/expenses/Domain/models"
	"github.com/apmath-web/expenses/Infrastructure/Mapper"
)

type calculationService struct {
	clientFetcher Domain.ClientFetchInterface
}

func (cs *calculationService) Calculate(ids models.IdsDomainModel) Domain.ExpensesInterface {
	var persons []Domain.PersonDomainModelInterface

	var pam = *cs.clientFetcher.Fetch(ids.ClientId)
	var pdm = Mapper.PersonApplicationMapper(pam)
	persons = append(persons, pdm)

	for _, value := range ids.CoborrowersIdSlice {
		var pam = *cs.clientFetcher.Fetch(value)
		var pdm = Mapper.PersonApplicationMapper(pam)
		persons = append(persons, pdm)
	}

	var maxValue = models.Calculate(persons)
	var expenses = models.GenExpensesDomainModel(maxValue)

	return expenses
}

func (cs *calculationService) GenCalculationService(clientFetch Domain.ClientFetchInterface) {
	cs.clientFetcher = clientFetch
}
