package chal13_test

import (
	"github.com/nanderv/aoc2024/chal13"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestChaAs(t *testing.T) {
	fl, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal13.Bfunc(fl)
	assert.Equal(t, 368, res)
}
