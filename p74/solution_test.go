package p74

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {

	type args struct {
		matrix [][]int
		target int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				matrix: [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}},
				target: 3,
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				matrix: [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}},
				target: 30,
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				matrix: [][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 50}},
				target: 3,
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				matrix: [][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 50}},
				target: 13,
			},
			want: false,
		},
		{
			name: "5",
			args: args{
				matrix: [][]int{},
				target: 13,
			},
			want: false,
		},
		{
			name: "6",
			args: args{
				matrix: [][]int{[]int{}},
				target: 13,
			},
			want: false,
		},
		{
			name: "7",
			args: args{
				matrix: [][]int{[]int{2}},
				target: 1,
			},
			want: false,
		},
		{
			name: "8",
			args: args{
				matrix: [][]int{[]int{1, 1}},
				target: 2,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchMatrix(tt.args.matrix, tt.args.target)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

}
