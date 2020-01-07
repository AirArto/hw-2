package stringunpack

import (
	"testing"
)

func TestDecode(t *testing.T) {
	PackedArr := []string{"a2b3c4", "abcd", "45", "qwe\\4\\5", "qwe\\45", "qwe\\\\5", "\\\\54", ""}
	UnpackedArr := []string{"aabbbcccc", "abcd", "", "qwe45", "qwe44444", "qwe\\\\\\\\\\", "", ""}
	ErrArr := []string{"", "", "Wrong packed string", "", "", "", "Wrong packed string", ""}
	for i, str := range PackedArr {
		data, err := Do(str)
		if data != UnpackedArr[i] {
			t.Errorf("Unexpected result while string unpacking: %s \nExpected:%s \nReceived:%s", str, data, UnpackedArr[i])
		}
		if err != nil {
			if err.Error() != ErrArr[i] {
				t.Errorf("Unexpected error while string unpacking: %s \nExpected:%s \nReceived:%s", str, err.Error(), ErrArr[i])
			}
		}
	}
}
