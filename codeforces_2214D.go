//go:build codeforces_2214D

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

	switch n := nextInt(); {
	case n == 1:
		fmt.Fprintln(wr, "walk")
	case n == 2:
		fmt.Fprintln(wr, "no")
	case n == 3:
		fmt.Fprintln(wr, "no")
	case n == 4:
		fmt.Fprintln(wr, "no")
	case n == 5:
		fmt.Fprintln(wr, "yes")
	case n == 6:
		fmt.Fprintln(wr, "yes")
	case n == 7:
		fmt.Fprintln(wr, "backwards")
	case n == 8:
		fmt.Fprintln(wr, "7")
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2214/D
  Tags: april fools
  Rating: -
*/
