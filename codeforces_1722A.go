//go:build codeforces_1722A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
    "slices"
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
    target := []byte("Timur")
    slices.Sort(target)
    for range t {
        n := nextInt()
        s := next()
        if n != 5 {
            fmt.Println("NO")
            continue
        }

        b := []byte(s)
        slices.Sort(b)
        
        if string(b) == string(target) {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1722/A
  Tags: implementation
  Rating: 800
*/
