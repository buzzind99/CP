//go:build codeforces_2214A

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

	fmt.Fprintln(wr, "C2")
}

/*
  Link: https://codeforces.com/problemset/problem/2214/A
  Tags: april fools
  Rating: -
*/
