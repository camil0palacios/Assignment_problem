package main

import (
	"math/rand"
	"sort"
)

const PERCENT_PARENTS = 40
const PERCENT_MUTATION = 10

var population [][]Pair

func generate_match(n, m int, drivers []Driver) []Pair {
	cap := make([]int, n)
	match := make([]bool, m)
	var sample []Pair
	mp := make(map[Pair]bool)
	for i := 0; i < 1000; i++ {
		x := rand.Intn(n)
		y := rand.Intn(m)
		if _, ok := mp[Pair{x, y}]; ok == false && cap[x] <= Get_cap(drivers[x].Vehicle) {
			mp[Pair{x, y}] = true
			match[y] = true
			sample = append(sample, Pair{x, y})
		}
	}
	return sample
}

func init_population(n, m, cnt int, drivers []Driver) [][]Pair {
	var population [][]Pair
	for i := 0; i < cnt; i++ {
		population = append(population, generate_match(n, m, drivers))
	}
	return population
}

func selection_parents(population [][]Pair, orders []Order, drivers []Driver, n_parents int) [][]Pair {
	sort.Slice(population, func(i, j int) bool {
		ptsi := Fill_pts(population[i], orders, drivers)
		ptsj := Fill_pts(population[j], orders, drivers)
		return Get_travel(ptsi, drivers) < Get_travel(ptsj, drivers)
	})
	return population[:n_parents]
}

func cross_parents(parent1 []Pair, parent2 []Pair) []Pair {
	var child []Pair
	mp := make(map[Pair]bool)
	n, m := len(parent1), len(parent2)
	for i := 0; i < n/2; i++ {
		child = append(child, parent1[i])
		mp[parent1[i]] = true
	}
	for i := m / 2; i < m; i++ {
		if _, ok := mp[parent2[i]]; ok == false {
			mp[parent2[i]] = true
			child = append(child, parent2[i])
		}
	}
	return child
}

func cross(parents [][]Pair, tot_pop int) [][]Pair {
	n := len(parents)
	sz_child := tot_pop - len(parents)
	var ans [][]Pair
	for i := 0; i < sz_child; i++ {
		parent1 := parents[rand.Intn(n)]
		parent2 := parents[rand.Intn(n)]
		ans = append(ans, cross_parents(parent1, parent2))
	}
	return ans
}

func mutation(population [][]Pair, tot_pop int) [][]Pair {
	for i := 0; i < tot_pop; i++ {

	}
	return population
}

func Genetic_algorithm(orders []Order, drivers []Driver, tot_pop, it int) float64 {
	n, m := len(drivers), len(orders)
	population := init_population(n, m, tot_pop, drivers)
	for i := 0; i < it; i++ {
		best_parents := selection_parents(population, orders, drivers, int((float64(tot_pop)*PERCENT_PARENTS)/100))
		population = best_parents
		population = append(population, cross(best_parents, tot_pop)...)
		population = mutation(population, int((float64(tot_pop)*PERCENT_MUTATION)/100))
	}
	population = selection_parents(population, orders, drivers, int((float64(tot_pop)*PERCENT_PARENTS)/100))
	ptsbest := Fill_pts(population[0], orders, drivers)
	best := Get_travel(ptsbest, drivers)
	return best
}
