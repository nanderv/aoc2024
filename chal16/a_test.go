package chal16_test

import (
	"github.com/nanderv/aoc2024/chal16"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestChalExample(t *testing.T) {
	fl, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal16.Afunc(fl)
	assert.Equal(t, 7036, res)
}

func TestChalLarger(t *testing.T) {
	fl, err := os.Open("larger.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal16.Afunc(fl)
	assert.Equal(t, 11048, res)
}
