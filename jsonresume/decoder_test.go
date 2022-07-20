package jsonresume

import "testing"

func TestDecoder(t *testing.T) {
	_, err := DecodeFile("./mocks/missing.json")
	if err == nil {
		t.Fatal("expected err for missing file")
	}

	_, err = DecodeFile("./mocks/invalid.json")
	if err == nil {
		t.Fatal("expected err for invalid file")
	}

	r, err := DecodeFile("./mocks/resume.json")
	if err != nil {
		t.Fatalf("unxpected error: %s", err)
	}
	if r == nil {
		t.Fatalf("resume should not be nil")
	}

	if r.Basics.Name != "Name" {
		t.Fatalf("parsed values is incorrect: %v", r)
	}
}
