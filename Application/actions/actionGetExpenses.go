package actions

import (
	"github.com/apmath-web/expenses/Domain/models"
	"github.com/apmath-web/expenses/Domain/services"
	"github.com/apmath-web/expenses/Infrastructure"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetExpenses(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		//do something
		return
	}

	tempSlice := []int{2, 3, 4}
	dm := models.GenIds(id, tempSlice)
	service := services.CalculationService{}

	clientFetchService := Infrastructure.GenClientFetchService()
	service.GenCalculationService(clientFetchService)
	ei, err := service.Calculate(dm)

	amount := ei.GetAmount()
}
