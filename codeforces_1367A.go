//go:build codeforces_1367A

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
        s := next()
        var ans []string
        for i := 0; i < len(s)-1; i += 2 {
            if i == 0 {
                ans = append(ans, s[0:2])
            } else {
                ans = append(ans, string(s[i+1]))
            }
        }
        fmt.Println(strings.Join(ans, ""))
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1367/A
  Tags: implementation, strings
  Rating: 800
*/
