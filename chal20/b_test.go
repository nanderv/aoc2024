package chal20_test

import (
	"github.com/nanderv/aoc2024/chal20"
	"github.com/nanderv/aoc2024/common"
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
	res := chal20.Algo(fl, 76, 20)
	assert.Equal(t, 3, res)
	fl, err = os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res = chal20.Algo(fl, 68, 20)
	assert.Equal(t, 55, res)
}
func TestGetNbs(t *testing.T) {
	test6 := chal20.GetNbs(6)
	assert.Contains(t, test6, common.Pos{Y: 4, X: 2})

}
