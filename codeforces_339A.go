//go:build codeforces_339A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"slices"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
    
	str := next()
	strSlice := strings.Split(str, "+")
	slices.Sort(strSlice)

	fmt.Println(strings.Join(strSlice, "+"))
}

/*
  Link: https://codeforces.com/problemset/problem/339/A
  Tags: greedy, implementation, sortings, strings
  Rating: 800
*/
