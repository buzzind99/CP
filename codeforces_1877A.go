//go:build codeforces_1877A

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
		n := nextInt()
		sum := 0
		for range n-1 {
			sum += nextInt()
		}
		fmt.Fprintln(wr, -sum)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1877/A
  Tags: math
  Rating: 800
*/
