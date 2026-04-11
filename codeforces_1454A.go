//go:build codeforces_1454A

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

	n := nextInt()
	for range n {
		n := nextInt()
		for i := range n {
			fmt.Fprint(wr, (i+1)%n+1, " ")
		}
		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1454/A
  Tags: constructive algorithms, probabilities
  Rating: 800
*/
