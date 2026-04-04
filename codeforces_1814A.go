//go:build codeforces_1814A

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

func nextInt64() int64 {
	i, _ := strconv.ParseInt(next(), 10, 64)
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n, k := nextInt64(), nextInt64()
		if n%2 == 0 {
			fmt.Fprintln(wr, "YES")
		} else {
			if k%2 == 0 {
				fmt.Fprintln(wr, "NO")
			} else {
				fmt.Fprintln(wr, "YES")
			}
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1814/A
  Tags: implementation, math
  Rating: 800
*/
