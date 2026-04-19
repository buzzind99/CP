//go:build codeforces_978B

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

	nextInt()
	s := next()
	count, sol := 0, 0
	for i := range s {
		if s[i] == 'x' {
			count++
		} else {
			if count > 2 { sol += count-2 }
			count = 0
		}
	}
	if count > 2 { sol += count-2 }

	fmt.Fprintln(wr, sol)
}

/*
  Link: https://codeforces.com/problemset/problem/978/B
  Tags: greedy, strings
  Rating: 800
*/
