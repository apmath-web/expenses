package viewModels

import "github.com/apmath-web/expenses/Application/Validation"

type JsonIds struct {
	ClientId           int   `json:"clientId"`    //ID заемщика
	CoborrowersIdSlice []int `json:"coBorrowers"` //Cписок ID созаемщиков
}

type IdsViewModel struct {
	JsonIds
	validation Validation.Validation
}

func (idsViewModel *IdsViewModel) GetClienId() int {
	return idsViewModel.ClientId
}

func (idsViewModel *IdsViewModel) GetCoborrowersIdSlice() []int {
	return idsViewModel.CoborrowersIdSlice
}

func (idsViewModel *IdsViewModel) validateClientId() bool {
	if idsViewModel.ClientId < 0 {
		idsViewModel.validation.AddMessage(Validation.GenMessage("clienId", "Is negative"))
		return false
	}
	return true
}

func (idsViewModel *IdsViewModel) validateCoBorrowerIdSlice() bool {
	for _, id := range idsViewModel.CoborrowersIdSlice {
		if id < 0 {
			idsViewModel.validation.AddMessage(Validation.GenMessage("coBorrowers", "Is negative"))
			return false
		}
	}
	return true
}

func (idsViewModel *IdsViewModel) Validate() bool {
	if idsViewModel.validateClientId() && idsViewModel.validateCoBorrowerIdSlice() {
		return true
	}
	return false
}
