//go:build codeforces_1766A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		s := next()
		nDigits := len(s)
		firstDigit, _ := strconv.Atoi(string(s[0]))
		ans := (nDigits-1)*9 + firstDigit
		
		fmt.Fprintln(wr, ans)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1766/A
  Tags: brute force, implementation
  Rating: 800
*/
