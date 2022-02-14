package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetDefault("application.name", "virgo")
	viper.SetDefault("users", map[int64]string{1 : "Tom", 2 : "Jerry"})

	usersMap := viper.Get("users").(map[int64]string)

	for userId, username := range usersMap {
		fmt.Printf("%d : %s\n", userId, username)
	}

	fmt.Println("The properties read from cfg_001.yaml .")
	viper.SetConfigFile("cfg_001.yaml")
	viper.SetConfigName("cfg_001")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("examples/viper")
	var err error
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	keys := viper.AllKeys()
	for _, key := range keys {
		fmt.Printf("%s : %v\n", key, viper.Get(key))
	}

	viper.Set("mysql.username", "virgo_r")
	err = viper.WriteConfig()
	if err != nil {
		log.Fatal(err)
	}
	
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed : ", e.Name)
	})
	viper.WatchConfig()

	for ;;{}

}
