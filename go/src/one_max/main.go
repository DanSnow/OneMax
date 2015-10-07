package main

import (
	"./evoluation"
	"./pool"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	pool := new(pool.Pool)
	evoluation.Evoluation(pool)
	fmt.Println(pool)
}
