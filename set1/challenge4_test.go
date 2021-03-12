package set1

import "testing"

func TestFindEncryptedMessage(t *testing.T) {
	tests := []struct {
		name          string
		wantDecrypted string
	}{
		{
			name:          "ok",
			wantDecrypted: "haha:)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDecrypted := FindEncryptedMessage(); gotDecrypted != tt.wantDecrypted {
				t.Errorf("FindEncryptedMessage() = %v, want %v", gotDecrypted, tt.wantDecrypted)
			}
		})
	}
}
