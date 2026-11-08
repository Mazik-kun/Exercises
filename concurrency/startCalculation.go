package main

type Outcome struct { Value int; Err error }

func StartCalculation(calculate func() (int, error)) <-chan Outcome {
	out := make(chan Outcome, 1)
	go func() {
		val, err := calculate()
		out <- Outcome{Value: val, Err: err}
		close(out)
	}()
	return out
}