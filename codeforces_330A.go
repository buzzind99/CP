//go:build codeforces_330A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	row, col := nextInt(), nextInt()
	arr := make([]string, 0, row)
	ans := 0
	for range row {
		s := next()
		arr = append(arr, s)
	}
	colEat := 0
	for i := range arr {
		if !strings.Contains(arr[i], "S") { ans += col; colEat++ }
	}
	rowEat := 0
	for i := range col {
		eat := true
		for j := range row {
			if arr[j][i] == 'S' { eat = false }
		}
		if eat { ans += row; rowEat++ }
	}

	fmt.Fprintln(wr, ans-rowEat*colEat)
}

/*
  Link: https://codeforces.com/problemset/problem/330/A
  Tags: brute force, implementation
  Rating: 800
*/
