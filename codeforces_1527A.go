//go:build codeforces_1527A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math/bits"
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
		n := nextInt()
		bitLen := bits.Len(uint(n))-1
		sol := powInt(2, bitLen)-1
		
		fmt.Fprintln(wr, sol)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1527/A
  Tags: bitmasks
  Rating: 800
*/
