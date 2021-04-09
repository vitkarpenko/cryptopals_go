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

func Test_hammingDistanceTwoBytes(t *testing.T) {
	type args struct {
		first  byte
		second byte
	}
	tests := []struct {
		name         string
		args         args
		wantDistance int
	}{
		{
			name: "0_1",
			args: args{
				first:  0,
				second: 1,
			},
			wantDistance: 1,
		},
		{
			name: "0_2",
			args: args{
				first:  0,
				second: 2,
			},
			wantDistance: 1,
		},
		{
			name: "1_2",
			args: args{
				first:  1,
				second: 2,
			},
			wantDistance: 2,
		},
		{
			name: "8_19",
			args: args{
				first:  8,
				second: 19,
			},
			wantDistance: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance := hammingDistanceTwoBytes(tt.args.first, tt.args.second); gotDistance != tt.wantDistance {
				t.Errorf("hammingDistanceTwoBytes() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
