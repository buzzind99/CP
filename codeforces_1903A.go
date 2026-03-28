//go:build codeforces_1903A

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

	t := nextInt()
    for range t {
        n, k := nextInt(), nextInt()
		arr := make([]int, n)
		sorted := true
		for i := range n {
			a := nextInt()
			arr[i] = a
			if i > 0 && arr[i] < arr[i-1] { sorted = false }
		}
		if k == 1 && !sorted {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1903/A
  Tags: brute force, greedy, sortings
  Rating: 800
*/
