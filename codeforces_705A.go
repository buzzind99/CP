//go:build codeforces_705A

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
	for i := range n {
		if (i+1)%2 != 0 {
			fmt.Print("I hate ")
		} else {
			fmt.Print("I love ")
		}
		if i != n-1 {
			fmt.Print("that ")
		}
	}
	fmt.Print("it")
}

/*
  Link: https://codeforces.com/problemset/problem/705/A
  Tags: implementation
  Rating: 800
*/
