//go:build codeforces_1985A

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
        a, b := []byte(next()), []byte(next())
		a[0], b[0] = b[0], a[0]
		fmt.Println(string(a), string(b))
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1985/A
  Tags: implementation, math
  Rating: 800
*/
