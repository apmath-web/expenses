package routing

import (
	"github.com/apmath-web/expenses/Application/actions"
	"github.com/gin-gonic/gin"
)

func GenRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/:clientId", actions.GetExpenses)
	}
	return router
}
