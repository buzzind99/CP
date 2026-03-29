//go:build codeforces_1367A

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
        s := next()
        res := []byte{s[0],s[1]}
        for i := 3; i < len(s); i += 2 {
            res = append(res, s[i])
        }
        fmt.Println(string(res))
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1367/A
  Tags: implementation, strings
  Rating: 800
*/
