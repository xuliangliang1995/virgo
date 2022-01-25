package main

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"log"
	"time"
)

type VirgoUser struct {
	Id int64
	Name string
	Age int32
	Gender    int8
	CreateAt    time.Time
	UpdateAt   time.Time
	DeleteFlag int8
}

func main() {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: "virgo",
			Username: "",
			Password: "",
		},
		//Debug:           true,
		DialTimeout:     time.Second,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := clickhouse.Context(context.Background(),
		clickhouse.WithSettings(clickhouse.Settings{
			"max_block_size": 10,
		}),
		clickhouse.WithProgress(func(progress *clickhouse.Progress) {
			fmt.Println("progress : ", progress)
		}))

	err = conn.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := conn.Query(ctx, "SELECT * FROM virgo_user")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var user VirgoUser
		if err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Gender, &user.CreateAt, &user.UpdateAt, &user.DeleteFlag); err != nil {
			log.Fatal(err)
		}
		log.Printf("%v\n", user)
	}

	rows.Close()
}
