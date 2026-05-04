//go:build codeforces_1919A

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

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		a, b := nextInt(), nextInt()

		if !isEven(a+b) {
			fmt.Fprintln(wr, "Alice")
		} else {
			fmt.Fprintln(wr, "Bob")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1919/A
  Tags: games, math
  Rating: 800
*/
