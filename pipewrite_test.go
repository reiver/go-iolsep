package iolsep

import (
	"testing"

	"io"
	"strings"
)

func TestPipeWrite(t *testing.T) {

	tests := []struct{
		Data string
		Expected string
	}{
		{
			Data: "",
			Expected: "",
		},



		{
			Data:     "Hello world!\n\nOnce Twice Thrice Fource",
			Expected: "Hello world!\n",
		},
		{
			Data:     "Hello world!\n\r\n\rOnce Twice Thrice Fource",
			Expected: "Hello world!\n",
		},
		{
			Data:     "Hello world!\n\r\u0085Once Twice Thrice Fource",
			Expected: "Hello world!\n",
		},
		{
			Data:     "Hello world!\r\rOnce Twice Thrice Fource",
			Expected: "Hello world!\r\rOnce Twice Thrice Fource",
		},
		{
			Data:     "Hello world!\r\n\r\nOnce Twice Thrice Fource",
			Expected: "Hello world!\r\n",
		},
		{
			Data:     "Hello world!\r\n\u0085Once Twice Thrice Fource",
			Expected: "Hello world!\r\n",
		},
		{
			Data:     "Hello world!\u0085\u0085Once Twice Thrice Fource",
			Expected: "Hello world!\u0085",
		},
		{
			Data:     "Hello world!\u0085\n\rOnce Twice Thrice Fource",
			Expected: "Hello world!\u0085",
		},
		{
			Data:     "Hello world!\u0085\r\nOnce Twice Thrice Fource",
			Expected: "Hello world!\u0085",
		},
		{
			Data:     "Hello world!\u2028Once Twice Thrice Fource",
			Expected: "Hello world!\u2028",
		},
		{
			Data:     "Hello world!\u2029Once Twice Thrice Fource",
			Expected: "Hello world!\u2029",
		},



		{
			Data:     "Hello world!\n\n",
			Expected: "Hello world!\n",
		},
		{
			Data:     "Hello world!\n\r\n\r",
			Expected: "Hello world!\n",
		},
		{
			Data:     "Hello world!\r\r",
			Expected: "Hello world!\r\r",
		},
		{
			Data:     "Hello world!\r\n\r\n",
			Expected: "Hello world!\r\n",
		},
		{
			Data:     "Hello world!\u0085\u0085",
			Expected: "Hello world!\u0085",
		},
		{
			Data:     "Hello world!\u2028",
			Expected: "Hello world!\u2028",
		},
		{
			Data:     "Hello world!\u2029",
			Expected: "Hello world!\u2029",
		},



		{
			Data:     "\n\n",
			Expected: "\n",
		},
		{
			Data:     "\n\r\n\r",
			Expected: "\n",
		},
		{
			Data:     "\r\r",
			Expected: "\r\r",
		},
		{
			Data:     "\r\n\r\n",
			Expected: "\r\n",
		},
		{
			Data:     "\u0085\u0085",
			Expected: "\u0085",
		},
		{
			Data:     "\u2028",
			Expected: "\u2028",
		},
		{
			Data:     "\u2029",
			Expected: "\u2029",
		},



		{
			Data:     "\n\nOnce Twice Thrice Fource",
			Expected: "\n",
		},
		{
			Data:     "\n\r\n\rOnce Twice Thrice Fource",
			Expected: "\n",
		},
		{
			Data:     "\r\rOnce Twice Thrice Fource",
			Expected: "\r\rOnce Twice Thrice Fource",
		},
		{
			Data:     "\r\n\r\nOnce Twice Thrice Fource",
			Expected: "\r\n",
		},
		{
			Data:     "\u0085\u0085Once Twice Thrice Fource",
			Expected: "\u0085",
		},
		{
			Data:     "\u2028Once Twice Thrice Fource",
			Expected: "\u2028",
		},
		{
			Data:     "\u2029Once Twice Thrice Fource",
			Expected: "\u2029",
		},
	}

	for testNumber, test := range tests {
		var buffer strings.Builder
		writerune := func(r rune)(exit bool) {
			buffer.WriteRune(r)
			return false
		}

		var eof bool
		returneof := func() {
			eof = true
		}

		returnerror := func(err error) {
			t.Errorf("For test #%d, did not expect a 'return error' but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("DATA:   %q", test.Data)
			t.Logf("BUFFER: %q", buffer.String())
		}

		var reader io.Reader = strings.NewReader(test.Data)

		pipewrite(writerune, returneof, returnerror, reader)

		{
			if !eof {
				t.Errorf("For test #%d, expected 'eof' but did not actually get one." , testNumber)
				t.Logf("DATA:   %q", test.Data)
				t.Logf("BUFFER: %q", buffer.String())
			}
		}

		{
			expected := test.Expected
			actual   := buffer.String()

			if expected != actual {
				t.Errorf("For test #%d, the actual 'written-runes' is not what was expected" , testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("DATA:     %q", test.Data)
			}
		}
	}
}
