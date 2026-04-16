//go:build codeforces_1907A

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
		s := next()
		for i := range 8 {
			if i+1 == int(s[1]-'0') { continue }
			fmt.Fprintf(wr, "%s%d\n", string(s[0]), i+1)
		}
		r := "abcdefgh"
		for i := range 8 {
			if r[i] == s[0] { continue }
			fmt.Fprintf(wr, "%s%s\n", string(r[i]), string(s[1]))
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1907/A
  Tags: implementation
  Rating: 800
*/
