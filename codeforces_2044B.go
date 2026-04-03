//go:build codeforces_2044B

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
		a := next()
		for i := len(a)-1; i >= 0; i-- {
			if a[i] == 'p' {
				fmt.Fprint(wr, "q")
			} else if a[i] == 'q' {
				fmt.Fprint(wr, "p")
			} else {
				fmt.Fprint(wr, string(a[i]))
			}
		}

		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2044/B
  Tags: implementation, strings
  Rating: 800
*/
