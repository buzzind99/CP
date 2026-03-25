//go:build codeforces_236A

package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
    
	str := next()
	set := make(map[rune]struct{})

	for _, char := range str {
		set[char] = struct{}{}
	}

	if len(set)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}

/*
  Link: https://codeforces.com/problemset/problem/236/A
  Tags: brute force, implementation, strings
  Rating: 800
*/
