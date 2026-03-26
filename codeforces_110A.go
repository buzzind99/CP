//go:build codeforces_110A

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

	s := next()
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '4' || s[i] == '7' { count++ }
	}

	if count == 4 || count == 7 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

/*
  Link: https://codeforces.com/problemset/problem/110/A
  Tags: implementation
  Rating: 800
*/
