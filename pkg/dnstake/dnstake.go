package dnstake

import "github.com/projectdiscovery/retryabledns"

// Verify if the response is not NOERROR
func Verify(IPs []string, hostname string) (bool, error) {
	retries := 3

	if len(IPs) < 1 {
		return false, nil
	}

	for i, k := range IPs {
		if k == "" {
			continue
		}

		IPs[i] += ":53"
	}

	if len(IPs) > retries {
		retries = len(IPs)
	}

	res, err := Resolve(retryabledns.New(IPs, retries), hostname, 1)
	if err != nil {
		return false, err
	}

	if res.StatusCode == "SERVFAIL" || res.StatusCode == "REFUSED" {
		return true, nil
	}

	return false, nil
}
