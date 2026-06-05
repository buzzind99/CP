//go:build codeforces_2121B

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

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
	defer wr.Flush()

	t := nextInt()
	for range t {
		n, s := nextInt(), next()
		freq := make([]int, 26)
		for i := 0; i < n; i++ {
			freq[s[i]-'a']++
		}

		found := false
		for i := 1; i < n-1; i++ {
			if freq[s[i]-'a'] > 1 {
				found = true
				break
			}
		}

		if found {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2121/B
  Tags: constructive algorithms, greedy, strings
  Rating: 800
*/
