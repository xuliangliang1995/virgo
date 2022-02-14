package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
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
	viper.SetConfigName("virgo_cfg")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs/mysql")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	keys := viper.AllKeys()
	for _, key := range keys {
		fmt.Printf("%s : %s\n", key, viper.Get(key))
	}

	var mysqlOptions options.MySQLOptions
	_ = mapstructure.WeakDecode(viper.Get("mysql"), &mysqlOptions)

	fmt.Println(mysqlOptions)

	storeFactory, err := mysql.GetMySQLFactory(&mysqlOptions)
	if err != nil {
		log.Fatal(err)
	}
	return storeFactory
}