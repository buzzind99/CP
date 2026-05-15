//go:build codeforces_1997A

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

func isEven (n int) bool {
	return n%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		s := next()
		a := "a"
		if len(s) == 1 {
			if s[0] == byte("a"[0]) { a = "b" }
			fmt.Fprintf(wr, "%s%s\n", a, string(s[0]))
			continue
		}
		solved := false
		for i := range len(s)-1 {
			if s[i+1] == s[i] {
				if s[i] == byte("a"[0]) { a = "b" }
				fmt.Fprintf(wr, "%s%s%s\n", s[:i+1], a, s[i+1:])
				solved = true
				break
			}
		}

		if !solved {
			if s[len(s)-1] == byte("a"[0]) { a = "b" }
			fmt.Fprintf(wr, "%s%s\n", s, a)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1997/A
  Tags: brute force, implementation, strings
  Rating: 800
*/
