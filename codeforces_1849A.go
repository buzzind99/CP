//go:build codeforces_1849A

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
		b, c, h := nextInt(), nextInt(), nextInt()

		fmt.Fprintln(wr, min(b-1, c+h)*2+1)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1849/A
  Tags: implementation, math
  Rating: 800
*/
