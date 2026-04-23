//go:build codeforces_1385B

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
		arr := make([]bool, n+1)
		for range 2*n {
			a := nextInt()
			if !arr[a] {
				fmt.Fprint(wr, a, " ")
				arr[a] = true
			}
		}

		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1385/B
  Tags: greedy
  Rating: 800
*/
