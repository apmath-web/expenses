package viewModels

import (
	"testing"
)

func TestIdsViewModel_Validate(t *testing.T) {

	type TestData struct {
		clientId       int
		coborrowersIds []int
		field          string
		text           string
		isErrorCase    bool
	}

	dataSlice := []TestData{}

	//Ok cases
	{
		clientId := 1
		coborrowersIds := []int{2}
		data := TestData{clientId, coborrowersIds, "",
			"", false}
		dataSlice = append(dataSlice, data)
	}
	{
		clientId := 2
		coborrowersIds := []int{1, 3, 4, 5}
		data := TestData{clientId, coborrowersIds, "",
			"", false}
		dataSlice = append(dataSlice, data)
	}
	{
		clientId := 20
		coborrowersIds := []int{2, 3, 4, 5, 412, 43}
		data := TestData{clientId, coborrowersIds, "",
			"", false}
		dataSlice = append(dataSlice, data)
	}

	//Client ID negative cases
	{
		clientId := -1
		coborrowersIds := []int{2}
		data := TestData{clientId, coborrowersIds, "clienId",
			"Is negative", true}
		dataSlice = append(dataSlice, data)
	}
	{
		clientId := -2
		coborrowersIds := []int{1, 3, 4, 5}
		data := TestData{clientId, coborrowersIds, "clienId",
			"Is negative", true}
		dataSlice = append(dataSlice, data)
	}
	{
		clientId := -20
		coborrowersIds := []int{2, 3, 4, 5, 412, 43}
		data := TestData{clientId, coborrowersIds, "clienId",
			"Is negative", true}
		dataSlice = append(dataSlice, data)
	}

	//Coborrower's ID negative cases
	{
		clientId := 1
		coborrowersIds := []int{-2, 90}
		data := TestData{clientId, coborrowersIds, "coBorrowers",
			"Is negative", true}
		dataSlice = append(dataSlice, data)
	}
	{
		clientId := 2
		coborrowersIds := []int{1, 3, 4, -5}
		data := TestData{clientId, coborrowersIds, "coBorrowers",
			"Is negative", true}
		dataSlice = append(dataSlice, data)
	}
	{
		clientId := 20
		coborrowersIds := []int{2, 3, -4, 5, 412, 43}
		data := TestData{clientId, coborrowersIds, "coBorrowers",
			"Is negative", true}
		dataSlice = append(dataSlice, data)
	}

	//Coborrower's ID's are eaqual cases
	{
		clientId := 1
		coborrowersIds := []int{2, 3, 90, 3}
		data := TestData{clientId, coborrowersIds, "coBorrowers",
			"Some coborrower's IDs are equal to each other", true}
		dataSlice = append(dataSlice, data)
	}
	{
		clientId := 2
		coborrowersIds := []int{1, 3, 3, 5}
		data := TestData{clientId, coborrowersIds, "coBorrowers",
			"Some coborrower's IDs are equal to each other", true}
		dataSlice = append(dataSlice, data)
	}
	{
		clientId := 20
		coborrowersIds := []int{2, 3, 4, 5, 412, 2}
		data := TestData{clientId, coborrowersIds, "coBorrowers",
			"Some coborrower's IDs are equal to each other", true}
		dataSlice = append(dataSlice, data)
	}

	//Coborrower's ID is eaqual to client's ID
	{
		clientId := 1
		coborrowersIds := []int{1, 2, 90, 3}
		data := TestData{clientId, coborrowersIds, "coBorrowers",
			"Coborrower's ID can't be equal to Client's ID", true}
		dataSlice = append(dataSlice, data)
	}
	{
		clientId := 2
		coborrowersIds := []int{1, 3, 2, 5}
		data := TestData{clientId, coborrowersIds, "coBorrowers",
			"Coborrower's ID can't be equal to Client's ID", true}
		dataSlice = append(dataSlice, data)
	}
	{
		clientId := 20
		coborrowersIds := []int{2, 3, 4, 5, 412, 20}
		data := TestData{clientId, coborrowersIds, "coBorrowers",
			"Coborrower's ID can't be equal to Client's ID", true}
		dataSlice = append(dataSlice, data)
	}

	for _, currentData := range dataSlice {

		vm := IdsViewModel{}
		vm.ClientId = currentData.clientId
		vm.CoborrowersIdSlice = currentData.coborrowersIds
		vm.Validate()

		if currentData.isErrorCase {
			validation := vm.GetValidation()
			message := validation.GetMessages()[0]
			field := message.GetField()
			text := message.GetText()

			if text != currentData.text || field != currentData.field {
				t.Error("Expected", currentData.text, "in", currentData.field,
					"got", text, "in", field)
			}
		} else {
			if !vm.Validate() {
				t.Error("Expected true case",
					"got false case")
			}
		}
	}
}
