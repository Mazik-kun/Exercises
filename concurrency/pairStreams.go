package main


type NumberPair struct { Left, Right int }

func PairStreams(left, right <-chan int) ([]NumberPair, error){
	pairs:= []NumberPair{}
	for v:= range left{
		pairs = append(pairs, NumberPair{Left: v}) 
	}
	counter := len(pairs)
	counter2 := 0
	for v := range right{
		if counter2 < counter{
			pairs[counter2].Right = v
		}
		counter2++
	}
	return pairs, nil
}