package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/andersonribeir0/market/model"
	"github.com/andersonribeir0/market/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type MockIMarketRepository struct {
	Save                   func(market model.Record) error
	GetItem                func(id string) (model.Record, error)
	GetItemsByDistrictId   func(id string) ([]model.Record, error)
	Delete                 func(id string) error

}

type MarketRepositoryStub struct {
	MockRepo	MockIMarketRepository
}

func (m MarketRepositoryStub) Save(market model.Record) error {
	return m.MockRepo.Save(market)
}

func (m MarketRepositoryStub) GetItem(id string) (model.Record, error) {
	return m.MockRepo.GetItem(id)
}

func (m MarketRepositoryStub) GetItemsByDistrictId(id string)  ([]model.Record, error) {
	return m.MockRepo.GetItemsByDistrictId(id)
}

func (m MarketRepositoryStub) Delete(id string) error {
	return m.MockRepo.Delete(id)
}

func getGinContextMock(id string, method string, body io.Reader, requestId string, queryParamKey string) (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	if id != "" {
		c.Params = gin.Params{gin.Param{
			Key:   "id",
			Value: id,
		}}
	}
	queryParam := ""
	if queryParamKey != "" {
		queryParam = "?" + queryParamKey + "=" + strconv.Itoa(rand.Int())
	}
	c.Request, _ = http.NewRequest(method, "/" + queryParam, body)
	c.Request.Header.Add("X-Request-Id", requestId)

	return c, w
}

func getBody() io.Reader{
	s, _ := json.Marshal(map[string]interface{}{
		"ID": 190,
		"CODDIST": "50",
		"DISTRITO": "LIMAO",
		"NUMERO": "417.000000",
		"BAIRRO": "VL STA MARIA",
	})
	return bytes.NewBuffer(s)
}

func TestMarketHandler_PutWithInvalidRequest(t *testing.T) {
	mockRepo := &MarketRepositoryStub{
		MockRepo: MockIMarketRepository{
			Save: func(market model.Record) error {
				return nil
			},
		},
	}
	s, _ := json.Marshal(map[string]interface{}{
		"rada": "invalid",
	})
	b := bytes.NewBuffer(s)
	c, _ := getGinContextMock("anyId", "PUT", b, "", "")
	handler := MarketHandler{
		requestId:  "a_request_id",
		marketRepo: mockRepo,
		log:        log.New().WithField("anyKey","anyValue"),
	}

	handler.Put(c)

	if c.Writer.Status() != 400 {
		t.Fatalf("Expected 400. Received %d", c.Writer.Status())
	}
}

func TestMarketHandler_PutWithRepoError(t *testing.T) {
	mockRepo := &MarketRepositoryStub{
		MockRepo: MockIMarketRepository{
			Save: func(market model.Record) error {
				return errors.New("repo error")
			},
		},
	}

	c, _ := getGinContextMock("anyId", "PUT", getBody(), "", "")
	handler := MarketHandler{
		requestId:  "a_request_id",
		marketRepo: mockRepo,
		log:        log.New().WithField("anyKey","anyValue"),
	}

	handler.Put(c)

	if c.Writer.Status() != 500 {
		t.Fatalf("Expected 500. Received %d", c.Writer.Status())
	}
}

func TestMarketHandler_Put(t *testing.T) {
	mockRepo := &MarketRepositoryStub{
		MockRepo: MockIMarketRepository{
			Save: func(market model.Record) error {
				return nil
			},
		},
	}

	c, _ := getGinContextMock("anyId", "PUT", getBody(), "", "")
	handler := MarketHandler{
		requestId:  "a_request_id",
		marketRepo: mockRepo,
		log:        log.New().WithField("anyKey","anyValue"),
	}

	handler.Put(c)

	if c.Writer.Status() != 200 {
		t.Fatalf("Expected 200. Received %d", c.Writer.Status())
	}
}

func TestMarketHandler_Delete(t *testing.T) {
	mockRepo := &MarketRepositoryStub{
		MockRepo: MockIMarketRepository{
			Delete: func(id string) error {
				return nil
			},
		},
	}

	c, _ := getGinContextMock("anyId", "DELETE", nil, "", "")
	handler := MarketHandler{
		requestId:  "a_request_id",
		marketRepo: mockRepo,
		log:        log.New().WithField("anyKey","anyValue"),
	}

	handler.Delete(c)

	if c.Writer.Status() != 200 {
		t.Fatalf("Expected 200. Received %d", c.Writer.Status())
	}
}

func TestMarketHandler_DeleteError(t *testing.T) {
	mockRepo := &MarketRepositoryStub{
		MockRepo: MockIMarketRepository{
			Delete: func(id string) error {
				return errors.New("an_error")
			},
		},
	}

	c, _ := getGinContextMock("anyId", "DELETE", nil, "", "")
	handler := MarketHandler{
		requestId:  "a_request_id",
		marketRepo: mockRepo,
		log:        log.New().WithField("anyKey","anyValue"),
	}

	handler.Delete(c)

	if c.Writer.Status() != 500 {
		t.Fatalf("Expected 500. Received %d", c.Writer.Status())
	}
}

func TestMarketHandler_Get(t *testing.T) {
	mockRepo := &MarketRepositoryStub{
		MockRepo: MockIMarketRepository{
			GetItem: func(id string) (model.Record, error) {
				return utils.GetFakeRecord(), nil
			},
		},
	}

	c, _ := getGinContextMock("anyId", "GET", nil, "", "")
	handler := MarketHandler{
		requestId:  "a_request_id",
		marketRepo: mockRepo,
		log:        log.New().WithField("anyKey","anyValue"),
	}

	handler.Get(c)

	if c.Writer.Status() != 200 {
		t.Fatalf("Expected 200. Received %d", c.Writer.Status())
	}
}

func TestMarketHandler_Get404(t *testing.T) {
	mockRepo := &MarketRepositoryStub{
		MockRepo: MockIMarketRepository{
			GetItem: func(id string) (model.Record, error) {
				return model.Record{}, nil
			},
		},
	}

	c, _ := getGinContextMock("anyId", "GET", nil, "", "")
	handler := MarketHandler{
		requestId:  "a_request_id",
		marketRepo: mockRepo,
		log:        log.New().WithField("anyKey","anyValue"),
	}

	handler.Get(c)

	if c.Writer.Status() != 404 {
		t.Fatalf("Expected 404. Received %d", c.Writer.Status())
	}
}

func TestMarketHandler_GetByDistCodeEmpty(t *testing.T) {
	mockRepo := &MarketRepositoryStub{
		MockRepo: MockIMarketRepository{
			GetItemsByDistrictId: func(id string) ([]model.Record, error) {
				var records []model.Record
				return records, nil
			},
		},
	}

	c, w := getGinContextMock("", "GET", nil, "", "codDist")

	handler := MarketHandler{
		requestId:  "a_request_id",
		marketRepo: mockRepo,
		log:        log.New().WithField("anyKey","anyValue"),
	}

	handler.GetByDistCode(c)

	if c.Writer.Status() != 200 {
		t.Fatalf("Expected 200. Received %d", c.Writer.Status())
	}

	var got []model.Record
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	if len(got) != 0 {
		t.Fatalf("Expected empty array.")
	}
}

func TestMarketHandler_GetByDistCodeWithHits(t *testing.T) {
	mockRepo := &MarketRepositoryStub{
		MockRepo: MockIMarketRepository{
			GetItemsByDistrictId: func(id string) ([]model.Record, error) {
				var records []model.Record
				records = append(records, utils.GetFakeRecord(), utils.GetFakeRecord())
				return records, nil
			},
		},
	}

	c, w := getGinContextMock("", "GET", nil, "", "codDist")

	handler := MarketHandler{
		requestId:  "a_request_id",
		marketRepo: mockRepo,
		log:        log.New().WithField("anyKey","anyValue"),
	}

	handler.GetByDistCode(c)

	if c.Writer.Status() != 200 {
		t.Fatalf("Expected 200. Received %d", c.Writer.Status())
	}

	var got []model.Record
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	if len(got) != 2 {
		t.Fatalf("Expected 2 records.")
	}
}
