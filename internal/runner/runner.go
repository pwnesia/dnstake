package runner

import (
	"github.com/pwnesia/dnstake/internal/executor"
	"github.com/pwnesia/dnstake/internal/option"
	"github.com/remeh/sizedwaitgroup"
)

// New of program runner
func New(opt *option.Options) error {
	if err := validate(opt); err != nil {
		return err
	}

	job := make(chan string)
	con := opt.Concurrency
	swg := sizedwaitgroup.New(con)

	for i := 0; i < con; i++ {
		swg.Add()
		go func() {
			defer swg.Done()
			for hostname := range job {
				executor.New(opt, hostname)
			}
		}()
	}

	for _, hostname := range opt.List {
		job <- hostname
	}

	close(job)
	swg.Wait()

	return nil
}
