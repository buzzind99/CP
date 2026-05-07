//go:build codeforces_2195A

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
		one, sixSeven := false, false
		for range n {
			a := nextInt()
			if a == 1 { one = true } else if a == 67 { sixSeven = true }
		}

		if (one && sixSeven) || sixSeven {
			fmt.Fprintln(wr, "YES")
		} else {
			fmt.Fprintln(wr, "NO")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2195/A
  Tags: math, number theory
  Rating: 800
*/
