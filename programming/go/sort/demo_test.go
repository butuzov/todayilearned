package sortingexample

import (
	"sort"
	"testing"
)

// Example of struct we going to sort.

type Point struct {
	X, Y int
}

// --- Struct / Raw Data
var TestCases = []Point{
	{10, 3},
	{10, 4},
	{10, 35},
	{10, 5},
	{10, 51},
	{10, 25},
	{10, 59},
	{10, 15},
	{10, 22},
	{10, 91},
}

// Example One - Sorting Slice Directly
// somehow - slowest way to sort it.
func SortSlice(points []Point) {

	sort.Slice(points, func(i, j int) bool {
		return points[i].Y < points[j].Y
	})
}

func BenchmarkSlice(b *testing.B) {
	tmp := make([]Point, len(TestCases))
	for i := 0; i < b.N; i++ {
		copy(tmp, TestCases)
		SortSlice(tmp)
	}
}

// Example Two - Sorting Slice Directly
// much faster performance
type Points []Point

// Sort interface implementation
func (p Points) Less(i, j int) bool { return p[i].Y < p[j].Y }
func (p Points) Len() int           { return len(p) }
func (p Points) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func SortStruct(points []Point) {
	sort.Sort(Points(points))
}

func BenchmarkStruct(b *testing.B) {
	tmp := make([]Point, len(TestCases))
	for i := 0; i < b.N; i++ {
		copy(tmp, TestCases)
		SortStruct(tmp)
	}
}

// --- Pointers
var TestCasesPoints = []*Point{
	&Point{10, 3},
	&Point{10, 4},
	&Point{10, 35},
	&Point{10, 5},
	&Point{10, 51},
	&Point{10, 25},
	&Point{10, 59},
	&Point{10, 15},
	&Point{10, 22},
	&Point{10, 91},
}

// Example Three - Sorting Slice of Pointers

func SortSlicePointers(points []*Point) {
	sort.Slice(points, func(i, j int) bool {
		return points[i].Y < points[j].Y
	})
}

func BenchmarkSlicePointers(b *testing.B) {
	tmp := make([]*Point, len(TestCasesPoints))
	for i := 0; i < b.N; i++ {
		copy(tmp, TestCasesPoints)
		SortSlicePointers(tmp)
	}
}

// Example Four - Sorting Struct (with Slice of pointers beneath it)
type PointsPointer []*Point

func (pp PointsPointer) Less(i, j int) bool { return pp[i].Y < pp[j].Y }
func (pp PointsPointer) Len() int           { return len(pp) }
func (pp PointsPointer) Swap(i, j int)      { pp[i], pp[j] = pp[j], pp[i] }

func SortStructOfSlicePointers(points []*Point) {
	sort.Sort(PointsPointer(points))
}

func BenchmarkStructOfSlicePointers(b *testing.B) {
	tmp := make([]*Point, len(TestCasesPoints))

	for i := 0; i < b.N; i++ {
		copy(tmp, TestCasesPoints)
		SortStructOfSlicePointers(tmp)
	}
}
