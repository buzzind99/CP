//go:build codeforces_265A
 
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

	s, t := next(), next()
	pos := 1
	for i := range t {
		if s[pos-1] == t[i] { pos++ }
	}

	fmt.Fprintln(wr, pos)
}
 
/*
  Link: https://codeforces.com/problemset/problem/265/A
  Tags: implementation
  Rating: 800
*/
