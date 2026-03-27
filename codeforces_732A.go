//go:build codeforces_732A

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

    k, r := nextInt(), nextInt()
	count := 1
	for (k*count)%10 != 0 && (k*count)%10 != r { count++ }

	fmt.Println(count)
}

/*
  Link: https://codeforces.com/problemset/problem/732/A
  Tags: brute force, constructive algorithms, implementation, math
  Rating: 800
*/
