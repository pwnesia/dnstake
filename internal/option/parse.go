package option

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Parse user-supplied options
func Parse() *Options {
	opt = &Options{}

	flag.StringVar(&opt.Target, "t", "", "")
	flag.StringVar(&opt.Target, "target", "", "")

	flag.IntVar(&opt.Concurrency, "c", 25, "")
	flag.IntVar(&opt.Concurrency, "concurrent", 25, "")

	flag.BoolVar(&opt.Silent, "s", false, "")
	flag.BoolVar(&opt.Silent, "silent", false, "")

	flag.StringVar(&opt.Output, "o", "", "")
	flag.StringVar(&opt.Output, "output", "", "")

	flag.Usage = func() {
		showBanner()
		h := []string{
			"Usage:" + usage,
			"",
			"Options:" + options,
			"",
			"Examples:" + examples,
			"",
		}

		fmt.Fprint(os.Stderr, strings.Join(h, "\n"))
	}

	flag.Parse()

	if !opt.Silent {
		showBanner()
	}

	return opt
}
