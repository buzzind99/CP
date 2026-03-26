//go:build codeforces_41A

package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
    
	s, t := next(), next()
	if len(s) != len(t) {
        fmt.Println("NO")
        return
    }
	n := len(s)
	for i := range n {
		if (s[i] != t[n-1-i]) {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}

/*
  Link: https://codeforces.com/problemset/problem/41/A
  Tags: implementation, strings
  Rating: 800
*/
