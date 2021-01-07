package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 10

	fmt.Println(rand.Intn(max-min+1) + min)
}
