package chal20_test

import (
	"github.com/nanderv/aoc2024/chal20"
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
	res := chal20.Algo(fl, 38, 2)
	assert.Equal(t, 3, res)
	fl, err = os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res = chal20.Algo(fl, 63, 2)
	assert.Equal(t, 1, res)
}
