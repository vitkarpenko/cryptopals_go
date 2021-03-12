package set1

import (
	"testing"
)

func Test_scoreStringAsEnglish(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name      string
		args      args
		wantScore int
	}{
		{
			name:      "all_in",
			args:      args{message: "riot eeeeah!"},
			wantScore: 75,
		}, {
			name:      "some_in",
			args:      args{message: "ear Zz"},
			wantScore: 27,
		}, {
			name:      "none_in",
			args:      args{message: "xyz"},
			wantScore: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotScore := scoreStringAsEnglish(tt.args.message); gotScore != tt.wantScore {
				t.Errorf("scoreStringAsEnglish() = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}

func TestDecryptXORedMessage(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{message: "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"},
			want: "Cooking MC's like a pound of bacon",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecryptXORedMessage(tt.args.message); got != tt.want {
				t.Errorf("DecryptXORedMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
