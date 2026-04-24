//go:build codeforces_894A

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

	s := next()
	m := make(map[byte][]int)
	for i := range s {
		if s[i] == 'Q' || s[i] == 'A' { m[s[i]] = append(m[s[i]], i) }
	}

	sum := 0
	for i := range len(m['Q']) {
		for j := range len(m['A']) {
			for k := range len(m['Q']) {
				if m['Q'][i] < m['A'][j] && m['A'][j] < m['Q'][k] { sum += len(m['Q'])-k; break }
			}
		}
	}

	fmt.Fprintln(wr, sum)
}

/*
  Link: https://codeforces.com/problemset/problem/894/A
  Tags: brute force, dp
  Rating: 800
*/
