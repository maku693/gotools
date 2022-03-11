package test

//go:generate gotools run -- stringer -type=Animal

type Animal int

const (
	Dog Animal = iota
	Cat
	Bird
)
