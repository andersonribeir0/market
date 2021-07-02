package handler

import (
	"github.com/gin-gonic/gin"
)

type HealthHandler struct { }

func (*HealthHandler) GetHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"ok": true,
	})
}
