package customencode

import "testing"

var testData = []struct {
	decoded string
	encoded string
}{
	{"hello", "h2ll4"},
	{"hi there", "h3 th2r2"},
	{"", ""},
}

func TestEncode(t *testing.T) {
	for _, d := range testData {
		actualEncoded := Encode(d.decoded)
		if actualEncoded != d.encoded {
			t.Error(d.decoded, actualEncoded, d.encoded)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, d := range testData {
		actualDecoded := Decode(d.decoded)
		if actualDecoded != d.decoded {
			t.Error(d.decoded, actualDecoded, d.encoded)
		}
	}
}
