package fse

// DecSymbol contains information about a state entry.
type DecSymbol struct {
	newState uint16 // offset base
	symbol   byte   // output symbol
	nbBits   uint8  // bits to read
}

// Decoder contains information for decoding an FSE stream.
type Decoder struct {
	state   uint16               // current state
	table   []DecSymbol          // decoding table
	getBits func(n uint8) uint16 // get bits from stream
}

// Next returns the next symbol and sets the next state.
func (d *Decoder) Next() byte {
	n := d.table[d.state]
	lowBits := d.getBits(n.nbBits)
	d.state = n.newState + lowBits
	return n.symbol
}
