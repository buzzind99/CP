//go:build codeforces_2214J

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
	fmt.Fprintln(wr, "Yes, I can attest to my status as a thoroughly validated, carbon-based biological entity.")
}

/*
  Link: https://codeforces.com/problemset/problem/2214/J
  Tags: april fools
  Rating: -
*/
