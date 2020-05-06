package OptionalPattern

import (
	"log"
	"testing"
)

func TestNewHuman(t *testing.T) {
	h := NewHuman(WithName("QQQ"))
	log.Println(h)
}
