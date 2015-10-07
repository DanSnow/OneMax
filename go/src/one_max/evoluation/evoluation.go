package evoluation

import (
	"../pool"
	"math/rand"
	"sort"
	"strings"
)

const (
	STR_LEN         = 100
	CROSS_MAX       = 5
	EVOLUATION_TIME = 250
)

func Evoluation(pool *pool.Pool) {
	Seed(pool)
	for i := 0; i < EVOLUATION_TIME; i++ {
		Multiplication(pool)
		Eliminate(pool)
	}
}

func Seed(pool *pool.Pool) {
	pool.Push(rand_str())
	pool.Push(rand_str())
}

func Multiplication(p *pool.Pool) {
	target := pool.GetCaptive() * 2
	for p.Size() < target {
		parent1, parent2 := pick_parent(p)
		child1, child2 := cross(parent1, parent2)
		child1 = mutation(child1)
		child2 = mutation(child2)
		p.Push(child1)
		p.Push(child2)
	}
}

func Eliminate(p *pool.Pool) {
	new_pool := p.Pool()
	sort.Sort(sort.Reverse(new_pool))
	new_pool = new_pool[:pool.GetCaptive()]
	p.SetPool(new_pool)
}

func mutation(g string) string {
	gene := []byte(g)
	n := rand.Intn(STR_LEN) / (rand.Intn(STR_LEN/2) + 1)
	for i := 0; i < n; i++ {
		pos := rand.Intn(STR_LEN)
		gene[pos] = reverse(gene[pos])
	}
	return string(gene)
}

func rand_str() string {
	str := make([]string, STR_LEN)
	for i := 0; i < STR_LEN; i++ {
		str[i] = string(rand.Intn(2) + '0')
	}
	return strings.Join(str, "")
}

func pick_parent(pool *pool.Pool) (string, string) {
	var pos int
	size := pool.Size()
	pos = rand.Intn(size)
	p1, _ := pool.At(pos)
	pos = rand.Intn(size)
	p2, _ := pool.At(pos)
	return p1, p2
}

func cross(parent_a, parent_b string) (child_a, child_b string) {
	n := rand.Intn(CROSS_MAX-1) + 1
	child_a = parent_a
	child_b = parent_b
	for i := 0; i < n; i++ {
		child_a, child_b = make_cross(child_a, child_b)
	}
	return
}

func make_cross(parent_a, parent_b string) (child_a, child_b string) {
	n := rand.Intn(STR_LEN)
	child_a = parent_a[:n] + parent_b[n:]
	child_b = parent_b[:n] + parent_a[n:]
	return
}

func reverse(ch byte) byte {
	if ch == '1' {
		return '0'
	} else {
		return '1'
	}
}
