//go:build codeforces_1714C

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
		s := nextInt()
		ans := 0
		multiplier := 1
		currentDigit := 9
		for s > 0 {
			take := currentDigit
			if s < currentDigit { take = s }
			ans += take*multiplier
			s -= take
			multiplier *= 10
			currentDigit--
		}

		fmt.Fprintln(wr, ans)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1714/C
  Tags: greedy
  Rating: 800
*/
