package stringunpack

import (
	"errors"
	"strings"
	"unicode"
)

// Do for strings array unpacking
func Do(packed string) (unpacked string, err error) {
	var (
		unpackedB   strings.Builder
		packedArray []rune = []rune(packed)
	)

	if len(packed) != 0 {
		if !unicode.IsNumber(packedArray[0]) {
			unpackedB = buildString(packedArray)
		}

		if unpackedB.Len() == 0 {
			err = errors.New("Wrong packed string")
		}
	}

	unpacked = unpackedB.String()
	return
}

func buildString(packedArray []rune) (unpackedB strings.Builder) {
	var (
		lastRune    rune
		isEscaped   bool = false
		isMultipled bool = false
	)
	for i, tempRune := range packedArray {
		isEscaped = !isEscaped && lastRune == rune('\\')
		if i < len(packedArray)-1 {
			isMultipled = unicode.IsNumber(packedArray[i+1])
		} else {
			isMultipled = false
		}

		if !isEscaped && isMultipled && unicode.IsNumber(tempRune) {
			unpackedB.Reset()
			break
		}
		if tempRune == rune('\\') {
			if isEscaped && !isMultipled {
				unpackedB.WriteRune(tempRune)
			}
		} else if unicode.IsNumber(tempRune) {
			if isEscaped {
				if !isMultipled {
					unpackedB.WriteRune(tempRune)
				}
			} else {
				for j := 0; j < int(tempRune-'0'); j++ {
					unpackedB.WriteRune(lastRune)
				}
			}
		} else {
			if isEscaped {
				unpackedB.Reset()
				break
			}
			if !isMultipled {
				unpackedB.WriteRune(tempRune)
			}
		}

		lastRune = tempRune
	}
	if lastRune == rune('\\') && !isEscaped {
		unpackedB.Reset()
	}
	return
}
