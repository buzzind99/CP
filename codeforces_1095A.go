//go:build codeforces_1095A

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

	n, s := nextInt(), next()
	arr := make([]byte, 0)
	currLen := 1
	for i := 0; i < n; {
		arr = append(arr, s[i])
		currLen++
		i += currLen
	}

	fmt.Fprintln(wr, string(arr))
}

/*
  Link: https://codeforces.com/problemset/problem/1095/A
  Tags: implementation
  Rating: 800
*/
