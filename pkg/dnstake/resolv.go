package dnstake

import "github.com/projectdiscovery/retryabledns"

// Resolve target hostname by sending a query
func Resolve(client *retryabledns.Client, hostname string, qtype uint16) (*retryabledns.DNSData, error) {
	return client.Query(hostname, qtype)
}
