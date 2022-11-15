package playfair

import "testing"

func TestEncrypt(t *testing.T) {
	// Test 1 //
	message := "am"
	key := "palmerston"

	want := "LE"
	got, err := Encrypt(message, key)
	if err != nil {
		t.Fatalf("Got error: %s", err)
	}

	if want != got {
		t.Fatalf("\n%10s %v\n%10s %v", "Expected:", want, "Got:", got)
	}

	// Test 2 //
	message = "pv"
	key = "palmerston"

	want = "RP"
	got, err = Encrypt(message, key)
	if err != nil {
		t.Fatalf("Got error: %s", err)
	}

	if want != got {
		t.Fatalf("\n%10s %v\n%10s %v", "Expected:", want, "Got:", got)
	}

	// Test 3 //
	message = "lo"
	key = "palmerston"

	want = "MT"
	got, err = Encrypt(message, key)
	if err != nil {
		t.Fatalf("Got error: %s", err)
	}

	if want != got {
		t.Fatalf("\n%10s %v\n%10s %v", "Expected:", want, "Got:", got)
	}

	// Test 4 //
	message = "lordgranvillesletter"
	key = "palmerston"

	want = "MTTBBNESWHTLMPTALNNLNV"
	got, err = Encrypt(message, key)
	if err != nil {
		t.Fatalf("Got error: %s", err)
	}

	if want != got {
		t.Fatalf("\n%10s %v\n%10s %v", "Expected:", want, "Got:", got)
	}
}
