//go:build codeforces_443A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	s := next()
	if s == "{}" {
        fmt.Println(0)
        return
    }

	slice := strings.Split(s[1:len(s)-1], ", ")
	set := make(map[string]struct{})
	for _, char := range slice {
		set[char] = struct{}{}
	}

	fmt.Println(len(set))
}

/*
  Link: https://codeforces.com/problemset/problem/443/A
  Tags: constructive algorithms, implementation
  Rating: 800
*/
