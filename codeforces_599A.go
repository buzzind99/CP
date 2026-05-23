//go:build codeforces_599A

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

	d1, d2, d3 := nextInt(), nextInt(), nextInt()
	path1 := d1+d2+d3
	path2 := 2*d1+2*d2
	path3 := 2*d1+2*d3
	path4 := 2*d2+2*d3
	ans := min(path1, min(path2, min(path3, path4)))

	fmt.Println(ans)
}

/*
  Link: https://codeforces.com/problemset/problem/599/A
  Tags: implementation
  Rating: 800
*/
