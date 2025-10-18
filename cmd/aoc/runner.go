package aoc

type Runner interface {
	Run(input string) (string, error)
}

// RunnerFunc Adapter to allow the use of ordinary functions as runner
type RunnerFunc func(input string) (string, error)

func (f RunnerFunc) Run(input string) (string, error) {
	return f(input)
}
