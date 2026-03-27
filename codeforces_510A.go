//go:build codeforces_510A

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

    n, m := nextInt(), nextInt()
	r := true
	for i := range n {
		for j := range m {
			if i%2 == 0 {
				fmt.Print("#")
			} else {
				if j == 0 && !r {
					fmt.Print("#")
				} else if j == m-1 && r {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
		}
		if i%2 != 0 { r = !r }
		fmt.Println()
	}
}

/*
  Link: https://codeforces.com/problemset/problem/510/A
  Tags: implementation
  Rating: 800
*/
