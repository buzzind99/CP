//go:build codeforces_1426A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
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
		n, x := nextInt(), nextInt()
		if n == 1 || n == 2 {
			fmt.Println(1)
		} else {
			fmt.Println(math.Ceil(float64(n-2)/float64(x))+1)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1426/A
  Tags: implementation, math
  Rating: 800
*/
