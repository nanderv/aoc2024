package main

import (
	"flag"
	"fmt"
	"github.com/nanderv/aoc2024/chal1"
	"github.com/nanderv/aoc2024/chal2"
	"github.com/nanderv/aoc2024/chal3"
	"github.com/nanderv/aoc2024/chal4"
	"github.com/nanderv/aoc2024/chal5"
	"github.com/nanderv/aoc2024/chal6"
	"log"
	"os"
)

const fileName = "/input.txt"

func main() {
	var challenge string
	flag.StringVar(&challenge, "challenge", "last", "")
	flag.Parse()
	if challenge != "last" {
		doChallenge(challenge)
	}
	for _, challenge := range []string{"6a", "6b"} {
		doChallenge(challenge)
	}
}

func doChallenge(challenge string) {
	f := chal1.Afunc
	file := "chal1"
	switch challenge {
	case "1a":
		f = chal1.Afunc
		file = "chal1"
	case "1b":
		f = chal1.Bfunc
		file = "chal1"
	case "2a":
		f = chal2.Afunc
		file = "chal2"
	case "2b":
		f = chal2.Bfunc
		file = "chal2"
	case "3a":
		f = chal3.Afunc
		file = "chal3"
	case "3b":
		f = chal3.Bfunc
		file = "chal3"
	case "4a":
		f = chal4.Afunc
		file = "chal4"
	case "4b":
		f = chal4.Bfunc
		file = "chal4"
	case "5a":
		f = chal5.Afunc
		file = "chal5"
	case "5b":
		f = chal5.Bfunc
		file = "chal5"
	case "6a":
		f = chal6.Afunc
		file = "chal6"
	case "6b":
		f = chal6.Bfunc
		file = "chal6"
	}

	fmt.Println(file + fileName)
	fl, err := os.Open(file + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	res := f(fl)
	fmt.Println(challenge, res)
}
