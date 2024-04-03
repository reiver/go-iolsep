package iolsep_test

import (
	"testing"

	"io"
	"reflect"
	"strings"

	"github.com/reiver/go-iolsep"
)

func TestNewLineReadCloser(t *testing.T) {

	tests := []struct{
		Data string
		Expected []string
	}{
		{
			Data: "",
			Expected: []string(nil),
		},



		{
			Data:
				"apple banana cherry",
			Expected: []string{
				"apple banana cherry",
			},
		},
		{
			Data:
				"apple banana cherry"+"\n",
			Expected: []string{
				"apple banana cherry"+"\n",
			},
		},
		{
			Data:
				"apple banana cherry"+"\n\n",
			Expected: []string{
				"apple banana cherry"+"\n",
				""+"\n",
			},
		},



		{
			Data:
				"apple banana cherry"+"\n"+
				"once twice thrice fource",
			Expected: []string{
				"apple banana cherry"+"\n",
				"once twice thrice fource",
			},
		},
		{
			Data:
				"apple banana cherry"+"\n\n"+
				"once twice thrice fource",
			Expected: []string{
				"apple banana cherry"+"\n",
				""+"\n",
				"once twice thrice fource",
			},
		},



		{
			Data:
				"apple banana cherry"+"\r"+
				"once twice thrice fource",
			Expected: []string{
				"apple banana cherry"+"\r"+"once twice thrice fource",
			},
		},
		{
			Data:
				"apple banana cherry"+"\r\r"+
				"once twice thrice fource",
			Expected: []string{
				"apple banana cherry"+"\r\r"+"once twice thrice fource",
			},
		},



		{
			Data:
				"apple banana cherry"+"\n\r\n\r"+
				"once twice thrice fource",
			Expected: []string{
				"apple banana cherry"+"\n",
				""+"\r\n",
				""+"\r"+"once twice thrice fource",
			},
		},
		{
			Data:
				"apple banana cherry"+"\r\n\r\n"+
				"once twice thrice fource",
			Expected: []string{
				"apple banana cherry"+"\r\n",
				""+"\r\n",
				"once twice thrice fource",
			},
		},
		{
			Data:
				"apple banana cherry"+"\n\r\r\n"+
				"once twice thrice fource",
			Expected: []string{
				"apple banana cherry"+"\n",
				""+"\r\r\n",
				"once twice thrice fource",
			},
		},
		{
			Data:
				"apple banana cherry"+"\r\n\n\r"+
				"once twice thrice fource",
			Expected: []string{
				"apple banana cherry"+"\r\n",
				""+"\n",
				""+"\r"+"once twice thrice fource",
			},
		},



		{
			Data:
				"apple banana cherry"+"\u0085\u0085"+
				"once twice thrice fource",
			Expected: []string{
				"apple banana cherry"+"\u0085",
				""+"\u0085",
				"once twice thrice fource",
			},
		},
		{
			Data:
				"apple banana cherry"+"\u0085\n\r"+
				"once twice thrice fource",
			Expected: []string{
				"apple banana cherry"+"\u0085",
				""+"\n",
				""+"\r"+"once twice thrice fource",
			},
		},
		{
			Data:
				"apple banana cherry"+"\u0085\r\n"+
				"once twice thrice fource",
			Expected: []string{
				"apple banana cherry"+"\u0085",
				""+"\r\n",
				"once twice thrice fource",
			},
		},
		{
			Data:
				"apple banana cherry"+"\n\r\u0085"+
				"once twice thrice fource",
			Expected: []string{
				"apple banana cherry"+"\n",
				""+"\r\u0085",
				"once twice thrice fource",
			},
		},
		{
			Data:
				"apple banana cherry"+"\r\n\u0085"+
				"once twice thrice fource",
			Expected: []string{
				"apple banana cherry"+"\r\n",
				""+"\u0085",
				"once twice thrice fource",
			},
		},



		{
			Data:
				"apple banana cherry"+"\u2029"+
				"once twice thrice fource",
			Expected: []string{
				"apple banana cherry"+"\u2029",
				"once twice thrice fource",
			},
		},
	}

	for testNumber, test := range tests {

		var actual []string

		var reader io.Reader = strings.NewReader(test.Data)

		for {
			var actualBytes []byte

			func(){
				var readcloser io.ReadCloser = iolsep.NewLineReadCloser(reader)
				defer func() {
					err := readcloser.Close()
					if nil != err {
						t.Errorf("For test #%d, did not expected an error when 'closing' but actually got one.", testNumber)
						t.Logf("ERROR: (%T) %s", err, err)
						t.Logf("DATA: %q", test.Data)
					}
				}()

				var err error

				actualBytes, err = io.ReadAll(readcloser)
				if nil != err {
					t.Errorf("For test #%d, did not expected an error when 'reading-all' but actually got one.", testNumber)
					t.Logf("ERROR: (%T) %s", err, err)
					t.Logf("DATA: %q", test.Data)
					return
				}
			}()

			if len(actualBytes) <= 0 {
				break
			}

			actual = append(actual, string(actualBytes))
		}

		{
			expected := test.Expected

			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d, the actual result is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("DATA: %q", test.Data)
				continue
			}
		}
	}
}
