package internal

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"virgo/internal/controller/user"
	"virgo/internal/store"
	"virgo/internal/store/mysql"
	"virgo/pkg/options"
)

func InitRouter()  {
	storeFactory := generateStoreFactory()
	router := gin.Default()

	router.GET("PING", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "PONG")
	})

	users := router.Group("users")
 	{
		 userController := user.NewUserController(storeFactory)
		 users.GET(":id", userController.GetUserById)
	}
	_ = router.Run("localhost:8080")
}

func generateStoreFactory() store.Factory {
	storeFactory, err := mysql.GetMySQLFactory(&options.MySQLOptions{
		Host: "localhost:3306",
		Username: "virgo",
		Password: "123456",
		Database: "virgo",
		LogLevel: 3,
	})
	if err != nil {
		log.Fatal(err)
	}
	return storeFactory
}