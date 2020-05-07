package profile

import (
	"math/big"
	"testing"
)

func TestProfile(t *testing.T) {
	i := big.NewInt(1000)
	BigIntFactorial(*i)
}