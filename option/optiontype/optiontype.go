package optiontype

type OptionType uint

const (
	Bool OptionType = iota
	Int
	String
	StringSlice
	Uint
	Uint32
	Uint64
)
