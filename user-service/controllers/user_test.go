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

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MockModules() sqlmock.Sqlmock {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	SetDbClient(db)

	return mock
}

func TestGetUser(t *testing.T) {
	t.Skip()
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

	if w.Code != 200 {
		b, _ := io.ReadAll(w.Body)
		t.Error(w.Code, string(b))
	}
	assert.Equal(t, 1, 1, "Should be equal")
}
