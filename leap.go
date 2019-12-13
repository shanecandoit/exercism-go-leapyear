// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	leap := false

	// on every year that is evenly divisible by 4
	if year%4 == 0 {
		leap = true

		// except every year that is evenly divisible by 100
		if year%100 == 0 {
			leap = false

			// unless the year is also evenly divisible by 400
			if year%400 == 0 {
				leap = true
			}
		}
	}

	return leap
}
