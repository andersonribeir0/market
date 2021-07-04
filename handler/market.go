package handler

import (
	"fmt"
	"github.com/andersonribeir0/market/logger"
	"github.com/andersonribeir0/market/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MarketHandler struct {
	requestId  string
	marketRepo repository.IMarketRepository
	Logger     *logger.Log
}


func (m *MarketHandler) setRequestId(c *gin.Context) {
	if m.requestId = c.Request.Header.Get("X-Request-Id"); c.Request.Header.Get("X-Request-Id") == "" {
		m.requestId = uuid.New().String()
		c.Request.Header.Set("X-Request-Id", m.requestId)
	}
}

func (m *MarketHandler) Initialize(c *gin.Context) {
	m.setRequestId(c)
	m.setMarketRepo()
}

func (m *MarketHandler) Get(c *gin.Context) {
	m.Logger.Info("Getting market record")
	m.Initialize(c)
	marketId := c.Param("id")
	if marketId == "" {
		c.JSON(400, gin.H {
			"error": "missing marketId ",
		})
		return
	}

	item, err := m.marketRepo.GetItem(marketId)
	if err != nil {
		c.JSON(400, gin.H {
			"error": err.Error(),
		})
		return
	}

	if item.Id == nil {
		c.JSON(404, gin.H{})
		return
	}

	c.JSON(200, item)
}

func (m *MarketHandler) GetByDistCode(c *gin.Context) {
	m.Initialize(c)
	codDist := c.Query("codDist")
	if codDist == "" {
		c.JSON(400, gin.H {
			"error": "missing codDist",
		})
		return
	}

	item, err := m.marketRepo.GetItemsByDistrictId(codDist)
	if err != nil {
		c.JSON(400, gin.H {
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, item)
}

func (m *MarketHandler) setMarketRepo() {
	if m.marketRepo == nil {
		marketRepo := repository.MarketRepository{}
		err := marketRepo.New()
		if err != nil {
			m.Logger.Error("Error creating repo instance",
				err,
				fmt.Sprintf("requestId:%s", m.requestId))
		}
		m.marketRepo = &marketRepo
	}
}

