package token

import (
	"errors"
	"fmt"
	"strconv"
)

// normalize converts '\u2138' to 'ℸ' and '\xab' to '«' from the input byte slice. A byte slice is returned with the
// conversions. An error is returned if the unicode literals are invalid.
func normalize(input []byte) (out []byte, err error) {
	// NOTE: The google unicode normalize library is not used because it decomposes the unicode codepoints, and we want
	// to preserve them being a document format.
	r := 0
	for r < len(input) {
		if input[r] == '\\' && input[r+1] == 'u' {
			if len(input) < r+6 {
				err = errors.New("invalid unicode literal")
				return
			}
			i, err2 := strconv.ParseUint(string(input[r+2:r+6]), 16, 64)
			if err2 != nil {
				err = fmt.Errorf("invalid unicode literal: %s", err2)
				return
			}
			out = append(out, []byte(string(rune(i)))...)
			r += 6
		} else if input[r] == '\\' && input[r+1] == 'x' {
			if len(input) < r+4 {
				err = errors.New("invalid unicode literal")
				return
			}
			i, err2 := strconv.ParseUint(string(input[r+2:r+4]), 16, 64)
			if err2 != nil {
				err = fmt.Errorf("invalid unicode literal: %s", err2)
				return
			}
			out = append(out, []byte(string(rune(i)))...)
			r += 4
		} else if input[r] == '\\' && (input[r+1] == '\\') {
			out = append(out, '\\')
			r += 2
		} else {
			out = append(out, input[r])
			r++
		}
	}
	return
}
