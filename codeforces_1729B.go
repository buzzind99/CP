//go:build codeforces_1729B

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
		i := n-1
		var result []byte
		for i >= 0 {
			if s[i] == '0' {
				num := int(s[i-2]-'0')*10 + int(s[i-1]-'0')
				result = append(result, byte('a'+num-1))
				i -= 3
			} else {
				num := int(s[i] - '0')
				result = append(result, byte('a'+num-1))
				i -= 1
			}
		}

		for j, k := 0, len(result)-1; j < k; j, k = j+1, k-1 {
			result[j], result[k] = result[k], result[j]
		}

		fmt.Fprintln(wr, string(result))
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1729/B
  Tags: greedy, strings
  Rating: 800
*/
