package core

type NodeInfo struct {
	street        string
	name          string
	originalValue float64
	value         float64
}

type Node struct {
	X int
	Y int
}

type StreetInfo struct {
	Street     string
	IsEven     bool
	Multiplier float64
	Length     int
}
