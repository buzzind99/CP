//go:build codeforces_1950C

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
		hour, _ := strconv.Atoi(s[:2])
		minute, _ := strconv.Atoi(s[3:])

		h := hour%12
		period := "AM"
		if h == 0 { h = 12 }
		if hour >= 12 { period = "PM" }

		fmt.Fprintf(wr, "%02d:%02d %s\n", h, minute, period)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1950/C
  Tags: implementation, math
  Rating: 800
*/
