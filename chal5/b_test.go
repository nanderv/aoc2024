package chal5_test

import (
	"github.com/nanderv/aoc2024/chal5"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestChal1B(t *testing.T) {
	fl, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal5.Bfunc(fl)
	assert.Equal(t, 123, res)
}
