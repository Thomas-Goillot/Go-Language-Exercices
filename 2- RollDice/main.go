package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("|", rand.Intn(6)+1, "|        |", rand.Intn(20)+1, "|        |", rand.Intn(100)+1, "|")
}
