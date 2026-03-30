//go:build codeforces_1900A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
        _, s := nextInt(), next()
		if strings.Contains(s, "...") {
			fmt.Println(2)
		} else {
			fmt.Println(strings.Count(s, "."))
		}
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1900/A
  Tags: constructive algorithms, greedy, implementation, strings
  Rating: 800
*/
