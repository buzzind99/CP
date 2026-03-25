//go:build codeforces_617A

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

    x := nextInt()
	
	fmt.Println((x + 4) / 5)
}

/*
  Link: https://codeforces.com/problemset/problem/617/A
  Tags: math
  Rating: 800
*/
