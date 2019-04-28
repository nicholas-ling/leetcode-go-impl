package p56

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type args struct {
		intervals [][]int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				[][]int{
					[]int{1, 3},
					[]int{2, 4},
					[]int{5, 6},
				},
			},
			want: [][]int{
				[]int{1, 4},
				[]int{5, 6},
			},
		},
		{
			name: "2",
			args: args{
				[][]int{
					[]int{1, 4},
					[]int{4, 5},
				},
			},
			want: [][]int{
				[]int{1, 5},
			},
		},
		{
			name: "3",
			args: args{
				[][]int{
					[]int{1, 3},
					[]int{2, 6},
					[]int{8, 10},
					[]int{15, 18},
				},
			},
			want: [][]int{
				[]int{1, 6},
				[]int{8, 10},
				[]int{15, 18},
			},
		},
		{
			name: "4",
			args: args{
				[][]int{
					[]int{8, 10},
					[]int{1, 3},
					[]int{15, 18},
					[]int{2, 6},
				},
			},
			want: [][]int{
				[]int{1, 6},
				[]int{8, 10},
				[]int{15, 18},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.args.intervals)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

}
