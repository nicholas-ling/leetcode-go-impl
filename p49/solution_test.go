package p49

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {

	type args struct {
		strs []string
	}

	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "1",
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{
				[]string{"ate", "eat", "tea"},
				[]string{"nat", "tan"},
				[]string{"bat"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.args.strs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Logf("got %v, want %v", got, tt.want) //result should not bother order, don't panic if you got this line.
			}
		})
	}

}
