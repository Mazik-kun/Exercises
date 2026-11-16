package main

func ApplyUntilError(in <-chan int, apply func(int) error) error {
	var err error
	for v := range in {
		if err == nil {
			err = apply(v)
		}
	}
	return err
}