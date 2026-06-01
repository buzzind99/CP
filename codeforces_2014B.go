//go:build codeforces_2014B

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
		n, k := nextInt(), nextInt()
		oddsToN := (n+1)/2
		oddsToStart := (n-k+1)/2
		oddCount := oddsToN-oddsToStart

		if oddCount % 2 == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2014/B
  Tags: math
  Rating: 800
*/
