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
	defer wr.Flush()

	t := nextInt()
	for range t {
		x, k := nextInt(), nextInt()
		sol := []int{}
		for x > 0 {
			i := x
			for i%k == 0 { i-- }
			sol = append(sol, i)
			x -= i
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
  Link: https://codeforces.com/problemset/problem/1837/A
  Tags: constructive algorithms, math
  Rating: 800
*/
