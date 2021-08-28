package runner

import (
	"fmt"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/pwnesia/dnstake/internal/errors"
	"github.com/pwnesia/dnstake/internal/option"
)

func validate(opt *option.Options) error {
	var (
		err  error
		file *os.File
	)

	if isStdin() {
		opt.List, err = read(os.Stdin)
		if err != nil {
			return err
		}
	} else if opt.Target != "" {
		_, err = os.Stat(opt.Target)

		if err != nil && govalidator.IsHost(opt.Target) {
			opt.List = []string{opt.Target}
		} else {
			file, err = os.Open(opt.Target)
			if err != nil {
				return err
			}

			opt.List, err = read(file)
			if err != nil {
				return err
			}
		}
	} else {
		return fmt.Errorf("%s", errors.ErrNoTarget)
	}

	return nil
}
