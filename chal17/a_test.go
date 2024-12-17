package chal17_test

import (
	"github.com/nanderv/aoc2024/chal17"
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
	res := chal17.Afunc(fl)
	assert.Equal(t, "4,6,3,5,6,3,5,2,1,0", res)
}
