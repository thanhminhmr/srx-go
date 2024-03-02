package primary_context

type ByteHistory struct {
	firstByte  uint8
	secondByte uint8
	thirdByte  uint8
	stateIndex uint8
}

func (history ByteHistory) FirstByte() uint8 {
	return history.firstByte
}

func (history ByteHistory) SecondByte() uint8 {
	return history.secondByte
}

func (history ByteHistory) ThirdByte() uint8 {
	return history.thirdByte
}

func (history ByteHistory) State() HistoryState {
	return stateTable[history.stateIndex]
}

func (history *ByteHistory) Matching(currentState HistoryState, nextByte uint8) ByteMatched {
	if history.firstByte == nextByte {
		history.stateIndex = currentState.nextIfFirst
		return ByteMatchedFirst
	} else if history.secondByte == nextByte {
		history.secondByte = history.firstByte
		history.firstByte = nextByte
		history.stateIndex = currentState.nextIfSecond
		return ByteMatchedSecond
	} else if history.thirdByte == nextByte {
		history.thirdByte = history.secondByte
		history.secondByte = history.firstByte
		history.firstByte = nextByte
		history.stateIndex = currentState.nextIfThird
		return ByteMatchedThird
	} else {
		history.thirdByte = history.secondByte
		history.secondByte = history.firstByte
		history.firstByte = nextByte
		history.stateIndex = currentState.nextIfMiss
		return ByteMatchedNone
	}
}

func (history *ByteHistory) Matched(currentState HistoryState, nextByte uint8, matched ByteMatched) {
	switch matched {
	case ByteMatchedFirst:
		history.stateIndex = currentState.nextIfFirst
	case ByteMatchedNone:
		history.thirdByte = history.secondByte
		history.secondByte = history.firstByte
		history.firstByte = nextByte
		history.stateIndex = currentState.nextIfMiss
	case ByteMatchedSecond:
		history.secondByte = history.firstByte
		history.firstByte = nextByte
		history.stateIndex = currentState.nextIfSecond
	case ByteMatchedThird:
		history.thirdByte = history.secondByte
		history.secondByte = history.firstByte
		history.firstByte = nextByte
		history.stateIndex = currentState.nextIfThird
	}
}
