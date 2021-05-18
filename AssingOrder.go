package main

import (
	"sort"
)

// this function assign in order of de input
// complexity: O(n) or in the worst case O(n^2)
func Asign(orders []Order, drivers []Driver) []Pair {
	n, m := len(orders), len(drivers)
	var ans []Pair
	cap := make([]int, n)
	for i, j := 0, 0; i < n; i++ { // x = n / m
		// iterate all the array if the driver does not have enough capacity
		for it := 0; it < n && cap[j]+1 > Get_cap(drivers[j].Vehicle); it++ {
			j = (j + 1) % m
		}
		if cap[j]+1 <= Get_cap(drivers[j].Vehicle) {
			cap[j]++
			ans = append(ans, Pair{j, i})
		}
	}
	return ans
}

// this function assign a driver to an order depending of their radius distance
// complexity: O(n*m*log(n*m))
func Asign_by_aprox(orders []Order, drivers []Driver) []Pair {
	n := len(orders)
	var dist []Pair_dist
	for i, D := range drivers {
		lat1 := D.Latitude
		lon1 := D.Longitude
		vehicle := D.Vehicle
		for j, O := range orders {
			lat2 := O.Latitude
			lon2 := O.Longitude
			d := Get_dist(lat1, lon1, lat2, lon2)
			if d <= Get_rad(vehicle) {
				dist = append(dist, Pair_dist{d, Pair{i, j}})
			}
		}
	}
	cap := make([]int, n)
	var ans []Pair
	sort.Slice(dist, func(i, j int) bool {
		return dist[i].Dist < dist[j].Dist
	})
	match := make([]int, n)
	for i := 0; i < len(dist); i++ {
		dri := dist[i].Match.Ft
		ord := dist[i].Match.Sd
		if match[ord] == 0 && cap[dri]+1 <= Get_cap(drivers[dri].Vehicle) {
			match[ord] = 1
			cap[dri]++
			ans = append(ans, dist[i].Match)
		}
	}
	return ans
}
