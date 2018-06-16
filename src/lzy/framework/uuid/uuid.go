package uuid

import (
	crand "crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	mrand "math/rand"
	"regexp"
	"strings"
	"time"
)

var seeded bool = false

var uuidRegex *regexp.Regexp = regexp.MustCompile(`^\{?([a-fA-F0-9]{8})-?([a-fA-F0-9]{4})-?([a-fA-F0-9]{4})-?([a-fA-F0-9]{4})-?([a-fA-F0-9]{12})\}?$`)

type UUID [16]byte

func (this UUID) Hex() string {
	x := [16]byte(this)
	return fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		x[0], x[1], x[2], x[3], x[4],
		x[5], x[6],
		x[7], x[8],
		x[9], x[10], x[11], x[12], x[13], x[14], x[15])

}

func Rand() UUID {
	var x [16]byte
	randBytes(x[:])
	x[6] = (x[6] & 0x0F) | 0x40
	x[8] = (x[8] & 0x3F) | 0x80
	return x
}

func FromStr(s string) (id UUID, err error) {
	if s == "" {
		err = errors.New("Empty string")
		return
	}

	parts := uuidRegex.FindStringSubmatch(s)
	if parts == nil {
		err = errors.New("Invalid string format")
		return
	}

	var array [16]byte
	slice, _ := hex.DecodeString(strings.Join(parts[1:], ""))
	copy(array[:], slice)
	id = array
	return
}

func MustFromStr(s string) UUID {
	id, err := FromStr(s)
	if err != nil {
		panic(err)
	}
	return id
}

func randBytes(x []byte) {

	length := len(x)
	n, err := crand.Read(x)

	if n != length || err != nil {
		if !seeded {
			mrand.Seed(time.Now().UnixNano())
		}

		for length > 0 {
			length--
			x[length] = byte(mrand.Int31n(256))
		}
	}
}
