package mapper

import (
	"github.com/apmath-web/expenses/Application/viewModels"
	"github.com/apmath-web/expenses/Domain/models"
	"testing"
)

func TestIdsViewMapper(t *testing.T) {
	vm := viewModels.IdsViewModel{}
	id := 1
	idSlice := []int{2, 3, 4}
	vm.ClientId = id
	vm.CoborrowersIdSlice = idSlice

	dm := models.IdsDomainModel{}
	dm.ClientId = id
	dm.CoborrowersIdSlice = idSlice

	resultId := IdsViewMapper(vm).GetClientId()
	if resultId != dm.ClientId {
		t.Error("For", vm.ClientId, "expected", dm.ClientId, "got", resultId)
	}

	resultIdSlice := IdsViewMapper(vm).GetCoborrowersIdSlice()
	for i := 0; i < len(idSlice); i++ {
		if resultIdSlice[i] != dm.CoborrowersIdSlice[i] {
			t.Error("For", vm.CoborrowersIdSlice[i], "expected", dm.CoborrowersIdSlice[i], "got", resultIdSlice[i])
			break
		}
	}
}
