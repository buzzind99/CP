//go:build codeforces_486A

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
    
	n := nextInt()
	res := 0
	if n%2 != 0 { res = -(n+1)/2 } else { res = n/2 }

	fmt.Println(res)
}

/*
  Link: https://codeforces.com/problemset/problem/486/A
  Tags: implementation, math
  Rating: 800
*/
