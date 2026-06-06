//go:build codeforces_1974B

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"slices"
	"strings"
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

func removeDuplicates(input string) string {
	seen := make(map[rune]struct{})
	var result strings.Builder
	for _, char := range input {
		if _, exists := seen[char]; !exists {
			seen[char] = struct{}{}
			result.WriteRune(char)
		}
	}

	return result.String()
}

func main() {
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n, s := nextInt(), next()
		uniqueS := removeDuplicates(s)
		if len(uniqueS) == 1 {
			fmt.Fprintln(wr, s)
			continue
		}

		bytes := []byte(uniqueS)
		slices.Sort(bytes)
		m := make(map[byte]byte)
		for left := range len(bytes)-1 {
			right := len(bytes)-1-left
			if left > right { break }
			m[bytes[left]] = bytes[right]
			if left != right {
				m[bytes[right]] = bytes[left]
			}
		}

		for i := range n {
			fmt.Fprintf(wr, "%c", m[s[i]])
		}
		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1974/B
  Tags: implementation, strings, sortings
  Rating: 800
*/
