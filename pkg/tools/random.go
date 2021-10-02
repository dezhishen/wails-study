package tools

import (
	"math/rand"
	"time"
)

func RandomInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}
