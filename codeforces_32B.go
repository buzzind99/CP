//go:build codeforces_32B

package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
    
	s := next()
	var curr string
	var sol string
	for i := range s {
		curr += string(s[i])
		if curr == "." {
			sol += "0"
			curr = ""
		} else if curr == "-." {
			sol += "1"
			curr = ""
		} else if curr == "--" {
			sol += "2"
			curr = ""
		}
	}

	fmt.Println(sol)
}

/*
  Link: https://codeforces.com/problemset/problem/32/B
  Tags: expression parsing, implementation
  Rating: 800
*/
