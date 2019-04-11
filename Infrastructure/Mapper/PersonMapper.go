package Mapper

import (
	"github.com/apmath-web/expenses/Domain"
	"github.com/apmath-web/expenses/Domain/Models"
	"github.com/apmath-web/expenses/Infrastructure/applicationModels"
)

func SumWageJob (jobs []applicationModels.JobApplicationModel) int {
	var sum int
	for _, value := range jobs{
		sum += value.Wage
	}
	return sum
}

func PersonApplicationMapper (am applicationModels.PersonApplicationModel) Domain.PersonDomainModelInterface {
	DomainModel := new(Models.PersonDomainModel)
	DomainModel.FirstName = am.FirstName
	DomainModel.LastName = am.LastName
	DomainModel.SumWage = SumWageJob(am.Jobs)
	return DomainModel
}
