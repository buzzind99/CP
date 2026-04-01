//go:build codeforces_1837A

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

	t := nextInt()
	for range t {
		x, k := nextInt(), nextInt()
		sol := []int{}
		for i := x; x > 0; i-- {
			if i%k != 0 {
				x -= i
				sol = append(sol, i)
				i = x+1
			}
		}
		num := len(sol)
		fmt.Fprintln(wr, num)
		for _, v := range sol {
			fmt.Fprint(wr, v, " ")
		}
		fmt.Fprintln(wr)
	}

	defer wr.Flush()
}

/*
  Link: https://codeforces.com/problemset/problem/1837/A
  Tags: constructive algorithms, math
  Rating: 800
*/
