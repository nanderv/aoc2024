package chal21_test

import (
	"fmt"
	"github.com/nanderv/aoc2024/chal21"
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
	fmt.Println(chal21.GetBaseKeypad())
	chal21.Afunc(fl)
}
