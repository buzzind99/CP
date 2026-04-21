//go:build codeforces_1607A

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

func abs(x int) int {
	if x < 0 { return -x }
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		r, s := next(), next()
		keyboard := make([]int, 26)
		for i := range r {
			keyboard[r[i]-'a'] = i
		}
		last := keyboard[s[0]-'a']
		sum := 0
		for i := 1; i < len(s); i++ {
			curr := keyboard[s[i]-'a']
			sum += abs(curr-last)
			last = curr
		}

		fmt.Fprintln(wr, sum)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1607/A
  Tags: implementation, strings
  Rating: 800
*/
