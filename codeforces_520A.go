//go:build codeforces_520A

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

    n := nextInt()
	s := strings.ToLower(next())
	if n < 26 {
		fmt.Println("NO")
		return
	}

	set := make(map[rune]struct{})
	for _, char := range s {
		set[char] = struct{}{}
	}

	if len(set) == 26 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

/*
  Link: https://codeforces.com/problemset/problem/520/A
  Tags: implementation, strings
  Rating: 800
*/
