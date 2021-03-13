package set1

import "testing"

func Test_hammingDistanceTwoStrings(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name         string
		args         args
		wantDistance int
	}{
		{
			name: "ok",
			args: args{
				first:  "this is a test",
				second: "wokka wokka!!!",
			},
			wantDistance: 37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance := hammingDistanceTwoStrings(tt.args.first, tt.args.second); gotDistance != tt.wantDistance {
				t.Errorf("hammingDistanceTwoStrings() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
