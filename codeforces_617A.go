//go:build codeforces_617A

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

    x := float64(nextInt())
	if int(x)%5 == 0 {
		fmt.Println(math.Floor(x/5.0))
	} else {
		fmt.Println(math.Floor(x/5.0) + 1)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/617/A
  Tags: math
  Rating: 800
*/
