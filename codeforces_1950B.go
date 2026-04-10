//go:build codeforces_1950B

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
		for i := range 2*n {
			row := make([]byte, 2*n)
			for j := range 2*n {
				if (i/2+j/2)%2 == 0 {
					row[j] = '#'
				} else {
					row[j] = '.'
				}
			}
			fmt.Fprintf(wr, "%s\n", row)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1950/B
  Tags: implementation
  Rating: 800
*/
