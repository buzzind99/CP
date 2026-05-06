//go:build codeforces_1996B

package main

import (
	"bufio"
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
		count := 0
		for range n {
			s := next()
			if count%k == 0 {
				for i := 0; i < n; i += k {
					wr.WriteByte(s[i])
				}
				wr.WriteByte('\n')
			}
			count++
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1996/B
  Tags: greedy, implementation
  Rating: 800
*/
