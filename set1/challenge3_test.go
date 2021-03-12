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
			args:      args{message: "riot"},
			wantScore: 4,
		}, {
			name:      "some_in",
			args:      args{message: "ear Zzzzzz"},
			wantScore: 3,
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
