package roulette

import (
	"reflect"
	"testing"
)

func TestGetExpectedFor(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    []int
		wantErr bool
	}{
		{
			name:    "valid number 0",
			input:   0,
			want:    []int{34, 14, 32, 10},
			wantErr: false,
		},
		{
			name:    "valid number 36",
			input:   36,
			want:    []int{16, 36, 1, 12},
			wantErr: false,
		},
		{
			name:    "valid middle number",
			input:   18,
			want:    []int{5, 6, 22, 19},
			wantErr: false,
		},
		{
			name:    "negative number",
			input:   -1,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "number too large",
			input:   37,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetExpectedFor(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetExpectedFor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExpectedFor() = %v, want %v", got, tt.want)
			}
		})
	}
}
