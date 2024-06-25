package gingonic

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserId(c *gin.Context) {
	fmt.Println(c.Query("foo"))
	fmt.Println(c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, id)
}
