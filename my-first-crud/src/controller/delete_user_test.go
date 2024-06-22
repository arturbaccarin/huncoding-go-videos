package controller

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/arturbaccarin/go-my-first-crud/src/configuration/rest_err"
	"github.com/arturbaccarin/go-my-first-crud/src/test/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_DeleteUser(t *testing.T) {
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

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUserById(context)

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

		service.EXPECT().DeleteUser(id).Return(rest_err.NewInternalServerError("error test"))

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUserById(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("id is valid service returns error", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{{
			Key:   "id",
			Value: id},
		}

		service.EXPECT().DeleteUser(id).Return(nil)

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUserById(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}
