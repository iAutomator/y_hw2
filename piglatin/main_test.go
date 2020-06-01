package piglatin

import "testing"

var testData = []struct {
	input  string
	output string
}{
	{"ball", "allbay"},
	{"star", "arstay"},
	{"This is an example of Pig Latin", "Isthay isay anay exampleay ofay Igpay Atinlay"},
	{"", ""},
}

func TestEncode(t *testing.T) {
	for _, d := range testData {
		actual := Encode(d.input)
		if actual != d.output {
			t.Errorf("inp: %q\nact: %q\nexp: %q\n", d.input, actual, d.output)
		}
	}
}
