//go:build !solution

package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var command_args []string = os.Args[1:]

	for _, url := range command_args {
		resp, err := http.Get(url)
		if err != nil {
			// defer resp.Body.Close()
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(resp.Body)
		check(err)

		_, _ = os.Stdout.Write(body)
	}
}
