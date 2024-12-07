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
	assert.Equal(t, 11387, res)

}

func TestTryop(t *testing.T) {

	assert.True(t, chal7.RecurseOp2(15, []int{6}, 156))
}
