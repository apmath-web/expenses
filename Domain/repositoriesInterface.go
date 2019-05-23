package Domain

type repositoryInterface interface {
	GetModel(id int) HelloWorldApplicationModel
	PutModel(model HelloWorldApplicationModel) int
}

type ClientFetchInterface interface {
	Fetch(id int) (PersonDomainModelInterface, error)
}

type CalculateInterface interface {
	GetMaxValue() float64
	SetMaxValue(maxValue float64)
	GetDecimalPlaces() int
	SetDecimalPlaces(dp int)
	Rounder(maxValue float64)
	CalculateSum(persons []PersonDomainModelInterface) error
}