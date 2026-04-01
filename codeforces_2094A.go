//go:build codeforces_2094A

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
	t := nextInt()
	for range t {
		s := next()
		arr := strings.Split(s, " ")
		for _, v := range arr {
			fmt.Print(string(v[0]))
		}
		fmt.Println()
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2094/A
  Tags: strings
  Rating: 800
*/
