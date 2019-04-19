package p5

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {

	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "babad",
			},
			want: "aba",
		},
		{
			name: "2",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "3",
			args: args{
				s: "aaaa",
			},
			want: "aaaa",
		},
		{
			name: "4",
			args: args{
				s: "abbaaattuyttyu",
			},
			want: "uyttyu",
		},
		{
			name: "5",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "6",
			args: args{
				s: "",
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.args.s)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

}
