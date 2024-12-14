package chal14_test

import (
	"github.com/nanderv/aoc2024/chal14"
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
	res := chal14.Afunc(fl)
	assert.Equal(t, 12, res)
}
