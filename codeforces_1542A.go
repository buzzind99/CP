//go:build codeforces_1542A

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
		odd, even := 0, 0
		for range 2*n {
			a := nextInt()
			if a%2 == 0 { even++ } else { odd++ }
		}

		if odd == even {
			fmt.Fprintln(wr, "Yes")
		} else {
			fmt.Fprintln(wr, "No")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1542/A
  Tags: math
  Rating: 800
*/
