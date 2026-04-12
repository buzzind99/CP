//go:build codeforces_1385A

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
		x, y, z := nextInt(), nextInt(), nextInt()
		xA, yA, zA := min(x, y), min(x, z), min(y, z)
		maxVal := max(x, y, z)
		maxValA := max(xA, yA, zA)

		if maxValA != maxVal {
			fmt.Fprintln(wr, "NO")
		} else {
			fmt.Fprintln(wr, "YES")
			fmt.Fprintln(wr, xA, yA, zA)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1385/A
  Tags: math
  Rating: 800
*/
