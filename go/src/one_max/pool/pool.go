package pool

import (
	"errors"
	"fmt"
	"strings"
)

const captive = 100

type GenePool []string

type Pool struct {
	pool [captive * 3]string
	top  int
}

func GetCaptive() int {
	return captive
}

func (pool *Pool) Push(gene string) {
	pool.pool[pool.top] = gene
	pool.top++
}

func (pool *Pool) Pool() GenePool {
	return GenePool(pool.pool[:])
}

func (pool *Pool) SetPool(new_pool GenePool) {
	copy(pool.pool[:], []string(new_pool))
	pool.top = len(new_pool)
}

func (pool *Pool) At(pos int) (string, error) {
	if pos >= pool.top {
		return "", errors.New("Over the pool size")
	}
	return pool.pool[pos], nil
}

func (pool *Pool) Size() int {
	return pool.top
}

func (pool *Pool) String() string {
	result := fmt.Sprintf("Pool: (size: %d)", pool.top)
	for i := 0; i < pool.top; i++ {
		result += fmt.Sprintf("\n\"%s\": %d", pool.pool[i], strings.Count(pool.pool[i], "1"))
	}
	return result
}

func (gene_pool GenePool) Len() int {
	return len(gene_pool)
}

func (gene_pool GenePool) Less(i, j int) bool {
	return strings.Count(gene_pool[i], "1") < strings.Count(gene_pool[j], "1")
}

func (gene_pool GenePool) Swap(i, j int) {
	gene_pool[j], gene_pool[i] = gene_pool[i], gene_pool[j]
}
