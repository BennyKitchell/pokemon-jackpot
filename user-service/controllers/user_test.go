package controllers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"user-service/pkg/models"

	"github.com/DATA-DOG/go-sqlmock"
	miniredis "github.com/alicebob/miniredis/v2"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MockModules() sqlmock.Sqlmock {
	// mock db
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	SetDbClient(db)

	// mock kafka
	kafkaProducer, _ := kafka.NewProducer(&kafka.ConfigMap{"test.mock.num.brokers": 3})
	SetUserTopicProducer(kafkaProducer)

	// mock redis
	server, _ := miniredis.Run()

	SetRedis(redis.NewClient(&redis.Options{
		Addr: server.Addr(),
	}))
	return mock
}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mock := MockModules()
	mock2 := MockModules()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	content := models.User{
		Email:    "testaccount@aol.com",
		Password: "password",
		Username: "testaccount",
	}

	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"users\" WHERE email = $1 AND \"users\".\"id\" = $2 ORDER BY \"users\".\"id\" LIMIT $3")).WillReturnError(sql.ErrNoRows)
	mock.ExpectClose()
	mock2.ExpectBegin()
	mock2.ExpectQuery(regexp.QuoteMeta("INSERT INTO \"users\" (\"username\",\"email\",\"password\") VALUES ($1,$2,$3) RETURNING \"id\"")).WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "password"}).AddRow(1, "testaccount", "testaccount@aol.com", "password"))
	mock2.ExpectCommit()
	mock2.ExpectClose()

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
	CreateUser(c)

	mockUserResp := &models.User{
		ID:       1,
		Email:    "testaccount@aol.com",
		Password: "password",
		Username: "testaccount",
	}

	respBody, _ := json.Marshal(gin.H{
		"user": mockUserResp,
	})

	assert.Equal(t, respBody, w.Body.Bytes(), "They should be equal")
	if w.Code != 200 {
		b, _ := io.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}
}

func TestGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mock := MockModules()

	users := []models.User{
		{ID: 1, Username: "test1", Email: "test1@aol.com", Password: "password"},
		{ID: 2, Username: "test2", Email: "test2@aol.com", Password: "password"},
	}

	rows := mock.NewRows([]string{"id", "username", "email", "password"})
	for _, b := range users {
		rows.AddRow(b.ID, b.Username, b.Email, b.Password)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"users\" WHERE email = $1 AND password = $2 ORDER BY \"users\".\"id\" LIMIT $3")).
		WillReturnRows(rows)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")
	content := models.User{
		Email:    "test1@aol.com",
		Password: "password",
	}
	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
	GetUser(c)

	mockUserResp := &models.User{
		ID:       1,
		Email:    "test1@aol.com",
		Username: "test1",
		Password: "password",
	}
	respBody, err := json.Marshal(gin.H{
		"user": mockUserResp,
	})
	if err != nil {
		t.Error("Could not marshal json")
	}

	assert.Equal(t, respBody, w.Body.Bytes(), "They should be equal")
	if w.Code != 200 {
		b, _ := io.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}
}
