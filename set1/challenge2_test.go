package set1

import (
	"reflect"
	"testing"
)

func TestFixedXOR(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []byte
	}{
		{
			name: "ok",
			args: args{
				first:  "1c0111001f010100061a024b53535009181c",
				second: "686974207468652062756c6c277320657965",
			},
			wantResult: HexToBytes("746865206b696420646f6e277420706c6179"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := FixedXOR(tt.args.first, tt.args.second); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("FixedXOR() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
