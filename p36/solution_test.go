package p36

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {

	type args struct {
		board [][]byte
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				board: [][]byte{
					[]byte("53..7...."),
					[]byte("6..195..."),
					[]byte(".98....6."),
					[]byte("8...6...3"),
					[]byte("4..8.3..1"),
					[]byte("7...2...6"),
					[]byte(".6....28."),
					[]byte("...419..5"),
					[]byte("....8..79"),
				},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				board: [][]byte{
					[]byte("83..7...."),
					[]byte("6..195..."),
					[]byte(".98....6."),
					[]byte("8...6...3"),
					[]byte("4..8.3..1"),
					[]byte("7...2...6"),
					[]byte(".6....28."),
					[]byte("...419..5"),
					[]byte("....8..79"),
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidSudoku(tt.args.board)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

}
