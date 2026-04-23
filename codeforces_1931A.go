//go:build codeforces_1931A

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
		c3 := min(n-2, 26)
		n -= c3
		c2 := min(n-1, 26)
		c1 := n - c2
		fmt.Fprintf(wr, "%c%c%c\n", 'a'+c1-1, 'a'+c2-1, 'a'+c3-1)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1931/A
  Tags: brute force, strings
  Rating: 800
*/
