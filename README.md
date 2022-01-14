# DNSTake

<img src="https://user-images.githubusercontent.com/25837540/131214165-06cb74c3-2754-48a6-a13d-bfcf592e646a.png" width="480" alt="DNSTake" title="DNSTake">

A fast tool to check missing hosted DNS zones that can lead to subdomain takeover.

---


## What is a DNS takeover?

DNS takeover vulnerabilities occur when a subdomain (subdomain.example.com) or domain has its authoritative nameserver set to a provider (e.g. AWS Route 53, Akamai, Microsoft Azure, etc.) but the hosted zone has been removed or deleted. Consequently, when making a [request for DNS records](https://www.diggui.com/#type=A&hostname=github.technology&nameserver=public&public=8.8.8.8&specify=&clientsubnet=&tcp=def&transport=def&mapped=def&nssearch=def&trace=def&recurse=def&edns=def&dnssec=def&subnet=def&cookie=def&all=def&cmd=def&question=def&answer=def&authority=def&additional=def&comments=def&stats=def&multiline=def&short=def&colorize=on) the server responds with a `SERVFAIL` error. This allows an attacker to create the missing hosted zone on the service that was being used and thus control all DNS records for that (sub)domain.¹

## Installation

### from Binary

The ez way! You can download a pre-built binary from [releases page](https://github.com/pwnesia/dnstake/releases), just unpack and run!

### from Source

<table>
	<td><b>NOTE:</b> <a href="https://golang.org/doc/install">Go 1.16+ compiler</a> should be installed & configured!</td>
</table>

Very quick & clean!

```bash
▶ go install github.com/pwnesia/dnstake/cmd/dnstake@latest
```

#### — or

Manual building executable from source code:

```bash
▶ git clone https://github.com/pwnesia/dnstake
▶ cd dnstake/cmd/dnstake
▶ go build .
▶ (sudo) mv dnstake /usr/local/bin
```

## Usage

```console
$ dnstake -h

  ·▄▄▄▄   ▐ ▄ .▄▄ ·▄▄▄▄▄ ▄▄▄· ▄ •▄ ▄▄▄ .
  ██▪ ██ •█▌▐█▐█ ▀.•██  ▐█ ▀█ █▌▄▌▪▀▄.▀·
  ▐█· ▐█▌▐█▐▐▌▄▀▀▀█▄▐█.▪▄█▀▀█ ▐▀▀▄·▐▀▀▪▄
  ██. ██ ██▐█▌▐█▄▪▐█▐█▌·▐█ ▪▐▌▐█.█▌▐█▄▄▌
  ▀▀▀▀▀• ▀▀ █▪ ▀▀▀▀ ▀▀▀  ▀  ▀ ·▀  ▀ ▀▀▀

        (c) pwnesia.org — v0.0.1

Usage:
  [stdin] | dnstake [options]
  dnstake -t HOSTNAME [options]

Options:
  -t, --target <HOST/FILE>    Define single target host/list to check
  -c, --concurrent <i>        Set the concurrency level (default: 25)
  -s, --silent                Suppress errors and/or clean output
  -o, --output <FILE>         Save vulnerable hosts to FILE
  -h, --help                  Display its help

Examples:
  dnstake -t (sub.)domain.tld
  dnstake -t hosts.txt
  dnstake -t hosts.txt -o ./dnstake.out
  cat hosts.txt | dnstake
  subfinder -silent -d domain.tld | dnstake
```

## Workflow

**DNSTake** use [RetryableDNS client library](https://github.com/projectdiscovery/retryabledns) to send DNS queries. Initial engagement using Google & Cloudflare DNS as the resolver, then check & fingerprinting the nameservers of target host — if there is one, it will resolving the target host again with its nameserver IPs as resolver, if it gets weird DNS status response (other than `NOERROR`/`NXDOMAIN`), then it's vulnerable to be taken over. More or less [like this](https://0xpatrik.com/content/images/2018/08/ns_automation-2.png) in form of a diagram.

Currently supported DNS providers, see [here](https://github.com/indianajson/can-i-take-over-dns/blob/97104102c8ce911fd978521c703f26e1c547c613/README.md#dns-providers).

## References

- [1] https://github.com/indianajson/can-i-take-over-dns#what-is-a-dns-takeover
- https://0xpatrik.com/subdomain-takeover-ns/

## License

**DNSTake** is distributed under MIT. See `LICENSE`.
