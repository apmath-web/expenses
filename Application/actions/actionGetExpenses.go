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
	"strconv"
)

func GetExpenses(c *gin.Context) {

	clientId, err := strconv.Atoi(c.Param("clientId"))
	if err != nil {
		validator := Validation.GenValidation()
		validator.SetMessage("Client's ID must be numeric")
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}

	if clientId < 0 {
		validator := Validation.GenValidation()
		validator.SetMessage("Param error")
		validator.AddMessage(Validation.GenMessage("clientId", "Is negative"))
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}

	vm := viewModels.IdsViewModel{}

	if err := c.BindJSON(&vm); err != nil {
		validator := Validation.GenValidation()
		validator.SetMessage("Body error")
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}

	if !vm.Validate() {
		validator := vm.GetValidation()
		validator.SetMessage("Validation error")
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}

	for _, id := range vm.CoborrowersIdSlice {
		if clientId == id {
			validator := Validation.GenValidation()
			validator.SetMessage("Validation error")
			validator.AddMessage(Validation.GenMessage("coBorrowers", "Client's ID is equal to coborrower's ID"))
			str, _ := json.Marshal(validator)
			c.String(http.StatusBadRequest, string(str))
			return
		}
	}
	dm := models.GenIds(clientId, vm.GetCoborrowersIdSlice())
	service := services.CalculationService{}

	clientFetchService := Infrastructure.GetServiceManager().GetClientFetchService()
	service.GenCalculationService(clientFetchService)
	ei, err := service.Calculate(dm)

	if err != nil {
		validator := Validation.GenValidation()
		validator.SetMessage(err.Error())
		str, _ := json.Marshal(validator)
		if err.Error() == Infrastructure.NotAvaliableMessage {
			c.String(http.StatusInternalServerError, string(str))
			return
		}
		if err.Error() == Infrastructure.BadRequestMessage {
			c.String(http.StatusBadRequest, string(str))
			return
		}
		if err.Error() == Infrastructure.NotFoundMessage {
			c.String(http.StatusNotFound, string(str))
			return
		}
		c.String(http.StatusBadRequest, string(str))
		return
	}

	evm := new(viewModels.ExpensesViewModel)
	evm.Hydrate(ei)

	c.JSON(http.StatusOK, evm)
}
