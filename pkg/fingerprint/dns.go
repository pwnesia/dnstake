package fingerprint

// DNS define DNS service providers
type DNS struct {
	Provider string
	Status   []int // 0 not vuln, 1 vulnerable, 2 edge case, 3 need purchase
	Pattern  string
}
