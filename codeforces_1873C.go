//go:build codeforces_1873C

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
        points := 0
        for y := 1; y <= 10; y++ {
            shot := next()
            for x := 1; x <= 10; x++ {
                if shot[x-1] == 'X' {
                    distR := min(y, 11-y)
                    distC := min(x, 11-x)
                    points += min(distR, distC)
                }
            }
        }
        fmt.Println(points)
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1873/C
  Tags: implementation, math
  Rating: 800
*/
