package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomInt(min, max int64) int64 {
	return min + rng.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rng.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomUsername() string {
	return RandomString(6)
}

func RandomAmount() int64 {
	return RandomInt(100, 1000)
}

func RandomEmail() string {
	return RandomString(6) + "@gmail.com"
}

func RandomPassword() string {
	return RandomString(8)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD", "RUB"}
	n := len(currencies)
	return currencies[rng.Intn(n)]
}
