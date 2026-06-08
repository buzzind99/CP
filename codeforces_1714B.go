//go:build codeforces_1714B

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
		arr := make([]int, n)
		for i := range n {
			arr[i] = nextInt()
		}

		seen := make([]bool, n+1)
		ans := 0
		for i := n - 1; i >= 0; i-- {
			val := arr[i]
			if seen[val] {
				ans = i + 1
				break
			}
			seen[val] = true
		}

		fmt.Fprintln(wr, ans)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1714/B
  Tags: data structures, greedy, implementation
  Rating: 800
*/
