//go:build codeforces_200B

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
	total := 0
	for range n {
		total += nextInt()
	}
	avg := float64(total)/float64(n)

	fmt.Println(avg)
}

/*
  Link: https://codeforces.com/problemset/problem/200/B
  Tags: implementation, math
  Rating: 800
*/
