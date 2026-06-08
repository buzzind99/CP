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
		m := make(map[int]int)
		maxIdx := -1
		for i := range n {
			a := nextInt()
			_, exists := m[a]
			if !exists {
				m[a] = i
			} else {
				maxIdx = max(maxIdx, m[a])
				m[a] = i
			}
		}

		ans := 0
		if maxIdx >= 0 { ans = maxIdx+1 }

		fmt.Fprintln(wr, ans)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1714/B
  Tags: data structures, greedy, implementation
  Rating: 800
*/
