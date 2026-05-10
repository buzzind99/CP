//go:build codeforces_2123B

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
		n, j, k := nextInt(), nextInt(), nextInt()
		playerStr, maxStr := 0, -1
		for i := range n {
			a := nextInt()
			maxStr = max(maxStr, a)
			if i == j-1 { playerStr = a}
		}

		if k >= 2 {
			fmt.Fprintln(wr, "YES")
		} else {
			if playerStr == maxStr {
				fmt.Fprintln(wr, "YES")
			} else {
				fmt.Fprintln(wr, "NO")
			}
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2123/B
  Tags: greedy
  Rating: 800
*/
