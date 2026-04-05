//go:build codeforces_233A

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

	a := nextInt()
	if a%2 == 1 {
		fmt.Fprint(wr, -1)
		return
	}

	for i := a; i >= 1; i-- {
		fmt.Fprint(wr, i, " ")
	}
}

/*
  Link: https://codeforces.com/problemset/problem/233/A
  Tags: implementation, math
  Rating: 800
*/
