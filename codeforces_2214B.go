//go:build codeforces_2214B

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

	next()
	fmt.Fprintln(wr, "1F604")
}

/*
  Link: https://codeforces.com/problemset/problem/2214/B
  Tags: april fools
  Rating: -
  Contest: April Fools Day Contest 2026
*/
