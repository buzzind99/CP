//go:build codeforces_1541A

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
		n := nextInt()
		if n == 2 {
			fmt.Fprintln(wr, 2, 1)
		} else if n == 3 {
			fmt.Fprintln(wr, 3, 1, 2)
		} else {
			if isEven(n) {
				for i := 1; i <= n; i++ {
					if !isEven(i) {
						fmt.Fprint(wr, i+1, " ")
					} else {
						fmt.Fprint(wr, i-1, " ")
					}
				}
			fmt.Fprintln(wr)
			} else {
				fmt.Fprint(wr, 3, 1, 2, " ")
				for i := 4; i <= n; i++ {
					if isEven(i) {
						fmt.Fprint(wr, i+1, " ")
					} else {
						fmt.Fprint(wr, i-1, " ")
					}
				}
				fmt.Fprintln(wr)
			}
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1541/A
  Tags: constructive algorithms, greedy, implementation
  Rating: 800
*/
