//go:build !solution

package main

import (
	"fmt"
	"os"

	// "maps"
	s "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var command_args []string = os.Args[1:]

	m := make(map[string]int)

	// fmt.Println(
	// 	command_args,
	// )

	for _, file := range command_args {
		dat, err := os.ReadFile(string(file))
		check(err)

		var dat_split []string = s.Split(string(dat), "\n")

		for _, sym := range dat_split {
			if _, ok := m[sym]; !ok {
				m[sym] = 0
			}
			m[sym] += 1
		}
	}

	for k, v := range m {
		if v < 2 {
			continue
		}
		fmt.Printf("%v\t%v\n", v, k)
	}
}
