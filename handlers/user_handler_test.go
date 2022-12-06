package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"unittest/mocks"
	"unittest/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestCreateUserSuccess(t *testing.T) {
	testData := &models.CreateUserData{
		Email:    "trminh36@gmail.com",
		Name:     "Tran Minh",
		Password: "12345",
	}

	userQueriesMock := new(mocks.IUserQueries)
	userQueriesMock.On("CreateUser", &models.CreateUserData{Email: "trminh36@gmail.com",
		Name:     "Tran Minh",
		Password: "12345"}).Return(testData)

	handler := NewUserHandler(userQueriesMock)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	jsonbytes, _ := json.Marshal(testData)
	ctx.Request = &http.Request{
		Body: io.NopCloser(bytes.NewBuffer(jsonbytes)),
	}

	handler.CreateUser(ctx)

	userQueriesMock.AssertCalled(t, "CreateUser")
	require.Equal(t, http.StatusOK, ctx.Writer.Status())
}
