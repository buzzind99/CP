//go:build codeforces_1862B

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
		prev := 0
		sol := make([]int, 0, 2*n)
		for range n {
			a := nextInt()
			if a >= prev {
				sol = append(sol, a)
			} else {
				sol = append(sol, a)
				sol = append(sol, a)
			}
			prev = a
		}
		num := len(sol)
		fmt.Fprintln(wr, num)
		for _, v := range sol {
			fmt.Fprint(wr, v, " ")
		}
		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1862/B
  Tags: constructive algorithms
  Rating: 800
*/
