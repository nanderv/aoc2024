package chal10_test

import (
	"github.com/nanderv/aoc2024/chal10"
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
	res := chal10.Afunc(fl)
	assert.Equal(t, 36, res)
}

func TestChal1ASmol(t *testing.T) {
	fl, err := os.Open("smol.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal10.Afunc(fl)
	assert.Equal(t, 1, res)
}
