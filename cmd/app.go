package main

import (
	"math/rand"
	"time"
	"virgo/internal"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	internal.InitRouter()
}
