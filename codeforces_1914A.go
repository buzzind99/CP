//go:build codeforces_1914A

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
		n, s := nextInt(), next()
		counts := [26]int{}
		for i := range n {
			p := int(s[i]-'A')
			counts[p]++
		}

		ans := 0
		for i := range 26 {
			if counts[i] >= i+1 { ans++ }
		}

		fmt.Fprintln(wr, ans)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1914/A
  Tags: implementation, strings
  Rating: 800
*/
