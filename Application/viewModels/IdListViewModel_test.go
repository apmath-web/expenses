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
}
