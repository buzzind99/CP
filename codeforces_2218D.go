//go:build codeforces_2218D

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

func nextInt64() int64 {
	i, _ := strconv.ParseInt(next(), 10, 64)
	return i
}

func sieve(n int) []int {
	var primes []int
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	primes := sieve(200000)
	for range t {
		n := nextInt()
		for i := 0; i < n; i++ {
			val := uint64(primes[i])*uint64(primes[i+1])
			fmt.Fprintf(wr, "%d ", val)
		}
		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2218/D
  Tags: -
  Rating: -
  Contest: Codeforces Round 1090 (Div. 4)
*/
