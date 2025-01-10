package util

import (
	"math/rand"
	"time"
)

var rng *rand.Rand

func init() {

	seed := time.Now().UnixNano()
	rng = rand.New(rand.NewSource(seed))

}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// this should be a random string between min and max
func RandomInt(min, max int64) int64 {
	return min + rng.Int63n(max-min+1)
}

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[rng.Intn(len(alphabet))]
	}
	return string(b)
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"NGN", "USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rng.Intn(n)]
}
