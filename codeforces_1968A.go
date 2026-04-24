//go:build codeforces_1968A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func gcd(a int, b int) int {
	for b != 0 { a, b = b, a%b }
	return a
}

func isPrime(n int) bool {
	if n <= 1 { return false }
	for i := 2; i*i <= n; i++ {
		if n%i == 0 { return false }
	}
	return true
}

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		x := nextInt()

		fmt.Fprintln(wr, x-1)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1933/A
  Tags: brute force, math, number theory
  Rating: 800
*/
