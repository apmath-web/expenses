package mapper

import (
	"github.com/apmath-web/expenses/Application/viewModels"
	"github.com/apmath-web/expenses/Domain"
	"github.com/apmath-web/expenses/Domain/models"
)

func IdsViewMapper(vm viewModels.IdsViewModel, clientId int) Domain.IdsDomainModelInterface {
	return models.GenIds(clientId, vm.GetCoborrowersIdSlice())
}
