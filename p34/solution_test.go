package p34

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {

	type args struct {
		sorted []int
		target int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{
				sorted: []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: []int{3, 4},
		},
		{
			name: "example2",
			args: args{
				sorted: []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			want: []int{-1, -1},
		},
		{
			name: "example3",
			args: args{
				sorted: []int{7, 7, 7, 7},
				target: 7,
			},
			want: []int{0, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchRange(tt.args.sorted, tt.args.target)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQLERequest() = %v, want %v", got, tt.want)
			}
		})
	}

}
