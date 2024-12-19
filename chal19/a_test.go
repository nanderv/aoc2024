package chal19_test

import (
	"github.com/nanderv/aoc2024/chal19"
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
	res := chal19.Afunc(fl)
	assert.Equal(t, 6, res)
}
