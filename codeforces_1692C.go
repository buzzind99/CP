//go:build codeforces_1692C

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

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		ansX, ansY := 0, 0
		solved := false
		for j := range 8 {
			s := next()
			count := 0
			p1x, p2x := 0, 0

			if solved { continue }
			for i := range s {
				if s[i] == '#' {
					count++
					if count == 1 {
						p1x = i+1
					} else if count == 2 {
						p2x = i+1
						ansX = (p2x-p1x)/2+p1x
						ansY = j+ansX+1-p1x
						solved = true
					}
				}
			}
		}

		fmt.Fprintln(wr, ansY, ansX)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1692/C
  Tags: implementation
  Rating: 800
*/
