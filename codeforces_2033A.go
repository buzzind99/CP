//go:build codeforces_2033A

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

func abs(x int) int {
	if x < 0 { return -x }
	return x
}

func isEven(x int) bool {
	return x%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()

		if n % 2 == 0 {
			fmt.Fprintln(wr, "Sakurako")
		} else {
			fmt.Fprintln(wr, "Kosuke")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2033/A
  Tags: constructive algorithms, implementation, math
  Rating: 800
*/
