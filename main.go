package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	f, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)
	var data [][]string
	for i := 0; i < 10; i++ {
		branch_offices, orders, drivers := Generate_data(5000)
		for _, B := range branch_offices {
			for _, O := range B.Related_order {
				orders[O].Latitude = B.Latitude
				orders[O].Longitude = B.Longitude
			}
		}
		var tmp_drivers []Driver
		for _, D := range drivers {
			if D.status != "order_in_progrees" {
				tmp_drivers = append(tmp_drivers, D)
			}
		}
		drivers = tmp_drivers
		old_t := time.Now()
		ans := Asign(orders, drivers) // Asign
		cur_t := time.Now()
		diff := cur_t.Sub(old_t)
		x, y := Tot_cost(ans, orders, drivers)
		res := []string{fmt.Sprintf("%d", i), "Asignacion basica", fmt.Sprintf("%f", x), fmt.Sprintf("%f", y), fmt.Sprintf("%f", diff.Seconds())}
		data = append(data, res)
		old_t = time.Now()
		ans = Asign_by_aprox(orders, drivers) // Asign by aprox
		x, y = Tot_cost(ans, orders, drivers)
		cur_t = time.Now()
		diff = cur_t.Sub(old_t)
		res = []string{fmt.Sprintf("%d", i), "Asignacion por aprox radial", fmt.Sprintf("%f", x), fmt.Sprintf("%f", y), fmt.Sprintf("%f", diff.Seconds())}
		data = append(data, res)
	}
	w.WriteAll(data)
	w.Flush()
}
