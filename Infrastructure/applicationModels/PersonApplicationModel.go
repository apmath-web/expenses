package applicationModels

import "github.com/apmath-web/expenses/Domain"

type PersonApplicationModel struct {
	Id        int                   `json:"id"`
	FirstName string                `json:"firstName"`
	LastName  string                `json:"lastName"`
	Jobs      []JobApplicationModel `json:"jobs"`
}

func (person *PersonApplicationModel) GetId() int {
	return person.Id
}

func (person *PersonApplicationModel) GetFirstName() string {
	return person.FirstName
}

func (person *PersonApplicationModel) GeLastName() string {
	return person.LastName
}

func (person *PersonApplicationModel) GetWage() int {
	return person.Wage
}

func GenHelloWorldApplicationModel(message string) Domain.HelloWorldApplicationModel {
	hw := new(HelloWorld)
	hw.SetMessage(message)
	return hw
}
