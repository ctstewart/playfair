package playfair

import (
	"reflect"
	"testing"
)

func TestGetCoordFromLetter(t *testing.T) {
	strArr, err := convertKeyToStringArray("PALMERSTON")
	if err != nil {
		t.Fatalf("Got error: %s", err)
	}
	type testStruct struct {
		letter string
		x      int
		y      int
	}
	tests := []testStruct{{letter: "P", x: 0, y: 0}, {letter: "T", x: 2, y: 1}, {letter: "Z", x: 4, y: 4}, {letter: "Q", x: 3, y: 3}}
	for _, test := range tests {
		want := [2]int{test.x, test.y}
		got := getCoordFromLetter(test.letter, strArr)
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("Wanted %v, got %v", want, got)
		}
	}
}
