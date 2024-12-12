package chal11_test

import (
	"github.com/nanderv/aoc2024/chal11"
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
	res := chal11.Afunc(fl)
	assert.Equal(t, 55312, res)
}
