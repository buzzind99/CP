//go:build codeforces_266B

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

	n, t := nextInt(), nextInt()
	s := next()
	b := []byte(s)
	for range t {
		for i := 0; i < n-1; i++ {
			if b[i] == 'B' && b[i+1] =='G' {
				b[i], b[i+1] = 'G', 'B'
				i++
			}
		}
	}
    
	fmt.Println(string(b))
}

/*
  Link: https://codeforces.com/problemset/problem/266/B
  Tags: simulaiton, implementation
  Rating: 800
*/
