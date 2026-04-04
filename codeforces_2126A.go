//go:build codeforces_2126A

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
		x := next()
		minVal := 10
		for i := range x {
			minVal = min(minVal, int(x[i]-'0'))
			if x[i] == '0' { break }
		}

		fmt.Fprintln(wr, minVal)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2126/A
  Tags: brute force, implementation, math
  Rating: 800
*/
