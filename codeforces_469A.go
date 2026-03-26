//go:build codeforces_469A

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
	set := make(map[int]struct{})

	p := nextInt()
	for range p {
		lvl := nextInt()
		set[lvl] = struct{}{}
	}
	q := nextInt()
	for range q {
		lvl := nextInt()
		set[lvl] = struct{}{}
	}

	if len(set) == n {
		fmt.Println("I become the guy.")
	} else {
		fmt.Println("Oh, my keyboard!")
	}
}

/*
  Link: https://codeforces.com/problemset/problem/469/A
  Tags: greedy, implementation
  Rating: 800
*/
