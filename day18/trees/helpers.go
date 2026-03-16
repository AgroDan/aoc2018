package trees

import (
	"bytes"
	"crypto/sha256"
)

// I may add these to the aocutils package because I'm going to try
// and use a go standard package to perform hashes of data so that
// I can reasonably track differences in the runemap in the hopes
// that I'll find a cycle. We'll see I guess!

// func gobHash(v any) ([32]byte, error) {
// 	// EXTRA NOTE: DO NOT USE THIS FOR MAPS THAT AREN'T CONSISTENTLY ORDERED!
// 	// Because gob will sort this by iteration, and go randomizes iteration
// 	// by default, there really is no way to compare maps because the order
// 	// will be different and return a separate hash regardless. Just keep
// 	// that in mind I guess.
// 	var buf bytes.Buffer
// 	if err := gob.NewEncoder(&buf).Encode(v); err != nil {
// 		return [32]byte{}, err
// 	}
// 	return sha256.Sum256(buf.Bytes()), nil
// }

// func (a Acres) GobEncode() ([]byte, error) {
// 	var buf bytes.Buffer
// 	proxy := Acres{
// 		Runemap: a.Runemap,
// 	}
// 	if err := gob.NewEncoder(&buf).Encode(proxy); err != nil {
// 		return nil, err
// 	}
// 	return buf.Bytes(), nil
// }

// func (a Acres) GobEncode() ([]byte, error) {
// 	// implementing gobencode to capture the full Acres state
// 	var buf bytes.Buffer
// 	if err := gob.NewEncoder(&buf).Encode(a.Runemap); err != nil {
// 		return nil, err
// 	}
// 	return buf.Bytes(), nil
// }

// func (a *Acres) GobDecode(data []byte) error {
// 	// mirror of GobEncode to properly decode
// 	buf := bytes.NewBuffer(data)
// 	return gob.NewDecoder(buf).Decode(&a.Runemap)
// }

// This is probably a better way of doing it without relying on gob i guess
func hashRunes(m [][]rune) [32]byte {
	var buf bytes.Buffer
	for _, row := range m {
		buf.WriteString(string(row))
	}
	return sha256.Sum256(buf.Bytes())
}

func CheckForPattern(a Acres, iter int) int {
	// This is just a helper so it will look for patterns
	// and just print them the second we see a repeat.
	seen := make(map[[32]byte]int)

	var startingRep, diff int

	for i := 0; i < iter; i++ {
		hash := hashRunes(a.GetRaw())
		if prev, ok := seen[hash]; ok {
			// fmt.Printf("Pattern found at %d and %d\n", prev, i)
			startingRep = prev
			diff = i - prev
			break

		}
		seen[hash] = i
		a = a.Tick()
		// a.Print()
	}

	// fmt.Printf("Looking for offset...\n")
	// 1_000_000_000
	return ((1_000_000_000 - startingRep) % diff) + startingRep

	// so the way I see it, if we have 5 repeats in a row that's good enough
	// for me to say that there is a repetition pattern going on. Instead of confirming
	// that we have 5 repeats, I know for a fact that it will repeat every "diff" iterations,
	// so I guess just trust me bro

	// Anyways, for my input field I noticed that the first "collision" I got was at 408 and
	// 436, which means the diff is 28. This number is a constant for each item after it
	// ie 409 == 437, 410 == 438, 411 == 439, etc -- so now all I have to do is determine
	// the offset. Since I know that the repeated pattern starts at 408, I can subtract
	// 408 from 1_000_000_000 and get 999_999_592. This is the new number to loop for.

	// Now modulo 999_999_592 from the discovered diff of 28. The answer is 4. This means
	// that since the pattern repeats, the value after 412 ticks (408 + 4) will equal the
	// value after 1_000_000_000 ticks.
}
