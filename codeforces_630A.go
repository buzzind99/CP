//go:build codeforces_630A

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

    nextInt()
	
	fmt.Println(25)
}

/*
  Link: https://codeforces.com/problemset/problem/630/A
  Tags: number theory
  Rating: 800
*/
