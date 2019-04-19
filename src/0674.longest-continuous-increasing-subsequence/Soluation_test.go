package Soluation

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	// 测试用例
	cases := []struct {
		name 	string
		inputs	[]int
		expect 	int
	}{
		{
			"TestCases 1",
			[]int{1,3,5,4,7},
			3,
		},
		{
			"TestCases 2",
			[]int{2,2,2,2,2},
			1,
		},
	}

	// 开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := findLengthOfLCIS(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v", c.expect, got, c.inputs)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {

}

//	使用案列
func ExampleSolution() {

}
