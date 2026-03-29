//go:build codeforces_80A

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

func isPrime(n int) bool {
	if n <= 1 { return false }
	for i := 2; i*i <= n; i++ {
		if n%i == 0 { return false }
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
    
	n, m := nextInt(), nextInt()
	nextP := n + 1
	for !isPrime(nextP) {
		nextP++
	}

	if nextP == m {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

/*
  Link: https://codeforces.com/problemset/problem/80/A
  Tags: brute force
  Rating: 800
*/
