package module01

import "fmt"

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//	DecToBase(14, 16) => "E"
//	DecToBase(14, 2) => "1110"
func DecToBase(dec, base int) string {

	var res string

	for dec > 0 {
		rem := dec % base
		res = fmt.Sprintf("%d%s", rem, res)
		dec = dec / base
	}

	return res
}
