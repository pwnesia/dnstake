package fingerprint

import "regexp"

// Check do fingerprinting
func Check(NS []string) (DNS, string, error) {
	for _, f := range Get() {
		for _, r := range NS {
			m, e := regexp.MatchString(f.Pattern, r)
			if e != nil {
				return DNS{}, "", e
			}

			if m {
				return f, r, nil
			}
		}
	}

	return DNS{}, "", nil
}
