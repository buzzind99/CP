//go:build codeforces_734A

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
	const maxCapacity int = 100005 
    buf := make([]byte, maxCapacity)
    sc.Buffer(buf, maxCapacity)
	sc.Split(bufio.ScanWords)

    n := nextInt()
	s := next()
	count := 0
	for i := 0; i < n; i++ {
		if s[i] == 'A' { count++ } else { count-- }
	}

	if count > 0 {
		fmt.Println("Anton")
	} else if count < 0 {
		fmt.Println("Danik")
	} else {
		fmt.Println("Friendship")
	}
}

/*
  Link: https://codeforces.com/problemset/problem/734/A
  Tags: implementation, strings
  Rating: 800
*/
