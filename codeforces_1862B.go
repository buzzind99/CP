//go:build codeforces_1862B

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
		n := nextInt()
		prev := 0
		sol := []int{}
		for range n {
			a := nextInt()
			if a >= prev {
				sol = append(sol, a)
			} else {
				sol = append(sol, a)
				sol = append(sol, a)
			}
			prev = a
		}
		num := len(sol)
		fmt.Println(num)
		for _, v := range sol {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1862/B
  Tags: constructive algorithms
  Rating: 800
*/
