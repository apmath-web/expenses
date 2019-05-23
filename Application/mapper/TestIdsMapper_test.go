package mapper

import (
	"github.com/apmath-web/expenses/Application/viewModels"
	"github.com/apmath-web/expenses/Domain/models"
	"math/rand"
	"testing"
)

func TestIdsViewMapper(t *testing.T) {
	id := rand.Int()
	coborrowersIds := []int{rand.Int(), rand.Int(), rand.Int()}

	vm := viewModels.IdsViewModel{}
	vm.ClientId = id
	vm.CoborrowersIdSlice = coborrowersIds

	dm := models.IdsDomainModel{}
	dm.ClientId = id
	dm.CoborrowersIdSlice = coborrowersIds

	resultId := IdsViewMapper(vm).GetClientId()
	if resultId != dm.ClientId {
		t.Error("For", vm.ClientId, "expected", dm.ClientId, "got", resultId)
	}

	resultIdSlice := IdsViewMapper(vm).GetCoborrowersIdSlice()
	for i, id := range resultIdSlice {
		if id != dm.CoborrowersIdSlice[i] {
			t.Error("For", vm.CoborrowersIdSlice[i], "expected", dm.CoborrowersIdSlice[i], "got", id)
			break
		}
	}
}
