package chal21_test

import (
	"github.com/nanderv/aoc2024/chal21"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestChal1A(t *testing.T) {
	fl, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()

	v := chal21.Afunc(fl)
	assert.Equal(t, 126384, v)
}
