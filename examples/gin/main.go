package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type User struct {
	Id int64	`json:"id"`
	Name string	`json:"name"`
	Age int32	`json:"age"`
}

func main()  {
	r := gin.Default()

	r.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	users := r.Group("/users")
	{
		users.GET("/:id", func(context *gin.Context) {
			id, _ := strconv.ParseInt(context.Param("id"), 10, 64)
			context.JSON(http.StatusOK, &User{
				Id: id,
				Name: "jerry",
				Age: 18,
			})
		})
	}

	_ = r.Run()
}
