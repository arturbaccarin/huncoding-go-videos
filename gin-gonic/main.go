package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"test" binding:"required,min=1"`
	Age  int    `json:"testAge" binding:"required,min=18"`
}

// https://youtu.be/dd0Ot074VTE?list=PLm-xZWCprwYRAsLvf43sg5ZWuIvmmnDp9
func main() {

	// r := gin.New() // inicializando o gin no modo cru
	r := gin.Default() // já instancia logger e mais

	r.GET("/ping", MidJwt, func(context *gin.Context) { // aceita uma lista de handlers e se precisar abortar usa context.Abort()
		// /ping/:id -> c.Param("id")
		// /ping?id=123 -> c.Query("id")
		// {"name": "abc"} -> c.ShouldBindJSON(&user)
		user := User{}
		err := context.ShouldBindJSON(&user)
		if err != nil {
			return context.String(http.StatusBadRequest, err.Error())
		}

		context.JSON(http.StatusOK, user)
	}) // handler: a função que vai executar quando chegar sua requisição

	err := r.Run(":9090")
	if err != nil {
		return
	}
}

func MidJwt(context *gin.Context) {

	headerAuth := context.Request.Header.Get("Authorization")
	if !strings.HasPrefix(headerAuth, "Bearer ") {
		context.String(http.StatusBadRequest, "invalid token")
		context.Abort()
	}

}
