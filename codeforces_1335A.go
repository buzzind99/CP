//go:build codeforces_1335A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

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
		n := nextInt()
		if n <= 2 {
			fmt.Println(0)
			continue
		}
		if n%2 == 0 {
			fmt.Println((n/2)-1)
		} else {
			fmt.Println(n/2)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1335/A
  Tags: math
  Rating: 800
*/
