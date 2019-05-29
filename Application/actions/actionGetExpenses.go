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
		c.String(http.StatusBadRequest, string(err.Error()))
		return
	}

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

	dm := models.GenIds(clientId, vm.GetCoborrowersIdSlice())
	service := services.CalculationService{}

	clientFetchService := Infrastructure.GenClientFetchService()
	service.GenCalculationService(clientFetchService)
	ei, err := service.Calculate(dm)

	if err != nil {
		if err.Error() == "clients service not available" {
			c.String(http.StatusInternalServerError, string(err.Error()))
			return
		}
		if err.Error() == "bad request" {
			c.String(http.StatusBadRequest, string(err.Error()))
			return
		}
		if err.Error() == "client not found" {
			c.String(http.StatusNotFound, string(err.Error()))
			return
		}
		c.String(http.StatusBadRequest, string(err.Error()))
		return
	}

	evm := new(viewModels.ExpensesViewModel)
	evm.Hydrate(ei)

	c.JSON(http.StatusOK, evm)
}
