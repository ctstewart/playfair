package playfair

import (
	"reflect"
	"testing"
)

func Test_createLetterPairs(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		input string
		want  [][2]string
	}{
		{
			name:  "Odd number letters",
			input: "lordgranvillesletter",
			want:  [][2]string{{"L", "O"}, {"R", "D"}, {"G", "R"}, {"A", "N"}, {"V", "I"}, {"L", "X"}, {"L", "E"}, {"S", "L"}, {"E", "T"}, {"T", "E"}, {"R", "Z"}},
		},
		{
			name:  "Even number letters",
			input: "lord",
			want:  [][2]string{{"L", "O"}, {"R", "D"}},
		},
		{
			name:  "Invalid letters are dropped",
			input: "lord123@!",
			want:  [][2]string{{"L", "O"}, {"R", "D"}},
		},
		{
			name:  "Invalid letters are dropped (odd)",
			input: "lord123@!s",
			want:  [][2]string{{"L", "O"}, {"R", "D"}, {"S", "Z"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createLetterPairs(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createLetterPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
