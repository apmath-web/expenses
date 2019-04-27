package Mapper

import (
	"github.com/apmath-web/expenses/Domain/models"
	"github.com/apmath-web/expenses/Domain"
	"github.com/apmath-web/expenses/Infrastructure/applicationModels"
)

func PersonApplicationMapper (am applicationModels.PersonApplicationModel) Domain.PersonDomainModelInterface {
	DomainModel := new(models.PersonDomainModel)
	DomainModel.FirstName = am.FirstName
	DomainModel.LastName = am.LastName
	DomainModel.SumWage = am.GetSumWage()
	return DomainModel
}
