package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/andersonribeir0/market/logger"
	"github.com/andersonribeir0/market/model"
	"github.com/andersonribeir0/market/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
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

func (m *MarketHandler) Put(c *gin.Context) {
	m.Initialize(c)
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		msg := "Invalid data. JSON data is required."
		m.Logger.Error(msg, err)
		c.JSON(400, gin.H{
			"error": msg,
		})
		return
	}

	var market model.Record
	validate := validator.New()
	if err = json.Unmarshal(jsonData, &market); err != nil {
		msg := "Error unmarshalling request data."
		m.Logger.Error(msg, err)
		c.JSON(500, gin.H{
			"error": msg,
		})
		return
	}
	if err = validate.Struct(market); err != nil {
		msg := "Invalid data. JSON data is required."
		m.Logger.Error(msg, err)
		c.JSON(400, gin.H{
			"error": msg,
		})
		return
	}

	m.Logger.Info(fmt.Sprintf("Inserting record %s", string(jsonData)))
	if err := m.marketRepo.Save(market); err != nil {
		msg := fmt.Sprintf("Error saving record %s", string(jsonData))
		m.Logger.Error(msg, err)
		c.JSON(500, gin.H{
			"error": msg,
		})
		return
	}

	c.JSON(200, gin.H{})
	return
}

func (m *MarketHandler) Get(c *gin.Context) {
	m.Initialize(c)
	marketId := c.Param("id")
	m.Logger.Info(fmt.Sprintf("Getting record by id %s", marketId))
	if marketId == "" {
		c.JSON(400, gin.H {
			"error": "missing marketId ",
		})
		return
	}

	item, err := m.marketRepo.GetItem(marketId)
	if err != nil {
		m.Logger.Error(fmt.Sprintf("Error getting record by id %s", marketId), err)
		c.JSON(400, gin.H {
			"error": err.Error(),
		})
		return
	}

	if item.Id == nil {
		m.Logger.Info(fmt.Sprintf("Record with id %s does not exists", marketId))
		c.JSON(404, gin.H{})
		return
	}

	data, _ := json.Marshal(item)
	m.Logger.Info(fmt.Sprintf("Got %s", string(data)))
	c.JSON(200, item)
}

func (m *MarketHandler) GetByDistCode(c *gin.Context) {
	m.Initialize(c)
	codDist := c.Query("codDist")
	if codDist == "" {
		m.Logger.Error("Missing codDist", errors.New("missing codDist"))
		c.JSON(400, gin.H {
			"error": "missing codDist",
		})
		return
	}

	items, err := m.marketRepo.GetItemsByDistrictId(codDist)
	if err != nil {
		m.Logger.Error(fmt.Sprintf("Error getting records by codDist %s", codDist), err)
		c.JSON(400, gin.H {
			"error": err.Error(),
		})
		return
	}
	data, _ := json.Marshal(items)
	m.Logger.Info(fmt.Sprintf("Got %s", string(data)))
	c.JSON(200, items)
}

func (m *MarketHandler) Delete(c *gin.Context) {
	m.Initialize(c)
	marketId := c.Param("id")
	m.Logger.Info(fmt.Sprintf("Deleting record by id %s", marketId))
	if marketId == "" {
		c.JSON(400, gin.H {
			"error": "missing marketId ",
		})
		return
	}

	if err := m.marketRepo.Delete(marketId); err != nil {
		msg := fmt.Sprintf("Error deleting id %s", marketId)
		m.Logger.Error(msg, err)
		c.JSON(500, gin.H {
			"error": msg,
		})
		return
	}

	c.JSON(200, gin.H{})
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

