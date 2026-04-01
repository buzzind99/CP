//go:build codeforces_959A

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

	n := nextInt()
	if n%2 == 0 {
		fmt.Println("Mahmoud")
	} else {
		fmt.Println("Ehab")
	}
}

/*
  Link: https://codeforces.com/problemset/problem/959/A
  Tags: games, math
  Rating: 800
*/