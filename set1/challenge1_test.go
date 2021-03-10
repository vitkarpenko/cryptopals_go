package set1

import "testing"

func TestBase64ToString(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name: "ok",
			args: args{
				HexToBytes("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"),
			},
			wantResult: "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Base64ToString(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("Base64ToString() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
