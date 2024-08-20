package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Generates a random intereget between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Generates a random string of length of n (int)
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generate random owner name
func RandomOwner() string {
	return RandomString(6)
}

// Generate random amount of money
func RandomMoneyAmount() int64 {
	return RandomInt(0, 1000)
}

// Generate random currency code
func RandomCurrency() string {
	currencies := []string{EUR, USD, CAD}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}

// Generate random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
