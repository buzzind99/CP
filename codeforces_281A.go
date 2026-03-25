//go:build codeforces_281A

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
	sc.Split(bufio.ScanWords)
    
	str := next()

	fmt.Printf("%s%s\n", strings.ToUpper(string(str[0])), str[1:])
}

/*
  Link: https://codeforces.com/problemset/problem/281/A
  Tags: implementation, strings
  Rating: 800
*/
