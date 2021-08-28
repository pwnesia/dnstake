package fingerprint

// Get predefined DNS service providers
// taken from https://github.com/indianajson/can-i-take-over-dns#dns-providers
func Get() []DNS {
	return []DNS{
		{
			Provider: "000Domains",
			Status:   []int{1, 3},
			Pattern:  `^(fw)?ns[\d]\.000domains\.com$`,
		},
		{
			Provider: "AWS Route 53",
			Status:   []int{0},
			Pattern:  `^ns\-([\w]{4}\.awsdns\-[\w]{2}\.(co\.uk|org)|[\w]{3}\.awsdns\-[\w]{2}\.(com|net))$`,
		},
		{
			Provider: "Microsoft Azure",
			Status:   []int{1},
			Pattern:  `^ns(4\-[\w]{2}\.azure\-dns\.info|1\-[\w]{2}\.azure\-dns\.com|2\-[\w]{2}\.azure\-dns\.net|3\-[\w]{2}\.azure\-dns\.org)$`,
		},
		{
			Provider: "Bizland",
			Status:   []int{1},
			Pattern:  `^(clickme2?\.click2site\.com|ns[12]\.bizland\.com)$`,
		},
		{
			Provider: "Cloudflare",
			Status:   []int{2},
			Pattern:  `^[\w]+\.ns\.cloudflare\.com$`,
		},
		{
			Provider: "DigitalOcean",
			Status:   []int{1},
			Pattern:  `^ns[1-3]\.digitalocean\.com$`,
		},
		{
			Provider: "DNSMadeEasy",
			Status:   []int{1},
			Pattern:  `^ns[\w]{,2}\.dnsmade{2}asy\.com$`,
		},
		{
			Provider: "DNSimple",
			Status:   []int{1},
			Pattern:  `^ns[1-4]\.dnsimple\.com$`,
		},
		{
			Provider: "Domain.com",
			Status:   []int{1, 3},
			Pattern:  `^ns[1-2]\.domain\.com$`,
		},
		{
			Provider: "DomainPeople",
			Status:   []int{0},
			Pattern:  `^ns[1-2]\.domainpeople\.com$`,
		},
		{
			Provider: "Dotster",
			Status:   []int{1, 3},
			Pattern:  `^ns[12]\.(nameresolve|dotster)\.com$`,
		},
		{
			Provider: "Dotster",
			Status:   []int{1, 3},
			Pattern:  `^ns[12]\.(nameresolve|dotster)\.com$`,
		},
		{
			Provider: "EasyDNS",
			Status:   []int{1},
			Pattern:  `^dns(?:4\.easydns\.info|1\.easydns\.com|2\.easydns\.net|3\.easydns\.org)$`,
		},
		{
			Provider: "Gandi.net",
			Status:   []int{0},
			Pattern:  `^[\w]+\.dns\.gandi\.net$`,
		},
		{
			Provider: "Google Cloud",
			Status:   []int{1},
			Pattern:  `^ns\-cloud\-[\w]+\.go{2}gledomains\.com$`,
		},
		{
			Provider: "Hover",
			Status:   []int{0},
			Pattern:  `^ns[1-2]\.hover\.com$`,
		},
		{
			Provider: "Hurricane Electric",
			Status:   []int{1},
			Pattern:  `^ns[1-5]\.he\.net$`,
		},
		{
			Provider: "Linode",
			Status:   []int{1},
			Pattern:  `^ns[1-2]\.linode\.com$`,
		},
		{
			Provider: "MediaTemple",
			Status:   []int{0},
			Pattern:  `^ns[1-2]\.mediatemple\.net$`,
		},
		{
			Provider: "MyDomain",
			Status:   []int{1, 3},
			Pattern:  `^ns[1-2]\.mydomain\.com$`,
		},
		{
			Provider: "Name.com",
			Status:   []int{1, 3},
			Pattern:  `^ns[1-4][\w]+?\.name\.com$`,
		},
		{
			Provider: "Network Solutions",
			Status:   []int{0},
			Pattern:  `^ns[\w]+?\.worldnic\.com$`,
		},
		{
			Provider: "NS1",
			Status:   []int{1},
			Pattern:  `^dns[1-4]\.p[\d]{,2}\.nsone\.net$`,
		},
		{
			Provider: "TierraNet",
			Status:   []int{1},
			Pattern:  `^ns[1-2]\.domaindiscover\.com$`,
		},
		{
			Provider: "Reg.ru",
			Status:   []int{1, 3},
			Pattern:  `^ns[1-2]\.reg\.ru$`,
		},
		{
			Provider: "UltraDNS",
			Status:   []int{0},
			Pattern:  `^[psu]dns[\d]{,3}\.ultradns\.com$`,
		},
		{
			Provider: "Yahoo Small Business",
			Status:   []int{1, 3},
			Pattern:  `^yns[1-2]\.yahoo\.com$`,
		},
	}
}
