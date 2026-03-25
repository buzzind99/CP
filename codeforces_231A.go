//go:build codeforces_231A

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
	sol := 0
	sc.Split(bufio.ScanWords)
    
	n := nextInt()
	for range n {
		sure := count()
		if sure >= 2 {
			sol++
		}
	}

	fmt.Println(sol)
}

func count() int {
	petya := nextInt()
	vasya := nextInt()
	tonya := nextInt()

	return petya + vasya + tonya
}

/*
  Link: https://codeforces.com/problemset/problem/231/A
  Tags: brute force, greedy
  Rating: 800
*/
