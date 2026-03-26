//go:build codeforces_228A

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
    
	set := make(map[int]struct{})
	for range 4 {
		n := nextInt()
		set[n] = struct{}{}
	}

	fmt.Println(4-len(set))
}

/*
  Link: https://codeforces.com/problemset/problem/228/A
  Tags: implementation
  Rating: 800
*/
