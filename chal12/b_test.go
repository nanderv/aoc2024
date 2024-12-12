package chal12_test

import (
	"github.com/nanderv/aoc2024/chal12"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestChaAs(t *testing.T) {
	fl, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal12.Bfunc(fl)
	assert.Equal(t, 368, res)
}

func TestChaBSmol(t *testing.T) {
	fl, err := os.Open("smol.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := chal12.Bfunc(fl)
	assert.Equal(t, 80, res)
}
