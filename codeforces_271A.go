//go:build codeforces_271A

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
    
	y := nextInt()
	for {
		y++
		s := strconv.Itoa(y)
		if s[0] != s[1] && s[0] != s[2] && s[0] != s[3] && s[1] != s[2] && s[1] != s[3] && s[2] != s[3] {
			fmt.Println(y)
			return
		}
    }
}

/*
  Link: https://codeforces.com/problemset/problem/271/A
  Tags: brute force
  Rating: 800
*/
