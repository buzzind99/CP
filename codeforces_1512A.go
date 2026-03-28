//go:build codeforces_1512A

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
        n := nextInt()
        arr := make([]int, n)
        for i := range n {
            arr[i] = nextInt()
        }

        common := 0
        if arr[0] == arr[1] {
            common = arr[0]
        } else if arr[0] == arr[2] {
            common = arr[0]
        } else {
            common = arr[1]
        }
        for i, v := range arr {
            if v != common {
                fmt.Println(i+1)
                break
            }
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1512/A
  Tags: brute force, implementation
  Rating: 800
*/
