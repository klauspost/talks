package fse

// START OMIT
// Encoder contains information for encoding an FSE stream.
type Encoder struct {
	state     uint16                      // current state
	transform []symbolTransform           // input symbol transformation
	table     []uint16                    // encoding table
	writeBits func(value uint16, n uint8) // write n bits to stream
}

// symbolTransform contains the state transform for a symbol.
type symbolTransform struct {
	deltaFindState int32  // delta to find state of new symbol.
	deltaNbBits    uint32 // delta bits to output.
}

// Encode encodes a single byte, updates the state and writes it out.
func (e *Encoder) Encode(b byte) {
	symbolTT := e.transform[b]
	nbBitsOut := uint8((uint32(e.state) + symbolTT.deltaNbBits) >> 16)
	e.writeBits(e.state, nbBitsOut)
	dstState := int32(e.state>>nbBitsOut) + symbolTT.deltaFindState
	e.state = e.table[dstState]
}
// END OMIT