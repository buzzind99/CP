//go:build codeforces_2065B

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
		s := next()
		same := false
		for i := range len(s)-1 {
			if s[i] == s[i+1] { same = true; break }
		}

		if same {
			fmt.Fprintln(wr, 1)
		} else {
			fmt.Fprintln(wr, len(s))
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2065/B
  Tags: strings
  Rating: 800
*/
