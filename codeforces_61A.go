//go:build codeforces_61A

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
    
	n, m := next(), next()
	res := make([]byte, len(n))
	for i := range m {
		if n[i] != m[i] { res[i] = '1' } else { res[i] = '0'}
	}
	
	fmt.Println(string(res))
}

/*
  Link: https://codeforces.com/problemset/problem/61/A
  Tags: implementation
  Rating: 800
*/
