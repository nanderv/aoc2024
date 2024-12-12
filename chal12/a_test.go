package chal12_test

import (
	"github.com/nanderv/aoc2024/chal12"
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
	res := chal12.Afunc(fl)
	assert.Equal(t, 1930, res)
}

func TestChal1As(t *testing.T) {
	fl, err := os.Open("smol.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal12.Afunc(fl)
	assert.Equal(t, 140, res)
}

func TestChal1Aoxo(t *testing.T) {
	fl, err := os.Open("oxo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal12.Afunc(fl)
	assert.Equal(t, 772, res)
}
