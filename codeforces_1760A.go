//go:build codeforces_1760A

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
        arr := []int{nextInt(), nextInt(), nextInt()}
        slices.Sort(arr)
        fmt.Println(arr[1])
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1760/A
  Tags: implementation, sortings
  Rating: 800
*/
