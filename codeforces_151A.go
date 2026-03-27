//go:build codeforces_151A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

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

	n, k, l, c := nextInt(), nextInt(), nextInt(), nextInt()
	d, p, nl, np := nextInt(), nextInt(), nextInt(), nextInt()
	total_l := k*l
	total_d := c*d
	max_l := total_l/nl
	max_p := p/np
	sol := min(total_d, max_l, max_p)/n

	fmt.Println(sol)
}

/*
  Link: https://codeforces.com/problemset/problem/151/A
  Tags: implementation, math
  Rating: 800
*/
