package viewModels

import (
	"encoding/json"
	"github.com/apmath-web/clients/Domain"
)

type JsonIds struct {
	//Тут такой джейсон???
	ClientId int `json:"ClientId"`
	CoborrowersIdSlice []int `json:"CoBorrowersIds"`
}

type IdsViewModel struct {
	JsonIds
	//Тут как поключить валдиацию???
}

func (idsViewModel *IdsViewModel) GetClienId() int {
	return idsViewModel.ClientId
}

func (idsViewModel *IdsViewModel) GetCoborrowersIdSlice() []int {
	return idsViewModel.CoborrowersIdSlice
}

func (idsViewModel *IdsViewModel) validateClientId() {
	//if idsViewModel.ClienId < 0 { Почему не работает?
	//	//Валидация и ла ла ла
	//	//idsViewModel.validation.AddMessage(validation.GenMessage("Id", "Is negative"))
	//}
}

func (idsViewModel *IdsViewModel) validateCoBorrowerIdSlice() {
	for _, id := range idsViewModel.CoborrowersIdSlice {
		//if !id.Validate() { //Как это работает???
			//for _, msg := range id.GetValidation().GetMessages() {
		//idsViewModel.validation.AddMessage(msg)
			}
		}
	}
}

func (idsViewModel *IdsViewModel) Validate() bool {
	idsViewModel.validateClientId()
	idsViewModel.validateCoBorrowerIdSlice()
	return true;
}

//func (idsViewModel *IdsViewModel) GetValidation() Domain.ValidationInterface { Это для чего?
//	return &idsViewModel.validation
//}

func (idsViewModel *IdsViewModel) MarshalJSON() (b []byte, e error) { //Тута так?
	return json.Marshal(map[string]interface{}{
		"CientId": idsViewModel.ClientId,
		"CoBorrowersIds": idsViewModel.CoborrowersIdSlice,
	})
}

func (idsViewModel *IdsViewModel) UnmarshalJSON(b []byte) error {
	tmpClient := JsonIds{}
	err := json.Unmarshal(b, &tmpClient)
	if err := json.Unmarshal(b, &tmpClient); err != nil {
		return err
	}
	idsViewModel.JsonIds = tmpClient
	return err
}

//func (idsViewModel *IdsViewModel) Hydrate(client Domain.viewModelInreface) { Тут что должно быть???
//	c.FirstName = client.GetFirstName()
//	c.LastName = client.GetLastName()
//	c.BirthDate = client.GetBirthDate()
//	c.Sex = client.GetSex()
//	c.MaritalStatus = client.GetMaritalStatus()
//	c.Children = client.GetChildren()
//	for _, job := range client.GetJobs() {
//		tmpJob := JobViewModel{}
//		tmpJob.Hydrate(job)
//		c.Jobs = append(c.Jobs, tmpJob)
//	}

//}
