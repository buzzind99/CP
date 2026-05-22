//go:build codeforces_1567A

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
		n, s := nextInt(), []byte(next())
		for i := range n {
			char := s[i]
			if char == 'U' { s[i] = 'D' } else if char == 'D' { s[i] = 'U' }
		}

		fmt.Fprintf(wr, "%s\n", s)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1567/A
  Tags: implementation, strings
  Rating: 800
*/
