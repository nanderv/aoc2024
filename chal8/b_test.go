package chal8_test

import (
	"github.com/nanderv/aoc2024/chal8"
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
	res := chal8.Bfunc(fl)
	assert.Equal(t, 34, res)

}

func TestChal1y(t *testing.T) {
	fl, err := os.Open("smol.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal8.Bfunc(fl)
	assert.Equal(t, 5, res)
}

func TestChal1z(t *testing.T) {
	fl, err := os.Open("mid.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal8.Bfunc(fl)
	assert.Equal(t, 9, res)
}
