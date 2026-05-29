//go:build codeforces_1285A

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
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n, _ := nextInt(), next()

	fmt.Fprintln(wr, n+1)
}

/*
  Link: https://codeforces.com/problemset/problem/1285/A
  Tags: math
  Rating: 800
*/
