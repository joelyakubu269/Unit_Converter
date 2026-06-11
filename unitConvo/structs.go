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
