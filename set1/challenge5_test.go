package set1

import (
	"reflect"
	"testing"
)

func TestXORWithRepeatingKey(t *testing.T) {
	type args struct {
		messageBytes []byte
		key          []byte
	}
	tests := []struct {
		name       string
		args       args
		wantResult []byte
	}{
		{
			name: "ok",
			args: args{
				messageBytes: []byte(
					"Burning 'em, if you ain't quick and nimble\n" +
						"I go crazy when I hear a cymbal",
				),
				key: []byte("ICE"),
			},
			wantResult: HexToBytes(
				"0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c" +
					"2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b" +
					"2027630c692b20283165286326302e27282f",
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := XORWithRepeatingKey(tt.args.messageBytes, tt.args.key); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("XORWithRepeatingKey() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
