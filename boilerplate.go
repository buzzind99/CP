//go:build boilerplate

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ─── I/O ────────────────────────────────────────────────────────────────────

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func init() { sc.Buffer(make([]byte, 1<<20), 1<<20) }
func flush() { wr.Flush() }

func next() string      { sc.Scan(); return sc.Text() }
func nextLine() string  { sc.Scan(); return sc.Text() }
func nextInt() int      { i, _ := strconv.Atoi(next()); return i }
func nextInt64() int64  { i, _ := strconv.ParseInt(next(), 10, 64); return i }
func nextFloat() float64 { f, _ := strconv.ParseFloat(next(), 64); return f }

// Read a whole line of space-separated ints
func nextInts(n int) []int {
	line := nextLine()
	fields := strings.Fields(line)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(fields[i])
	}
	return a
}

func pf(f string, a ...interface{}) { fmt.Fprintf(wr, f, a...) }
func pl(a ...interface{})           { fmt.Fprintln(wr, a...) }

// ─── MATH ────────────────────────────────────────────────────────────────────

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func clamp(x, lo, hi int) int { return max(lo, min(hi, x)) }

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int { return a / gcd(a, b) * b }

// Integer exponentiation
func powInt(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp&1 == 1 {
			result *= base
		}
		base *= base
		exp >>= 1
	}
	return result
}

// Modular exponentiation: (base^exp) % mod
func powMod(base, exp, mod int) int {
	result := 1
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			result = result * base % mod
		}
		base = base * base % mod
		exp >>= 1
	}
	return result
}

// Modular inverse (mod must be prime)
func modInv(a, mod int) int { return powMod(a, mod-2, mod) }

const MOD = 1_000_000_007

func addMod(a, b int) int { return (a + b) % MOD }
func mulMod(a, b int) int { return a % MOD * (b % MOD) % MOD }
func subMod(a, b int) int { return ((a-b)%MOD + MOD) % MOD }

// ─── NUMBER THEORY ───────────────────────────────────────────────────────────

// Sieve of Eratosthenes — returns slice where isPrime[i] == true if i is prime
func sieve(n int) []bool {
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	return isPrime
}

// All prime factors of n (with repetition)
func primeFactors(n int) []int {
	var factors []int
	for p := 2; p*p <= n; p++ {
		for n%p == 0 {
			factors = append(factors, p)
			n /= p
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

// All divisors of n (unsorted)
func divisors(n int) []int {
	var divs []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divs = append(divs, i)
			if i != n/i {
				divs = append(divs, n/i)
			}
		}
	}
	return divs
}

// Precomputed factorials and inverse factorials for combinations
type CombTable struct {
	fact, inv []int
	mod       int
}

func newCombTable(n, mod int) *CombTable {
	fact := make([]int, n+1)
	inv := make([]int, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * i % mod
	}
	inv[n] = powMod(fact[n], mod-2, mod)
	for i := n - 1; i >= 0; i-- {
		inv[i] = inv[i+1] * (i + 1) % mod
	}
	return &CombTable{fact, inv, mod}
}

func (c *CombTable) C(n, r int) int {
	if r < 0 || r > n {
		return 0
	}
	return c.fact[n] * c.inv[r] % c.mod * c.inv[n-r] % c.mod
}

// ─── SORTING & SEARCHING ─────────────────────────────────────────────────────

// Binary search: smallest index i in [lo,hi) where f(i) is true
func lowerBound(lo, hi int, f func(int) bool) int {
	for lo < hi {
		mid := (lo + hi) / 2
		if f(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// Sort helpers
func sortAsc(a []int)  { sort.Ints(a) }
func sortDesc(a []int) { sort.Sort(sort.Reverse(sort.IntSlice(a))) }

// Coordinate compression: returns compressed slice + sorted unique values
func compress(a []int) ([]int, []int) {
	uniq := append([]int{}, a...)
	sort.Ints(uniq)
	uniq = unique(uniq)
	out := make([]int, len(a))
	for i, v := range a {
		out[i] = sort.SearchInts(uniq, v)
	}
	return out, uniq
}

func unique(a []int) []int {
	if len(a) == 0 {
		return a
	}
	res := []int{a[0]}
	for i := 1; i < len(a); i++ {
		if a[i] != a[i-1] {
			res = append(res, a[i])
		}
	}
	return res
}

// ─── DATA STRUCTURES ─────────────────────────────────────────────────────────

// Union-Find (Disjoint Set Union)
type DSU struct {
	parent, rank []int
	components   int
}

func newDSU(n int) *DSU {
	p := make([]int, n)
	r := make([]int, n)
	for i := range p {
		p[i] = i
	}
	return &DSU{p, r, n}
}

func (d *DSU) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) union(x, y int) bool {
	px, py := d.find(x), d.find(y)
	if px == py {
		return false
	}
	if d.rank[px] < d.rank[py] {
		px, py = py, px
	}
	d.parent[py] = px
	if d.rank[px] == d.rank[py] {
		d.rank[px]++
	}
	d.components--
	return true
}

func (d *DSU) same(x, y int) bool { return d.find(x) == d.find(y) }

// Min-heap
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Fenwick Tree (Binary Indexed Tree) — 1-indexed, point update, prefix sum
type BIT struct {
	tree []int
	n    int
}

func newBIT(n int) *BIT { return &BIT{make([]int, n+1), n} }

func (b *BIT) add(i, delta int) {
	for ; i <= b.n; i += i & -i {
		b.tree[i] += delta
	}
}

func (b *BIT) sum(i int) int {
	s := 0
	for ; i > 0; i -= i & -i {
		s += b.tree[i]
	}
	return s
}

func (b *BIT) rangeSum(l, r int) int { return b.sum(r) - b.sum(l-1) }

// Sparse Table — range minimum query in O(1) after O(n log n) build
type SparseTable struct {
	table [][]int
	log2  []int
}

func newSparseTable(a []int) *SparseTable {
	n := len(a)
	k := bits(n) + 1
	table := make([][]int, k)
	table[0] = append([]int{}, a...)
	for j := 1; j < k; j++ {
		table[j] = make([]int, n-(1<<j)+1)
		for i := 0; i+(1<<j) <= n; i++ {
			table[j][i] = min(table[j-1][i], table[j-1][i+(1<<(j-1))])
		}
	}
	log2 := make([]int, n+1)
	for i := 2; i <= n; i++ {
		log2[i] = log2[i/2] + 1
	}
	return &SparseTable{table, log2}
}

// Query returns the minimum in a[l..r] (0-indexed, inclusive)
func (st *SparseTable) query(l, r int) int {
	k := st.log2[r-l+1]
	return min(st.table[k][l], st.table[k][r-(1<<k)+1])
}

func bits(n int) int { return int(math.Log2(float64(n))) }

// ─── GRAPH ───────────────────────────────────────────────────────────────────

type Graph struct {
	adj [][]int
	n   int
}

func newGraph(n int) *Graph { return &Graph{make([][]int, n), n} }

func (g *Graph) addEdge(u, v int)            { g.adj[u] = append(g.adj[u], v) }
func (g *Graph) addUndirected(u, v int)      { g.addEdge(u, v); g.addEdge(v, u) }

// BFS — returns distances from src (-1 if unreachable)
func (g *Graph) bfs(src int) []int {
	dist := make([]int, g.n)
	for i := range dist {
		dist[i] = -1
	}
	dist[src] = 0
	queue := []int{src}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range g.adj[u] {
			if dist[v] == -1 {
				dist[v] = dist[u] + 1
				queue = append(queue, v)
			}
		}
	}
	return dist
}

// DFS — returns visited order
func (g *Graph) dfs(src int) []int {
	visited := make([]bool, g.n)
	var order []int
	var dfsHelper func(u int)
	dfsHelper = func(u int) {
		visited[u] = true
		order = append(order, u)
		for _, v := range g.adj[u] {
			if !visited[v] {
				dfsHelper(v)
			}
		}
	}
	dfsHelper(src)
	return order
}

// Topological sort (Kahn's algorithm) — returns order, or nil if cycle exists
func (g *Graph) topoSort() []int {
	indeg := make([]int, g.n)
	for u := 0; u < g.n; u++ {
		for _, v := range g.adj[u] {
			indeg[v]++
		}
	}
	var queue []int
	for i, d := range indeg {
		if d == 0 {
			queue = append(queue, i)
		}
	}
	var order []int
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		order = append(order, u)
		for _, v := range g.adj[u] {
			indeg[v]--
			if indeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if len(order) != g.n {
		return nil // cycle detected
	}
	return order
}

// Dijkstra — shortest paths from src on a weighted graph
// edges: adj[u] = list of {v, weight}
func dijkstra(n, src int, adj [][][2]int) []int {
	const INF = math.MaxInt64 / 2
	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	dist[src] = 0
	// Using a simple priority queue via sorted slice (fine for most CP problems)
	type item struct{ d, u int }
	pq := []item{{0, src}}
	for len(pq) > 0 {
		// pop minimum
		best := 0
		for i, x := range pq {
			if x.d < pq[best].d {
				best = i
			}
		}
		cur := pq[best]
		pq[best] = pq[len(pq)-1]
		pq = pq[:len(pq)-1]

		if cur.d > dist[cur.u] {
			continue
		}
		for _, e := range adj[cur.u] {
			v, w := e[0], e[1]
			if nd := dist[cur.u] + w; nd < dist[v] {
				dist[v] = nd
				pq = append(pq, item{nd, v})
			}
		}
	}
	return dist
}

// ─── DYNAMIC PROGRAMMING HELPERS ─────────────────────────────────────────────

// Longest Increasing Subsequence — O(n log n), returns length
func lis(a []int) int {
	var tails []int
	for _, x := range a {
		pos := sort.SearchInts(tails, x)
		if pos == len(tails) {
			tails = append(tails, x)
		} else {
			tails[pos] = x
		}
	}
	return len(tails)
}

// 2D prefix sums — build then query rectangle sums in O(1)
type PrefixSum2D struct {
	ps [][]int
}

func newPrefixSum2D(grid [][]int) *PrefixSum2D {
	n, m := len(grid), len(grid[0])
	ps := make([][]int, n+1)
	for i := range ps {
		ps[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			ps[i][j] = grid[i-1][j-1] + ps[i-1][j] + ps[i][j-1] - ps[i-1][j-1]
		}
	}
	return &PrefixSum2D{ps}
}

// Sum of grid[r1..r2][c1..c2] (0-indexed, inclusive)
func (p *PrefixSum2D) query(r1, c1, r2, c2 int) int {
	r1++; c1++; r2++; c2++
	return p.ps[r2][c2] - p.ps[r1-1][c2] - p.ps[r2][c1-1] + p.ps[r1-1][c1-1]
}

// ─── STRING ──────────────────────────────────────────────────────────────────

// KMP failure function
func kmpFailure(pattern string) []int {
	m := len(pattern)
	fail := make([]int, m)
	for i := 1; i < m; i++ {
		j := fail[i-1]
		for j > 0 && pattern[i] != pattern[j] {
			j = fail[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		fail[i] = j
	}
	return fail
}

// KMP search — returns all start indices of pattern in text
func kmpSearch(text, pattern string) []int {
	fail := kmpFailure(pattern)
	var matches []int
	j := 0
	for i := 0; i < len(text); i++ {
		for j > 0 && text[i] != pattern[j] {
			j = fail[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == len(pattern) {
			matches = append(matches, i-j+1)
			j = fail[j-1]
		}
	}
	return matches
}

// ─── GEOMETRY ────────────────────────────────────────────────────────────────

type Point struct{ x, y float64 }

func (p Point) dist(q Point) float64 {
	dx, dy := p.x-q.x, p.y-q.y
	return math.Sqrt(dx*dx + dy*dy)
}

func cross(o, a, b Point) float64 {
	return (a.x-o.x)*(b.y-o.y) - (a.y-o.y)*(b.x-o.x)
}

// Convex hull (Andrew's monotone chain), returns hull in CCW order
func convexHull(pts []Point) []Point {
	n := len(pts)
	if n < 2 {
		return pts
	}
	sort.Slice(pts, func(i, j int) bool {
		if pts[i].x != pts[j].x {
			return pts[i].x < pts[j].x
		}
		return pts[i].y < pts[j].y
	})
	hull := make([]Point, 0, 2*n)
	// Lower
	for _, p := range pts {
		for len(hull) >= 2 && cross(hull[len(hull)-2], hull[len(hull)-1], p) <= 0 {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, p)
	}
	// Upper
	lower := len(hull) + 1
	for i := n - 2; i >= 0; i-- {
		for len(hull) >= lower && cross(hull[len(hull)-2], hull[len(hull)-1], pts[i]) <= 0 {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, pts[i])
	}
	return hull[:len(hull)-1]
}

// ─── MISC ────────────────────────────────────────────────────────────────────

// Permutations — calls f on every permutation of a (modifies a in place)
func permute(a []int, f func([]int)) {
	var helper func(start int)
	helper = func(start int) {
		if start == len(a) {
			f(a)
			return
		}
		for i := start; i < len(a); i++ {
			a[start], a[i] = a[i], a[start]
			helper(start + 1)
			a[start], a[i] = a[i], a[start]
		}
	}
	helper(0)
}

// Next permutation (lexicographic), returns false if already last permutation
func nextPerm(a []int) bool {
	n := len(a)
	i := n - 2
	for i >= 0 && a[i] >= a[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := n - 1
	for a[j] <= a[i] {
		j--
	}
	a[i], a[j] = a[j], a[i]
	for l, r := i+1, n-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
	return true
}

// ─── ENTRY POINT ─────────────────────────────────────────────────────────────

func main() {
	defer flush()
}
