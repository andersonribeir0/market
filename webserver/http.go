package webserver

import (
	"github.com/andersonribeir0/market/handler"
	"github.com/andersonribeir0/market/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var log = logger.NewLogger()

func getRequestId(c *gin.Context) {
	requestId := c.GetHeader("X-Request-Id")
	if requestId == "" {
		requestId = uuid.New().String()
		c.Request.Header.Add("X-Request-Id", requestId)
	}
	log.WithField("requestId", requestId)
	c.Next()
}

func getRouter() *gin.Engine {
	router := gin.Default()
	log.WithField("service", "market_app")
	healthRoute := handler.HealthHandler{  }
	marketRoute := handler.MarketHandler{ Logger: log }

	router.Use(getRequestId)

	router.GET("/health", healthRoute.GetHealth)
	v1 := router.Group("/v1")
	{
		v1.GET("/market/:id", marketRoute.Get)
		v1.GET("/market", marketRoute.GetByDistCode)
		v1.DELETE("/market/:id", marketRoute.Delete)
		v1.PUT("/market", marketRoute.Put)
	}
	return router
}

func Run(port string) {
	router := getRouter()
	router.Run(port)
}