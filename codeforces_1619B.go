//go:build codeforces_1619B

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
		n := nextInt()
		liked := make(map[int]bool)
		for i := 1; i*i <= n; i++ {
			liked[i*i] = true
		}
		for i := 1; i*i*i <= n; i++ {
			liked[i*i*i] = true
		}
		fmt.Println(len(liked))
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1619/B
  Tags: implementation, math
  Rating: 800
*/
