package main


// Problem14
// Longest Collatz sequence
func Problem14() interface{} {

	var longestChainLength int
	var startNum int

	for i := 0; i < 1000000; i++ {
		seq := produceCollatzSequence(i)
		if len(seq) > longestChainLength {
			startNum = i
			longestChainLength = len(seq)
		}
	}

	return startNum
}

func produceCollatzSequence(start int) []int {

	var nextInSeq int = start
	var collatzSeq []int
	collatzSeq = append(collatzSeq, nextInSeq)

	for nextInSeq != 1 && nextInSeq > 0 {
		if nextInSeq%2 == 0 {
			nextInSeq = nextInSeq / 2
		} else {
			nextInSeq = 3*nextInSeq + 1
		}

		collatzSeq = append(collatzSeq, nextInSeq)
	}

	return collatzSeq
}
