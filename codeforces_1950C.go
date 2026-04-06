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

		if hour >= 12 {
			if hour == 12 {
				fmt.Fprintf(wr, "%02d:%02d PM\n", 12, minute)
			} else {
				fmt.Fprintf(wr, "%02d:%02d PM\n", hour%12, minute)
			}
		} else {
			if hour == 0 {
				fmt.Fprintf(wr, "%02d:%02d AM\n", 12, minute)
			} else {
				fmt.Fprintf(wr, "%02d:%02d AM\n", hour%12, minute)
			}
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1950/C
  Tags: implementation, math
  Rating: 800
*/
