//go:build codeforces_1328A

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
		a, b := nextInt(), nextInt()
		mod := a%b
		if mod == 0 {
			fmt.Println(0)
		} else {
			fmt.Println(b-mod)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1328/A
  Tags: math
  Rating: 800
*/
