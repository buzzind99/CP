//go:build codeforces_1980A

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
		n, m := nextInt(), nextInt()
		s := next()
		arr := [7]int{}
		for i := range n {
			arr[int(s[i]-'A')]++
		}

		sum := 0
		for _, v := range arr {
			sum += m-min(v, m)
		}

		fmt.Fprintln(wr, sum)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1980/A
  Tags: math
  Rating: 800
*/
