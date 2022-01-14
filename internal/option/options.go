package option

// Options define available flag/options
type Options struct {
	Target      string
	Concurrency int
	Silent      bool
	List        []string
	Output      string
}
