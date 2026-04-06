//go:build codeforces_1296A

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
		for range n {
			a := nextInt()
			if a%2 == 0 { even++ } else { odd++ }
		}

		if odd == 0 {
			fmt.Fprintln(wr, "NO")
		} else if even == 0 {
			if n%2 == 0 {
				fmt.Fprintln(wr, "NO")
			} else {
				fmt.Fprintln(wr, "YES")
			}
		} else {
			fmt.Fprintln(wr, "YES")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1296/A
  Tags: math
  Rating: 800
*/
