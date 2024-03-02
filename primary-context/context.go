package primary_context

type ByteMatched uint

const (
	ByteMatchedFirst  ByteMatched = iota
	ByteMatchedNone   ByteMatched = iota
	ByteMatchedSecond ByteMatched = iota
	ByteMatchedThird  ByteMatched = iota
)
