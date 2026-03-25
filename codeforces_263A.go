//go:build codeforces_263A

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

func abs(x int) int {
    if x < 0 { return -x }
    return x
}

func main() {
	sc.Split(bufio.ScanWords)
    
	for y := 1; y <= 5; y++ {
		for x := 1; x <= 5; x++ {
			n := nextInt()
			if n == 1 {
				fmt.Println(abs(x-3) + abs(y-3))
				return
			}
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/263/A
  Tags: implementation
  Rating: 800
*/
