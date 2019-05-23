package mapper

import (
	"github.com/apmath-web/expenses/Application/viewModels"
	"github.com/apmath-web/expenses/Domain/models"
	"math/rand"
	"testing"
)

func TestIdsViewMapper(t *testing.T) {
	clientId := rand.Int()
	coborrowersIds := []int{rand.Int(), rand.Int(), rand.Int()}

	vm := viewModels.IdsViewModel{}
	vm.ClientId = clientId
	vm.CoborrowersIdSlice = coborrowersIds

	dm := models.GenIds(clientId, coborrowersIds)

	resultId := IdsViewMapper(vm).GetClientId()
	if resultId != dm.GetClientId() {
		t.Error("For", vm.ClientId, "expected", dm.GetClientId(), "got", resultId)
	}

	resultIdSlice := IdsViewMapper(vm).GetCoborrowersIdSlice()
	for i, id := range resultIdSlice {
		if id != dm.GetCoborrowersIdSlice()[i] {
			t.Error("For", vm.CoborrowersIdSlice[i], "expected", dm.GetCoborrowersIdSlice()[i], "got", id)
			break
		}
	}
}
