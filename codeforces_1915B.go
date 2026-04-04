//go:build codeforces_1915B

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
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		for range 3 {
			s := next()
			if strings.Contains(s, "?") {
				var res byte = 64
				for i := 0; i < 3; i++ {
					if s[i] != '?' {
						res ^= s[i]
					}
				}

				fmt.Fprintln(wr, string(res))
			}
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1915/B
  Tags: bitmasks, brute force, implementation
  Rating: 800
*/
