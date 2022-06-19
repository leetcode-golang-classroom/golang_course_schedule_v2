package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	for idx := 0; idx < b.N; idx++ {
		findOrder(numCourses, prerequisites)
	}
}
func Test_findOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "numCourses = 2, prerequisites = [[1,0]]",
			args: args{numCourses: 2, prerequisites: [][]int{{1, 0}}},
			want: []int{0, 1},
		},
		{
			name: "numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]",
			args: args{numCourses: 4, prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "numCourses = 2, prerequisites = [[1,0],[0,1]]",
			args: args{numCourses: 4, prerequisites: [][]int{{1, 0}, {0, 1}}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
