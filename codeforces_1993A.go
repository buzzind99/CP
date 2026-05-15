//go:build codeforces_1993A

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
		n, s := nextInt(), next()
		arr := [4]int{}
		for i := range s {
			if s[i] != '?' { arr[int(s[i]-'A')]++ }
		}

		sum := 0
		for _, v := range arr {
			inc := v
			if inc > n { inc = n }
			sum += inc
		}

		fmt.Fprintln(wr, sum)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1993/A
  Tags: greedy, implementation
  Rating: 800
*/
