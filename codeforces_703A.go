//go:build codeforces_703A

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
	mishka := 0
	chris := 0
	for range n {
		m, c := nextInt(), nextInt()
		if m > c { mishka++ } else if c > m { chris++ }
	}
	
	if mishka > chris {
		fmt.Println("Mishka")
	} else if chris > mishka {
		fmt.Println("Chris")
	} else {
		fmt.Println("Friendship is magic!^^")
	}
}

/*
  Link: https://codeforces.com/problemset/problem/703/A
  Tags: implementation
  Rating: 800
*/
