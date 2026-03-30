//go:build codeforces_1926A

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
        s := next()
		a := 0
		b := 0
		for i := range 5 {
			if s[i] == 'A' { a++ } else { b++ }
			if a == 3 || b == 3 { break }
		}
		if a > b {
			fmt.Println("A")
		} else {
			fmt.Println("B")
		}
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1926/A
  Tags: constructive algorithms, greedy, implementation, strings
  Rating: 800
*/
