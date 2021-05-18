package main

import (
	"math"
)

func Tot_cost(pairs []Pair, orders []Order, drivers []Driver) (float64, float64) {
	pts := Fill_pts(pairs, orders, drivers)
	ans1 := 0.0
	// time estimated without tsp
	x, y := 0, 0
	for i, e := range pts {
		if len(pts[i]) > 0 {
			ans1 += travel(e, drivers[i].Vehicle)
			x++
		}
	}
	// time estimated with tsp
	ans2 := 0.0
	for i, e := range pts {
		if len(pts[i]) > 0 {
			ans2 += Tsp(e, drivers[i].Vehicle)
			y++
		}
	}
	return ans1 / float64(x), ans2 / float64(y)
	// fmt.Printf("Total de asignaciones: %d\n", len(pairs))
	// fmt.Printf("Tiempo en promedio sin tsp: %f\n", ans1/float64(x))
	// fmt.Printf("Tiempo en promedio con tsp: %f\n", ans2/float64(y))
}

func Fill_pts(pairs []Pair, orders []Order, drivers []Driver) [][]Pt {
	m := len(drivers)
	// time estimated without tsp
	pts := make([][]Pt, m)
	for i, D := range drivers {
		pts[i] = append(pts[i], Pt{D.Latitude, D.Longitude})
	}
	for _, P := range pairs {
		pts[P.Ft] = append(pts[P.Ft], Pt{orders[P.Sd].Latitude, orders[P.Sd].Longitude})
	}
	return pts
}

// obtain the
func Get_travel(pts [][]Pt, drivers []Driver) float64 {
	ans, x := 0.0, 0
	for i, e := range pts {
		if len(pts[i]) > 0 {
			ans += travel(e, drivers[i].Vehicle)
			x++
		}
	}
	return ans / float64(x)
}

// travel point by point
func travel(pts []Pt, vehicle string) float64 {
	n := len(pts)
	ans := 0.0
	for i := 1; i < n; i++ {
		ans += Get_time(pts[i-1].X, pts[i-1].Y, pts[i].X, pts[i].Y, vehicle)
	}
	return ans / 60
}

// travel salesman problem
func Tsp(pts []Pt, vehicle string) float64 {
	n := len(pts)
	// dp[msk][lst] = min(dp[msk | (1 << i)][lst] + t(lst, i))
	var dp [1 << 11][11]float64
	for i := 0; i < (1 << n); i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = Oo
		}
	}
	for i := 0; i < n; i++ {
		dp[(1<<n)-1][i] = 0
	}
	for msk := (1 << n) - 2; msk >= 0; msk-- {
		for lst := 0; lst < n; lst++ {
			for i := 0; i < n; i++ {
				if ((msk >> i) & 1) == 0 {
					x1, y1 := pts[lst].X, pts[lst].Y
					x2, y2 := pts[i].X, pts[i].Y
					dp[msk][lst] = math.Min(
						dp[msk][lst],
						dp[msk|(1<<i)][i]+Get_time(x1, y1, x2, y2, vehicle))
				}
			}
		}
	}
	return dp[1][0] / 60
}

// get maximum capacity depending of the vehicle
func Get_cap(x string) int {
	if x == "car" {
		return 10
	}
	if x == "motorcycle" {
		return 6
	}
	return 4
}

// get average velocity depending of the vehicle
func Get_vel(s string) float64 {
	if s == "car" {
		return 35
	}
	if s == "motorcycle" {
		return 25
	}
	return 10
}

// get maximum radius
func Get_rad(s string) float64 {
	if s == "car" {
		return 25
	}
	if s == "motorcycle" {
		return 10
	}
	return 4
}

// convert to radians
func To_rad(x float64) float64 {
	return (math.Pi / 180.0) * x
}

// get distance between two points
func Get_dist(lat1, lon1, lat2, lon2 float64) float64 {
	R := 6378.0 // radius of the earth
	dlat := To_rad(lat2 - lat1)
	dlon := To_rad(lon2 - lon1)
	SinLat := math.Sin(dlat / 2)
	SinLong := math.Sin(dlon / 2)
	a := SinLat*SinLat + math.Cos(To_rad(lat1))*math.Cos(To_rad(lat2))*SinLong*SinLong
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

// get time depending of the vehicle
func Get_time(x1, y1, x2, y2 float64, vehicle string) float64 {
	return Get_dist(x1, y1, x2, y2) / Get_vel(vehicle)
}
