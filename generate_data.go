package main

import (
	"math/rand"

	"github.com/brianvoe/gofakeit"
)

const Min_lat = 3.3333627137544797
const Max_lat = 3.4913538359600014
const Min_lon = -76.58369867130855
const Max_lon = -76.4649090062173

func Generate_data(n int) ([]BranchOffice, []Order, []Driver) {
	driver_status := []string{"ready_to_delivery", "on_cooking"}
	order_status := []string{"free", "order_in_progress"}
	vehicle := []string{"car", "motorcycle", "bicycle"}
	var names []string
	for i := 0; i < n; i++ {
		names = append(names, gofakeit.Company())
	}
	var orders []Order
	var drivers []Driver
	map_branch_offices := make(map[string]BranchOffice)
	for i := 0; i < n; i++ {
		rand_name := names[rand.Intn(len(names))]
		rand_lat, err := gofakeit.LatitudeInRange(Min_lat, Max_lat)
		rand_lon, err := gofakeit.LongitudeInRange(Min_lon, Max_lon)

		if err != nil {
			break
		}

		if _, ok := map_branch_offices[rand_name]; ok {
			b := map_branch_offices[rand_name]
			b.Related_order = append(b.Related_order, i)
			map_branch_offices[rand_name] = b
		} else {
			var x []int
			map_branch_offices[rand_name] = BranchOffice{
				rand_name,
				rand_lat,
				rand_lon,
				x}
		}

		orders = append(orders, Order{
			i,
			rand_name,
			order_status[rand.Intn(2)], 0, 0})

		rand_lat, err = gofakeit.LatitudeInRange(Min_lat, Max_lat)
		rand_lon, err = gofakeit.LongitudeInRange(Min_lon, Max_lon)

		drivers = append(drivers, Driver{
			gofakeit.Name(),
			rand_lat,
			rand_lon,
			driver_status[rand.Intn(2)],
			vehicle[rand.Intn(3)]})
	}
	var branch_offices []BranchOffice
	for _, val := range map_branch_offices {
		branch_offices = append(branch_offices, val)
	}
	return branch_offices, orders, drivers
}
