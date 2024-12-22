package chal22_test

import (
	"fmt"
	"github.com/nanderv/aoc2024/chal22"
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

	v := chal22.Afunc(fl)
	assert.Equal(t, 37327623, v)
}
func TestNxt(t *testing.T) {
	num := 123
	num = chal22.Nxt(num)
	fmt.Println(num)
	num = chal22.Nxt(num)
	fmt.Println(num)
	num = chal22.Nxt(num)
	fmt.Println(num)
	num = chal22.Nxt(num)
	fmt.Println(num)
	num = chal22.Nxt(num)
	fmt.Println(num)
	num = chal22.Nxt(num)
	fmt.Println(num)
	num = chal22.Nxt(num)
	fmt.Println(num)
	num = chal22.Nxt(num)
	fmt.Println(num)
	num = chal22.Nxt(num)
	fmt.Println(num)
	num = chal22.Nxt(num)
	fmt.Println(num)

}
