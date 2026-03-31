//go:build codeforces_1881A

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
        nextInt()
        nextInt()
        x, s := next(), next()
        matched := false
        count := 0
        for count <= 6 {
            if strings.Contains(x, s) {
                fmt.Println(count)
                matched = true
                break
            }
            x += x
            count++
        }
        if !matched {
            fmt.Println(-1)
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1881/A
  Tags: brute force, strings
  Rating: 800
*/
