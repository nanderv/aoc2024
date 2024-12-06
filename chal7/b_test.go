package chal7_test

import (
	"github.com/nanderv/aoc2024/chal7"
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
	res := chal7.Bfunc(fl)
	assert.Equal(t, 6, res)
}
