package mapper

import (
	"github.com/apmath-web/expenses/Application/viewModels"
	"github.com/apmath-web/expenses/Domain/models"
	"math/rand"
	"testing"
)

func TestIdsViewMapper(t *testing.T) {
	id := rand.Int()
	idSlice := []int{rand.Int(), rand.Int(), rand.Int()}

	vm := viewModels.IdsViewModel{}
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
