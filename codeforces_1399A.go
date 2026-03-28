//go:build codeforces_1399A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
    "slices"
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
        n := nextInt()
        arr := make([]int, n)
        for i := range n {
            arr[i] = nextInt()
        }
        slices.Sort(arr)
        possible := true
        for i := 0; i < n-1; i++ {
            if arr[i+1]-arr[i] > 1 {
                possible = false
                break
            }
        }
        if possible {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1399/A
  Tags: greedy, sortings
  Rating: 800
*/
