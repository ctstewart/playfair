package playfair

import "testing"

func TestDecrypt(t *testing.T) {
	// Test 1 //
	message := "LE"
	key := "palmerston"

	want := "am"
	got, err := Decrypt(message, key)
	if err != nil {
		t.Fatalf("Got error: %s", err)
	}

	if want != got {
		t.Fatalf("\n%10s %v\n%10s %v", "Expected:", want, "Got:", got)
	}

	// Test 2 //
	message = "RP"
	key = "palmerston"

	want = "pv"
	got, err = Decrypt(message, key)
	if err != nil {
		t.Fatalf("Got error: %s", err)
	}

	if want != got {
		t.Fatalf("\n%10s %v\n%10s %v", "Expected:", want, "Got:", got)
	}

	// Test 3 //
	message = "MT"
	key = "palmerston"

	want = "lo"
	got, err = Decrypt(message, key)
	if err != nil {
		t.Fatalf("Got error: %s", err)
	}

	if want != got {
		t.Fatalf("\n%10s %v\n%10s %v", "Expected:", want, "Got:", got)
	}

	// Test 4 //
	message = "MTTBBNESWHTLMPTALNNLNV"
	key = "palmerston"

	want = "lordgranvilxlesletterz"
	got, err = Decrypt(message, key)
	if err != nil {
		t.Fatalf("Got error: %s", err)
	}

	if want != got {
		t.Fatalf("\n%10s %v\n%10s %v", "Expected:", want, "Got:", got)
	}
}
