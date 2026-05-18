//go:build codeforces_2055A

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

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		_, a, b := nextInt(), nextInt(), nextInt()
		if isEven(a-b) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2055/A
  Tags: constructive algorithms, games, greedy, math
  Rating: 800
*/
