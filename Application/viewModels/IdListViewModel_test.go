package viewModels

import (
	"github.com/apmath-web/expenses/Application/Validation"
	"testing"
)

func TestIdsViewModel_Validate(t *testing.T) {
	vm := IdsViewModel{}

	//OK case
	{
		vm.ClientId = 1
		vm.CoborrowersIdSlice = []int{2, 3, 4, 5}
		if !vm.Validate() {
			t.Error("Expected true, got", vm.Validate())
		}
	}
	{
		vm.ClientId = 2
		vm.CoborrowersIdSlice = []int{1}
		if !vm.Validate() {
			t.Error("Expected true, got", vm.Validate())
		}
	}
	{
		vm.ClientId = 10
		vm.CoborrowersIdSlice = []int{23, 31, 4, 15}
		if !vm.Validate() {
			t.Error("Expected true, got", vm.Validate())
		}
	}
	{
		vm.ClientId = 100
		vm.CoborrowersIdSlice = []int{2, 3, 4, 5, 6, 10, 33}
		if !vm.Validate() {
			t.Error("Expected true, got", vm.Validate())
		}
	}

	//Client's Id negative case
	{
		vm.ClientId = -1
		vm.CoborrowersIdSlice = []int{2, 3, 4, 5}
		vm.Validate()

		trueValidation := Validation.Validation{}
		trueValidation.AddMessage(Validation.GenMessage("clienId", "Is negative"))
		_, trueMessage := trueValidation.GetMessages()[0].MarshalJSON()

		resultValidation := vm.GetValidation()
		_, resultMessage := resultValidation.GetMessages()[0].MarshalJSON()

		if resultMessage != trueMessage {
			t.Error("Expected", trueMessage, "got", resultMessage)
		}
	}
	{
		vm.ClientId = -22
		vm.CoborrowersIdSlice = []int{2, 3, 4, 5}
		vm.Validate()

		trueValidation := Validation.Validation{}
		trueValidation.AddMessage(Validation.GenMessage("clienId", "Is negative"))
		_, trueMessage := trueValidation.GetMessages()[0].MarshalJSON()

		resultValidation := vm.GetValidation()
		_, resultMessage := resultValidation.GetMessages()[0].MarshalJSON()

		if resultMessage != trueMessage {
			t.Error("Expected", trueMessage, "got", resultMessage)
		}
	}
	{
		vm.ClientId = -123
		vm.CoborrowersIdSlice = []int{1, 2, 3}
		vm.Validate()

		trueValidation := Validation.Validation{}
		trueValidation.AddMessage(Validation.GenMessage("clienId", "Is negative"))
		_, trueMessage := trueValidation.GetMessages()[0].MarshalJSON()

		resultValidation := vm.GetValidation()
		_, resultMessage := resultValidation.GetMessages()[0].MarshalJSON()

		if resultMessage != trueMessage {
			t.Error("Expected", trueMessage, "got", resultMessage)
		}
	}

	//Coborrower's Id negative case
	{
		vm.ClientId = 1
		vm.CoborrowersIdSlice = []int{2, -3, 4, 5}
		vm.Validate()

		trueValidation := Validation.Validation{}
		trueValidation.AddMessage(Validation.GenMessage("coBorrowers", "Is negative"))
		_, trueMessage := trueValidation.GetMessages()[0].MarshalJSON()

		resultValidation := vm.GetValidation()
		_, resultMessage := resultValidation.GetMessages()[0].MarshalJSON()

		if resultMessage != trueMessage {
			t.Error("Expected", trueMessage, "got", resultMessage)
		}
	}
	{
		vm.ClientId = 1
		vm.CoborrowersIdSlice = []int{21, 13, 40, 50, -3, 22}
		vm.Validate()

		trueValidation := Validation.Validation{}
		trueValidation.AddMessage(Validation.GenMessage("coBorrowers", "Is negative"))
		_, trueMessage := trueValidation.GetMessages()[0].MarshalJSON()

		resultValidation := vm.GetValidation()
		_, resultMessage := resultValidation.GetMessages()[0].MarshalJSON()

		if resultMessage != trueMessage {
			t.Error("Expected", trueMessage, "got", resultMessage)
		}
	}
	{
		vm.ClientId = 1
		vm.CoborrowersIdSlice = []int{21, 13, 40, 50, 2, 22, -40}
		vm.Validate()

		trueValidation := Validation.Validation{}
		trueValidation.AddMessage(Validation.GenMessage("coBorrowers", "Is negative"))
		_, trueMessage := trueValidation.GetMessages()[0].MarshalJSON()

		resultValidation := vm.GetValidation()
		_, resultMessage := resultValidation.GetMessages()[0].MarshalJSON()

		if resultMessage != trueMessage {
			t.Error("Expected", trueMessage, "got", resultMessage)
		}
	}

	//Coborrower's Id is equal to clients's Id
	{
		vm.ClientId = 1
		vm.CoborrowersIdSlice = []int{1, 2, 4, 5}
		vm.Validate()

		trueValidation := Validation.Validation{}
		trueValidation.AddMessage(Validation.GenMessage("coBorrowers", "Coborrower's ID can't be equal to Client's ID"))
		_, trueMessage := trueValidation.GetMessages()[0].MarshalJSON()

		resultValidation := vm.GetValidation()
		_, resultMessage := resultValidation.GetMessages()[0].MarshalJSON()

		if resultMessage != trueMessage {
			t.Error("Expected", trueMessage, "got", resultMessage)
		}
	}
	{
		vm.ClientId = 1
		vm.CoborrowersIdSlice = []int{13, 2, 4, 5, 6, 1, 3, 5}
		vm.Validate()

		trueValidation := Validation.Validation{}
		trueValidation.AddMessage(Validation.GenMessage("coBorrowers", "Coborrower's ID can't be equal to Client's ID"))
		_, trueMessage := trueValidation.GetMessages()[0].MarshalJSON()

		resultValidation := vm.GetValidation()
		_, resultMessage := resultValidation.GetMessages()[0].MarshalJSON()

		if resultMessage != trueMessage {
			t.Error("Expected", trueMessage, "got", resultMessage)
		}
	}
	{
		vm.ClientId = 1
		vm.CoborrowersIdSlice = []int{13, 2, 4, 5, 6, 10, 3, 5, 1}
		vm.Validate()

		trueValidation := Validation.Validation{}
		trueValidation.AddMessage(Validation.GenMessage("coBorrowers", "Coborrower's ID can't be equal to Client's ID"))
		_, trueMessage := trueValidation.GetMessages()[0].MarshalJSON()

		resultValidation := vm.GetValidation()
		_, resultMessage := resultValidation.GetMessages()[0].MarshalJSON()

		if resultMessage != trueMessage {
			t.Error("Expected", trueMessage, "got", resultMessage)
		}
	}

	//Coborrower's Id's are equal
	{
		vm.ClientId = 1
		vm.CoborrowersIdSlice = []int{2, 2, 4, 5}
		vm.Validate()

		trueValidation := Validation.Validation{}
		trueValidation.AddMessage(Validation.GenMessage("coBorrowers", "Some coborrower's IDs are equal to each other"))
		_, trueMessage := trueValidation.GetMessages()[0].MarshalJSON()

		resultValidation := vm.GetValidation()
		_, resultMessage := resultValidation.GetMessages()[0].MarshalJSON()

		if resultMessage != trueMessage {
			t.Error("Expected", trueMessage, "got", resultMessage)
		}
	}
	{
		vm.ClientId = 1
		vm.CoborrowersIdSlice = []int{1, 2, 4, 5, 10, 11, 12, 10}
		vm.Validate()

		trueValidation := Validation.Validation{}
		trueValidation.AddMessage(Validation.GenMessage("coBorrowers", "Some coborrower's IDs are equal to each other"))
		_, trueMessage := trueValidation.GetMessages()[0].MarshalJSON()

		resultValidation := vm.GetValidation()
		_, resultMessage := resultValidation.GetMessages()[0].MarshalJSON()

		if resultMessage != trueMessage {
			t.Error("Expected", trueMessage, "got", resultMessage)
		}
	}
	{
		vm.ClientId = 111
		vm.CoborrowersIdSlice = []int{1, 2, 4, 1, 10, 1, 12, 10}
		vm.Validate()

		trueValidation := Validation.Validation{}
		trueValidation.AddMessage(Validation.GenMessage("coBorrowers", "Some coborrower's IDs are equal to each other"))
		_, trueMessage := trueValidation.GetMessages()[0].MarshalJSON()

		resultValidation := vm.GetValidation()
		_, resultMessage := resultValidation.GetMessages()[0].MarshalJSON()

		if resultMessage != trueMessage {
			t.Error("Expected", trueMessage, "got", resultMessage)
		}
	}
}
