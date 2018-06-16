package uuid

import (
	"strings"
	"testing"
)

func TestDuplicates(t *testing.T) {
	var id UUID
	var hex string

	dups := make(map[string]bool)
	for i := 0; i < 1024; i++ {
		id = Rand()
		hex = id.Hex()
		if dups[hex] {
			t.Errorf("Duplicates after %d iterations", i+1)
			t.FailNow()
		}
		dups[hex] = true
	}
}

func TestFromStrSanity(t *testing.T) {
	var id, id2 UUID
	for i := 0; i < 18; i++ {
		id = Rand()
		id2 = MustFromStr(id.Hex())
		if id2 != id {
			t.Errorf("Sanity check fail for UUID string %s\n\tid:  %v\n\tid2: %v", id.Hex(), id, id2)
			t.FailNow()
		}
	}
}

func TestHex(t *testing.T) {
	x := [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	s := strings.ToLower(UUID(x).Hex())
	if s != "01020304-0506-0708-090a-0b0c0d0e0f10" {
		t.Errorf("Hex fail:\n\tBinary: %v,\n\tBad hex: %s", x, s)
		t.FailNow()
	}
}

func TestRand(t *testing.T) {
	uuid := Rand()
	t.Log(uuid.Hex())
}
