package webserver

import (

	"github.com/andersonribeir0/market/handler"
	"github.com/andersonribeir0/market/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getRequestId(c *gin.Context) {
	requestId := c.GetHeader("X-Request-Id")
	if requestId == "" {
		requestId = uuid.New().String()
		c.Request.Header.Add("X-Request-Id", requestId)
	}
	c.Next()
}


func getRouter() *gin.Engine {
	router := gin.Default()
	log := logger.NewLogger("market_app")
	healthRoute := handler.HealthHandler{  }
	marketRoute := handler.MarketHandler{ Logger: log }

	router.Use(getRequestId)
	router.GET("/health", healthRoute.GetHealth)

	v1 := router.Group("/v1")
	{
		v1.GET("/market/:id", marketRoute.Get)
		v1.GET("/market", marketRoute.GetByDistCode)
	}
	return router
}

func Run(port string) {
	router := getRouter()
	router.Run(port)
}