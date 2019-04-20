package p54

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {

	type args struct {
		matrix [][]int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				matrix: [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "2",
			args: args{
				matrix: [][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}},
			},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name: "3",
			args: args{
				matrix: [][]int{},
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := spiralOrder(tt.args.matrix)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

}
