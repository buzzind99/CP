//go:build codeforces_1593A

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
		a, b, c := nextInt(), nextInt(), nextInt()
		ansA := max(0, max(b, c)+1-a)
		ansB := max(0, max(a, c)+1-b)
		ansC := max(0, max(a, b)+1-c)

		fmt.Fprintln(wr, ansA, ansB, ansC)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1593/A
  Tags: math
  Rating: 800
*/
