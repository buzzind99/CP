//go:build codeforces_59A

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
    
	s := next()
	upperCount := 0
	for _, char := range s {
		if unicode.IsUpper(char) { upperCount++ }
	}

	if upperCount > len(s)/2 {
		fmt.Println(strings.ToUpper(s))
	} else {
		fmt.Println(strings.ToLower(s))
	}
}

/*
  Link: https://codeforces.com/problemset/problem/59/A
  Tags: implementation, strings
  Rating: 800
*/
