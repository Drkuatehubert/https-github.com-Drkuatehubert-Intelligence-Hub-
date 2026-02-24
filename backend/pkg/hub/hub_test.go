package hub

import (
	"testing"
)

func TestClassifyDevice(t *testing.T) {
	tests := []struct {
		name     string
		origType string
		want     string
	}{
		{"My Tesla", "bluetooth", "car"},
		{"LG Smart TV", "router", "tv"},
		{"Bose QuietComfort", "bluetooth", "headphone"},
		{"Generic Cam", "router", "camera"},
		{"Unknown Device", "iot", "iot"},
	}

	for _, tt := range tests {
		if got := ClassifyDevice(tt.name, tt.origType); got != tt.want {
			t.Errorf("ClassifyDevice(%q, %q) = %q, want %q", tt.name, tt.origType, got, tt.want)
		}
	}
}

func TestGetRegistry(t *testing.T) {
	registry := GetRegistry()
	if len(registry) == 0 {
		t.Error("Registry should not be empty")
	}
}
