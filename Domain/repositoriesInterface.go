package Domain

import "github.com/apmath-web/expenses/Infrastructure/applicationModels"

type repositoryInterface interface {
	GetModel(id int) HelloWorldApplicationModel
	PutModel(model HelloWorldApplicationModel) int
}

type ExpensesInterface interface {
	GetAmount() float64
	SetAmount(amount float64)
}

type ClientFetchInterface interface {
	Fetch(id int) *applicationModels.PersonApplicationModel
}
