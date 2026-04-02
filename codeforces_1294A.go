//go:build codeforces_1294A

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
		a, b, c, n := nextInt(), nextInt(), nextInt(), nextInt()
		maxCoin := max(a, b, c)
		needToGive := 3*maxCoin-a-b-c
		if n < needToGive {
			fmt.Fprintln(wr, "NO")
		} else {
			if (n-needToGive)%3 != 0 {
				fmt.Fprintln(wr, "NO")
			} else {
				fmt.Fprintln(wr, "YES")
			}
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1294/A
  Tags: math
  Rating: 800
*/
