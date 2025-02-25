package main

import (
	"flag"
	"fmt"
	"github.com/nanderv/aoc2024/chal1"
	"github.com/nanderv/aoc2024/chal10"
	"github.com/nanderv/aoc2024/chal11"
	"github.com/nanderv/aoc2024/chal12"
	"github.com/nanderv/aoc2024/chal13"
	"github.com/nanderv/aoc2024/chal14"
	"github.com/nanderv/aoc2024/chal15"
	"github.com/nanderv/aoc2024/chal16"
	"github.com/nanderv/aoc2024/chal17"
	"github.com/nanderv/aoc2024/chal18"
	"github.com/nanderv/aoc2024/chal19"
	"github.com/nanderv/aoc2024/chal2"
	"github.com/nanderv/aoc2024/chal20"
	"github.com/nanderv/aoc2024/chal21"
	"github.com/nanderv/aoc2024/chal22"
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
	f    func(io.Reader) any
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
		{
			name: "13a",
			f:    chal13.Afunc,
			file: "chal13",
		},
		{
			name: "13b",
			f:    chal13.Bfunc,
			file: "chal13",
		},
		{
			name: "14a",
			f:    chal14.Afunc,
			file: "chal14",
		},
		{
			name: "14b",
			f:    chal14.Bfunc,
			file: "chal14",
		},
		{
			name: "15a",
			f:    chal15.Afunc,
			file: "chal15",
		},
		{
			name: "15b",
			f:    chal15.Bfunc,
			file: "chal15",
		},
		{
			name: "16a",
			f:    chal16.Afunc,
			file: "chal16",
		},
		{
			name: "16b",
			f:    chal16.Bfunc,
			file: "chal16",
		},
		{
			name: "17a",
			f:    chal17.Afunc,
			file: "chal17",
		},
		{
			name: "17b",
			f:    chal17.Bfunc,
			file: "chal17",
		},
		{
			name: "18a",
			f:    chal18.Afunc,
			file: "chal18",
		},
		{
			name: "18b",
			f:    chal18.Bfunc,
			file: "chal18",
		},
		{
			name: "19a",
			f:    chal19.Afunc,
			file: "chal19",
		},
		{
			name: "19b",
			f:    chal19.Bfunc,
			file: "chal19",
		},
		{
			name: "20a",
			f:    chal20.Afunc,
			file: "chal20",
		},
		{
			name: "20b",
			f:    chal20.Bfunc,
			file: "chal20",
		},
		{
			name: "21a",
			f:    chal21.Afunc,
			file: "chal21",
		},
		{
			name: "21b",
			f:    chal21.Bfunc,
			file: "chal21",
		},
		{
			name: "22a",
			f:    chal22.Afunc,
			file: "chal22",
		},
		{
			name: "22b",
			f:    chal22.Bfunc,
			file: "chal22",
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
