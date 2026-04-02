//go:build codeforces_1760B

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
		s := next()
		maxVal := -1
		for i := range n {
			maxVal = max(maxVal, int(s[i]-'a')+1)
		}
		
		fmt.Fprintln(wr, maxVal)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1760/B
  Tags: greedy, implementation, strings
  Rating: 800
*/
