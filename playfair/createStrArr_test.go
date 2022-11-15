package playfair

import (
	"reflect"
	"testing"
)

func Test_createStrArr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		key     string
		want    []string
		wantErr bool
	}{
		{
			name:    "Basic test",
			key:     "PALMERSTON",
			want:    []string{"P", "A", "L", "M", "E", "R", "S", "T", "O", "N", "B", "C", "D", "F", "G", "H", "I", "K", "Q", "U", "V", "W", "X", "Y", "Z"},
			wantErr: false,
		},
		{
			name:    "No repeating letters",
			key:     "BOB",
			want:    []string{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertKeyToStringArray(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("createStrArr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createStrArr() = %v, want %v", got, tt.want)
			}
		})
	}
}
