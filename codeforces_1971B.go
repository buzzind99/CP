//go:build codeforces_1971B

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
		s := []byte(next())
		swapped := false
		for i := 1; i < len(s); i++ {
			if s[i] != s[0] {
				s[i], s[0] = s[0], s[i]
				swapped = true
				break
			}
		}

		if !swapped {
			fmt.Fprintln(wr, "NO")
		} else {
			fmt.Fprintln(wr, "YES")
			fmt.Fprintln(wr, string(s))
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1971/B
  Tags: implementation, strings
  Rating: 800
*/
