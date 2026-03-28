//go:build codeforces_1915A

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
        a, b, c := nextInt(), nextInt(), nextInt()
        if a != b && a != c {
			fmt.Println(a)
		} else if b != a && b != c {
			fmt.Println(b)
		} else {
			fmt.Println(c)
		}
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1915/A
  Tags: implementation, math
  Rating: 800
*/
