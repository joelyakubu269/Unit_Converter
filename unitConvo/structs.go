package main

type PageData struct {
	Type   string
	Units  []string
	Result float64
	Error  string
}
type Input struct {
	Value float64
	From  string
	To    string
}
type Unit struct {
	ToBase   func(float64) float64
	FromBase func(float64) float64
}
