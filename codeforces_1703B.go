//go:build codeforces_1703B

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
        m := make([]bool, 26)
        n := nextInt()
        s := next()
        sum := 0
        for i := range n {
            idx := s[i]-'A'
            if !m[idx] {
                m[idx] = true
                sum += 2
            } else {
                sum++
            }
        }
        fmt.Println(sum)
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1703/B
  Tags: data structures, implementation
  Rating: 800
*/
