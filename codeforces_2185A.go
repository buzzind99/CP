//go:build codeforces_2185A

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
		for i := range n {
			fmt.Fprint(wr, i+1, " ")
		}

		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2185/A
  Tags: constructive algorithms, math
  Rating: 800
*/
