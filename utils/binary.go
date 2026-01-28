package utils

// This is a bitset object, think of it like a binary array.
// Using the boolean slice underneath the hood

type BitSet []bool

// NewBitSet creates a new BitSet of the given size
func NewBitSet(size int) BitSet {
	// note this will all be set to false by default
	return make(BitSet, size)
}

// this is kind of a cheater function that just makes sense to me
func NewFullBitSet(size int) BitSet {
	bs := make(BitSet, size)
	for i := range bs {
		bs[i] = true
	}
	return bs
}

// Set sets the bit at the given index to true
func (b BitSet) Set(index int) {
	b[index] = true
}

// Get the bit at specific index
func (b BitSet) Get(index int) bool {
	return b[index]
}

// Clear sets the bit at the given index to false
func (b BitSet) Clear(index int) {
	b[index] = false
}

// Toggle toggles the bit at the given index
func (b BitSet) Toggle(index int) {
	b[index] = !b[index]
}

func (b BitSet) IsEqual(other BitSet) bool {
	if len(b) != len(other) {
		return false
	}
	for i := range b {
		if b[i] != other[i] {
			return false
		}
	}
	return true
}

func (b BitSet) String() string {
	result := ""
	for _, bit := range b {
		if bit {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result
}

func BitSetFromString(s string) BitSet {
	bs := NewBitSet(len(s))
	for i, char := range s {
		if char == '1' {
			bs.Set(i)
		}
	}
	return bs
}

func BitSetFromInt(n int, size int) BitSet {
	bs := NewBitSet(size)
	for i := 0; i < size; i++ {
		if n&(1<<i) != 0 {
			// basically, use the AND operator to check if the bit at
			// position i is set. It shifts left by i positions and ANDs it with n.
			// I learned this stuff from Cryptopals!
			bs.Set(i)
		}
	}
	return bs
}

func (b BitSet) ToInt() int {
	result := 0
	for i, bit := range b {
		if bit {
			result |= 1 << i
		}
	}
	return result
}
