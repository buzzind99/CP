//go:build codeforces_977A

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

    n, k := nextInt(), nextInt()
	for range k {
		if n%10 != 0 { n-- } else { n /= 10 }
	}

	fmt.Println(n)
}

/*
  Link: https://codeforces.com/problemset/problem/977/A
  Tags: implementation
  Rating: 800
*/
