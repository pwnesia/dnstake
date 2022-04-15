package option

const (
  version = "0.1.1"
  author  = "pwnesia.org"
  banner  = `
  ·▄▄▄▄   ▐ ▄ .▄▄ ·▄▄▄▄▄ ▄▄▄· ▄ •▄ ▄▄▄ .
  ██▪ ██ •█▌▐█▐█ ▀.•██  ▐█ ▀█ █▌▄▌▪▀▄.▀·
  ▐█· ▐█▌▐█▐▐▌▄▀▀▀█▄▐█.▪▄█▀▀█ ▐▀▀▄·▐▀▀▪▄
  ██. ██ ██▐█▌▐█▄▪▐█▐█▌·▐█ ▪▐▌▐█.█▌▐█▄▄▌
  ▀▀▀▀▀• ▀▀ █▪ ▀▀▀▀ ▀▀▀  ▀  ▀ ·▀  ▀ ▀▀▀

        (c) ` + author + ` — v` + version
  usage = `
  [stdin] | dnstake [options]
  dnstake -t HOSTNAME [options]`
  options = `
  -t, --target <HOST/FILE>    Define single target host/list to check
  -c, --concurrent <i>        Set the concurrency level (default: 25)
  -s, --silent                Suppress errors and/or clean output
  -o, --output <FILE>         Save vulnerable hosts to FILE
  -h, --help                  Display its help`
  examples = `
  dnstake -t (sub.)domain.tld
  dnstake -t hosts.txt
  dnstake -t hosts.txt -o ./dnstake.out
  cat hosts.txt | dnstake
  subfinder -silent -d domain.tld | dnstake
  `
)
