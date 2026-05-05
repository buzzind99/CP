//go:build codeforces_2037A

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

func abs(x int) int {
	if x < 0 { return -x }
	return x
}

func isEven(x int) bool {
	return x%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		if n == 1 {
			nextInt()
			fmt.Fprintln(wr, 0)
		} else {
			arr := make([]int, 0, n)
			for range n {
				a := nextInt()
				arr = append(arr, a)
			}
			check := make([]bool, n)
			count := 0
			for i := 0; i < n-1; i++ {
				if check[i] { continue }
				for j := i+1; j < n; j++ {
					if arr[i] == arr[j] && !check[i] && !check[j] {
						count++
						check[i], check[j] = true, true
					}
				}
			}
			fmt.Fprintln(wr, count)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2037/A
  Tags: implementation
  Rating: 800
*/
