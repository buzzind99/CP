//go:build codeforces_71A

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
		solve()
	}
}

func solve() {
	str := next()
	n := len(str)

    if n > 10 {
        fmt.Printf("%c%d%c\n", str[0], n-2, str[n-1])
    } else {
        fmt.Println(str)
    }
}

/*
  Link: https://codeforces.com/problemset/problem/71/A
  Tags: strings
  Rating: 800
*/
