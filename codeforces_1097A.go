//go:build codeforces_1097A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

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
	defer wr.Flush()

	s := next()
	x := ""
	for range 5 {
		x += next()+" "
	}

	possible := false
	for _, char := range s {
		if strings.Contains(x, string(char)) { possible = true; break }
	}

	if possible {
		fmt.Fprintln(wr, "YES")
	} else {
		fmt.Fprintln(wr, "NO")
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1097/A
  Tags: brute force, implementation
  Rating: 800
*/
