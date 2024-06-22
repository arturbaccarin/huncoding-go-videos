package controller

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"
	"github.com/arturbaccarin/go-my-first-crud/src/model"
	"github.com/arturbaccarin/go-my-first-crud/src/test/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

// mockgen -source=src/model/service/user_interface.go -destination=src/test/mock/user_interface_mock.go -package=mock

func TestUserController_FindUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock.NewMockUserDomainService(ctrl)

	controller := NewUserControllerInterface(service)

	t.Run("email is invalid returns error", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{{
			Key:   "email",
			Value: "TEST_ERROR"},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("email is invalid service returns error", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{{
			Key:   "email",
			Value: "test@test.com"},
		}

		service.EXPECT().FindUserByEmail("test@test.com").Return(nil, rest_err.NewInternalServerError("error test"))

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("email is valid service returns error", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{{
			Key:   "email",
			Value: "test@test.com"},
		}

		service.EXPECT().FindUserByEmail("test@test.com").Return(model.NewUserDomain("test@test.com", "test", "name", 98), nil)

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

func TestUserController_FindUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock.NewMockUserDomainService(ctrl)

	controller := NewUserControllerInterface(service)

	t.Run("id is invalid returns error", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{{
			Key:   "id",
			Value: "TEST_ERROR"},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserById(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("id is invalid service returns error", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{{
			Key:   "id",
			Value: id},
		}

		service.EXPECT().FindUserById(id).Return(nil, rest_err.NewInternalServerError("error test"))

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserById(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("id is valid service returns success", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{{
			Key:   "id",
			Value: id},
		}

		service.EXPECT().FindUserById(id).Return(model.NewUserDomain("test@test.com", "test", "name", 98), nil)

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserById(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {

	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MakeRequest(c *gin.Context, param gin.Params, u url.Values, method string, body io.ReadCloser) {

	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param

	c.Request.Body = body
	c.Request.URL.RawQuery = u.Encode()
}
