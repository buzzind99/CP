//go:build codeforces_1925A

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
		for range n {
			for i := range k {
				wr.WriteByte(byte('a'+i))
			}
		}
		wr.WriteByte('\n')
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1925/A
  Tags: constructive algorithms, greedy, strings
  Rating: 800
*/
