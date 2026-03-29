//go:build codeforces_490A

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
	arr := make([]int, n)
	a, b, c := 0, 0, 0
	for i := range n {
		n := nextInt()
		if n == 1 { a++ } else if n == 2 { b++ } else { c++ }
		arr[i] = n
	}

	minC := min(a, b, c)
	fmt.Println(minC)
	for range minC {
		curr := 1
		for i := 0; i < len(arr); i++ {
			if curr > 3 { break }
			if arr[i] == curr {
				fmt.Print(i+1, " ")
				arr[i] = 0
				i = -1
				curr++
			}
		}
		fmt.Println()
	}
}

/*
  Link: https://codeforces.com/problemset/problem/490/A
  Tags: greedy, implementation, sortings
  Rating: 800
*/
