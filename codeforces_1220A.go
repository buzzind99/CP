//go:build codeforces_1220A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	_, s := nextInt(), next()
	nCount, zCount := strings.Count(s, "n"), strings.Count(s, "z")
	for range nCount {
		fmt.Fprint(wr, 1, " ")
	}
	for range zCount {
		fmt.Fprint(wr, 0, " ")
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1220/A
  Tags: implementation, sotrings, strings
  Rating: 800
*/
