package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/arturbaccarin/go-my-first-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {

	t.Run("user and password is valid", func(t *testing.T) {
		recorderCreateUser := httptest.NewRecorder()
		ctxCreateUser := GetTestGinContext(recorderCreateUser)

		recorderLoginUser := httptest.NewRecorder()
		ctxLoginUser := GetTestGinContext(recorderLoginUser)

		email := fmt.Sprintf("%d@test.com", rand.Int())
		password := "teste@#@123"

		userCreateRequest := request.UserRequest{
			Email:    email,
			Password: password,
			Name:     "Test User",
			Age:      10,
		}

		b, _ := json.Marshal(userCreateRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctxCreateUser, []gin.Param{}, url.Values{}, "POST", stringReader)
		UserController.CreateUser(ctxCreateUser)

		userRequest := request.UserLogin{
			Email:    email,
			Password: password,
		}

		b, _ = json.Marshal(userRequest)
		stringReader = io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctxLoginUser, []gin.Param{}, url.Values{}, "POST", stringReader)
		UserController.LoginUser(ctxLoginUser)

		assert.EqualValues(t, http.StatusOK, recorderLoginUser.Code)
		assert.NotEmpty(t, recorderLoginUser.Result().Header.Get("Authorization"))
	})

	t.Run("user and password is invalid", func(t *testing.T) {
		recorderCreateUser := httptest.NewRecorder()
		ctxCreateUser := GetTestGinContext(recorderCreateUser)

		recorderLoginUser := httptest.NewRecorder()
		ctxLoginUser := GetTestGinContext(recorderLoginUser)

		email := fmt.Sprintf("%d@test.com", rand.Int())
		password := "teste@#@123"

		userCreateRequest := request.UserRequest{
			Email:    email,
			Password: password,
			Name:     "Test User",
			Age:      10,
		}

		b, _ := json.Marshal(userCreateRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctxCreateUser, []gin.Param{}, url.Values{}, "POST", stringReader)
		UserController.CreateUser(ctxCreateUser)

		userRequest := request.UserLogin{
			Email:    email,
			Password: "teste123",
		}

		b, _ = json.Marshal(userRequest)
		stringReader = io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctxLoginUser, []gin.Param{}, url.Values{}, "POST", stringReader)
		UserController.LoginUser(ctxLoginUser)

		assert.EqualValues(t, http.StatusBadRequest, recorderLoginUser.Code)
		assert.Empty(t, recorderLoginUser.Result().Header.Get("Authorization"))
	})
}
