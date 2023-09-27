package main

import (
	"fmt"
	"sort"
)

type Data struct {
	Type string
	Size int
	Name string
}

type SortStrategy interface {
	Sort(data []Data)
}

type NameSorter struct{}

func (n NameSorter) Sort(data []Data) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Name < data[j].Name
	})
}

type TypeSorter struct{}

func (t TypeSorter) Sort(data []Data) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Type < data[j].Type
	})
}

type SizeSorter struct{}

func (s SizeSorter) Sort(data []Data) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Size < data[j].Size
	})
}

type SortContext struct {
	strategy SortStrategy
}

func (s *SortContext) SetStrategy(strategy SortStrategy) {
	s.strategy = strategy
}

func (s *SortContext) ExecuteSort(data []Data) {
	s.strategy.Sort(data)
}

func main() {
	data := []Data{
		{"bbbbbb", 5, "Olzhas"},
		{"ababab", 15, "Shredder"},
		{"cccccc", 30, "Leo"},
	}

	context := SortContext{strategy: TypeSorter{}}
	context.ExecuteSort(data)
	fmt.Println("Sorted by Type:", data)

	context.SetStrategy(NameSorter{})
	context.ExecuteSort(data)
	fmt.Println("Sorted by Name:", data)

	context.SetStrategy(SizeSorter{})
	context.ExecuteSort(data)
	fmt.Println("Sorted by Size:", data)

}
