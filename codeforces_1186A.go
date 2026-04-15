//go:build codeforces_1186A
 
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

	n, m, k := nextInt(), nextInt(), nextInt()

	if m < n || k < n {
		fmt.Fprintln(wr, "No")
	} else {
		fmt.Fprintln(wr, "Yes")
	}
}
 
/*
  Link: https://codeforces.com/problemset/problem/265/A
  Tags: implementation
  Rating: 800
*/
