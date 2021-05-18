package main

const Oo = 1e18 // infinite

type Pair struct {
	Ft, Sd int
}

type Pt struct {
	X, Y float64
}

type Pair_dist struct {
	Dist  float64
	Match Pair
}

type Driver struct {
	Name      string
	Latitude  float64
	Longitude float64
	status    string
	Vehicle   string
}

type Order struct {
	Reference     int
	Branch_office string
	Status        string
	Latitude      float64
	Longitude     float64
}

type BranchOffice struct {
	name          string
	Latitude      float64
	Longitude     float64
	Related_order []int
}
