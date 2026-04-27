//go:build codeforces_1702A

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

func powInt(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp&1 == 1 {
			result *= base
		}
		base *= base
		exp >>= 1
	}
	return result
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		m := nextInt()
		curr, pow := m, 0
		for curr > 0 {
			curr /= 10
			if curr > 0 { pow++ }
		}

		fmt.Fprintln(wr, m-powInt(10, pow))
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1702/A
  Tags: constructive algorithms
  Rating: 800
*/
