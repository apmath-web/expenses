package actions

import (
	"encoding/json"
	"github.com/apmath-web/expenses/Application/Validation"
	"github.com/apmath-web/expenses/Application/viewModels"
	"github.com/apmath-web/expenses/Domain/models"
	"github.com/apmath-web/expenses/Domain/services"
	"github.com/apmath-web/expenses/Infrastructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetExpenses(c *gin.Context) {

	vm := viewModels.IdsViewModel{}

	if err := c.BindJSON(&vm); err != nil {
		validator := Validation.GenValidation()
		validator.SetMessage("validation error")
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}

	if !vm.Validate() {
		validator := vm.GetValidation()
		validator.SetMessage("validation error")
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}

	dm := models.GenIds(vm.GetClienId(), vm.GetCoborrowersIdSlice())
	service := services.CalculationService{}

	clientFetchService := Infrastructure.GenClientFetchService()
	service.GenCalculationService(clientFetchService)
	ei, err := service.Calculate(dm)

	if err != nil {
		validator := vm.GetValidation()
		validator.SetMessage("validation error")
		validator.AddMessage(Validation.GenMessage("clientId", err.Error()))
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}

	evm := new(viewModels.ExpensesViewModel)
	evm.Hydrate(ei)

	c.JSON(http.StatusOK, evm)
}
