package main

import (
	"flag"
	"fmt"
	"github.com/nanderv/aoc2024/chal1"
	"github.com/nanderv/aoc2024/chal10"
	"github.com/nanderv/aoc2024/chal11"
	"github.com/nanderv/aoc2024/chal12"
	"github.com/nanderv/aoc2024/chal2"
	"github.com/nanderv/aoc2024/chal3"
	"github.com/nanderv/aoc2024/chal4"
	"github.com/nanderv/aoc2024/chal5"
	"github.com/nanderv/aoc2024/chal6"
	"github.com/nanderv/aoc2024/chal7"
	"github.com/nanderv/aoc2024/chal8"
	"github.com/nanderv/aoc2024/chal9"
	"io"
	"log"
	"os"
	"time"
)

const fileName = "/input.txt"

type chal struct {
	f    func(io.Reader) int
	file string
	name string
}

func main() {
	challenges := []chal{
		{
			name: "1a",
			f:    chal1.Afunc,
			file: "chal1",
		},
		{
			name: "1b",
			f:    chal1.Bfunc,
			file: "chal1",
		},
		{
			name: "2a",
			f:    chal2.Afunc,
			file: "chal2",
		},
		{
			name: "2b",
			f:    chal2.Bfunc,
			file: "chal2",
		},
		{
			name: "3a",
			f:    chal3.Afunc,
			file: "chal3",
		},
		{
			name: "3b",
			f:    chal3.Bfunc,
			file: "chal3",
		},
		{
			name: "4a",
			f:    chal3.Afunc,
			file: "chal3",
		},
		{
			name: "4b",
			f:    chal4.Bfunc,
			file: "chal4",
		},
		{
			name: "5a",
			f:    chal5.Afunc,
			file: "chal5",
		},
		{
			name: "5b",
			f:    chal5.Bfunc,
			file: "chal5",
		},
		{
			name: "6a",
			f:    chal6.Afunc,
			file: "chal6",
		},
		{
			name: "6b",
			f:    chal6.Bfunc,
			file: "chal6",
		},
		{
			name: "7a",
			f:    chal7.Afunc,
			file: "chal7",
		},
		{
			name: "7b",
			f:    chal7.Bfunc,
			file: "chal7",
		},
		{
			name: "8a",
			f:    chal8.Afunc,
			file: "chal8",
		},
		{
			name: "8b",
			f:    chal8.Bfunc,
			file: "chal8",
		},
		{
			name: "9a",
			f:    chal9.Afunc,
			file: "chal9",
		},
		{
			name: "9b",
			f:    chal9.Bfunc,
			file: "chal9",
		},
		{
			name: "10a",
			f:    chal10.Afunc,
			file: "chal10",
		},
		{
			name: "10b",
			f:    chal10.Bfunc,
			file: "chal10",
		},
		{
			name: "11a",
			f:    chal11.Afunc,
			file: "chal11",
		},
		{
			name: "11b",
			f:    chal11.Bfunc,
			file: "chal11",
		},
		{
			name: "12a",
			f:    chal12.Afunc,
			file: "chal12",
		},
		{
			name: "12b",
			f:    chal12.Bfunc,
			file: "chal12",
		},
	}
	var challenge string
	flag.StringVar(&challenge, "challenge", "last", "")
	flag.Parse()
	if challenge != "last" {
		for _, cc := range challenges {
			if cc.name == challenge {
				doChallenge(cc)
				return
			}
		}
	}
	for _, challenge := range challenges {
		doChallenge(challenge)
	}
}

func doChallenge(c chal) {
	fl, err := os.Open(c.file + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	tm := time.Now()
	res := c.f(fl)

	fmt.Println(c.name, res, time.Now().Sub(tm))
}
