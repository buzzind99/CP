//go:build codeforces_1337A

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
		_, b, c, _ := nextInt(), nextInt(), nextInt(), nextInt()

		fmt.Fprintln(wr, b, c, c)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1337/A
  Tags: constructive algorithms, math
  Rating: 800
*/
