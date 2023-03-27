package sorting

type testCase[T any] struct {
	Name     string
	Values   []T
	Expected []T
}

var testCasesInt = []testCase[int]{
	{
		Name: "simple integer",
		Values: []int{5,3,4,1},
		Expected: []int{1,3,4,5},
	},
	{
		Name: "with negative numbers",
		Values: []int{5,3,-4,1},
		Expected: []int{-4,1,3,5},
	},
	{
		Name: "already ordered",
		Values: []int{1,2,3,4,5},
		Expected: []int{1,2,3,4,5},
	},
}

var testCasesFloat32 = []testCase[float32]{
	{
		Name: "simple float32",
		Values: []float32{5,5.2,5.15,5.01},
		Expected: []float32{5,5.01,5.15,5.2},
	},
	{
		Name: "negative float32",
		Values: []float32{5,-5.2,-5.15,5.01},
		Expected: []float32{-5.2,-5.15,5,5.01},
	},
	{
		Name: "ordered float32",
		Values: []float32{-5,5.01,5.15,5.2},
		Expected: []float32{-5,5.01,5.15,5.2},
	},
}