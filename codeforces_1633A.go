//go:build codeforces_1633A

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
		if n%7 == 0 {
			fmt.Fprintln(wr, n)
			continue
		}

		base := n-(n%10)
		for i := 0; i <= 9; i++ {
			if (base+i)%7 == 0 {
				fmt.Fprintln(wr, base+i)
				break
			}
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1633/A
  Tags: brute force
  Rating: 800
*/
