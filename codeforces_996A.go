//go:build codeforces_996A

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
	bills := []int{100, 20, 10, 5, 1}
    count := 0
    for _, bill := range bills {
        count += n / bill
        n %= bill
    }

	fmt.Println(count)
}

/*
  Link: https://codeforces.com/problemset/problem/996/A
  Tags: implementation
  Rating: 800
*/
