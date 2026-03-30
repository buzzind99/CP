//go:build codeforces_1624A

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
        n := nextInt()
        minVal, maxVal := math.MaxInt, 0
        for range n {
            a := nextInt()
            if a < minVal { minVal = a }
            if a > maxVal { maxVal = a }
        }

        fmt.Println(maxVal-minVal)
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1624/A
  Tags: math
  Rating: 800
*/
