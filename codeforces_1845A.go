//go:build codeforces_1845A

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
		n, k, x := nextInt(), nextInt(), nextInt()

		if x != 1 {
			fmt.Fprintln(wr, "YES")
			fmt.Fprintln(wr, n)
			for range n {
				fmt.Fprint(wr, 1, " ")
			}
			fmt.Fprintln(wr)
		} else {
			if n%2 == 0 && k >= 2 {
				fmt.Fprintln(wr, "YES")
				fmt.Fprintln(wr, n/2)
				for range n/2 {
					fmt.Fprint(wr, 2, " ")
				}
				fmt.Fprintln(wr)
			} else if n%2 != 0 && k >= 3 && n > 1 {
				fmt.Fprintln(wr, "YES")
				fmt.Fprintln(wr, (n-3)/2+1)
				fmt.Fprint(wr, 3, " ")
				for range (n-3)/2 {
					fmt.Fprint(wr, 2, " ")
				}
				fmt.Fprintln(wr)
			} else {
				fmt.Fprintln(wr, "NO")
			}
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1845/A
  Tags: constructive algorithms, implementation, math, number theory
  Rating: 800
*/
