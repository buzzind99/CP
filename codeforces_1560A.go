//go:build codeforces_1560A

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
        k := nextInt()
        v := 0
        for i := 0; i < k; {
            v++
            if v%3 == 0 || v%10 == 3 { continue }
            i++
        }
        fmt.Println(v)
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1560/A
  Tags: implementation
  Rating: 800
*/
