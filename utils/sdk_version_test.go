package utils

import "testing"

func TestGetSDKVersion(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test GetSDKVersion",
			want: "2.0.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSDKVersion(); got != tt.want {
				t.Errorf("GetSDKVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
