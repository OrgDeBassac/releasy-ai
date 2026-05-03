package semver

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected Version
		wantErr  bool
	}{
		{"v1.2.3", Version{1, 2, 3}, false},
		{"1.2.3", Version{1, 2, 3}, false},
		{"v0.0.1", Version{0, 0, 1}, false},
		{"v1.0", Version{}, true},
		{"v1.2.3.4", Version{}, true},
		{"abc", Version{}, true},
	}

	for _, tt := range tests {
		v, err := Parse(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("Parse(%s) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			continue
		}
		if !tt.wantErr && v != tt.expected {
			t.Errorf("Parse(%s) = %v, want %v", tt.input, v, tt.expected)
		}
	}
}

func TestIncrements(t *testing.T) {
	v := Version{1, 2, 3}

	if next := v.NextMajor(); next != (Version{2, 0, 0}) {
		t.Errorf("NextMajor = %v, want v2.0.0", next)
	}

	if next := v.NextMinor(); next != (Version{1, 3, 0}) {
		t.Errorf("NextMinor = %v, want v1.3.0", next)
	}

	if next := v.NextPatch(); next != (Version{1, 2, 4}) {
		t.Errorf("NextPatch = %v, want v1.2.4", next)
	}
}
