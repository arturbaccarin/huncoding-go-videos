package gingonic

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUserId(t *testing.T) {
	recorder := httptest.NewRecorder()

	context := GetTestGinContext(recorder)

	params := []gin.Param{{
		Key:   "id",
		Value: "1",
	}}

	u := url.Values{}
	u.Set("foo", "bar")

	MakeGet(context, params, u)

	GetUserId(context)

	assert.EqualValues(t, http.StatusOK, recorder.Code)
	got, _ := strconv.Atoi(recorder.Body.String())

	assert.Equal(t, 1, got)
}

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder) // cria um contexto mockado
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MakeGet(c *gin.Context, params gin.Params, u url.Values) {
	c.Request.Method = "GET"

	c.Request.Header.Set("Content-Type", "application/json")

	c.Params = params

	c.Request.URL.RawQuery = u.Encode()
}
