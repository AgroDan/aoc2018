# Day 18: Settlers of The North Pole

I enjoyed this one quite a bit. Mostly because I kinda already knew how to tackle this one and I could smell this challenge from a mile away. This was your standard issue "Run this experiment and give me the result of 10 iterations. Now for part two, give me the result after a jillion-gajillion iterations." Basically part one is feasible to just let it run and display the end result. Part two on the other hand will not be completed in any reasonable amount of time, so either find ways to optimize the execution stack or find a shortcut. I opted for the latter, because I don't know how to optimize it any better than I have given the rules of the challenge.

The part that made me wince a little was the fact that each coordinate within the map had to operate based on its surrounding neighbors _at the exact same time_, which means I couldn't update in place. My way out of that was to make a deep copy of the map and update the copy based on the "original" copy of the map. I had to do this iteration over each coordinate included in the map and update a new copy of the map based on it. Once it's done, I throw the old one away and return the new copy of the map. Doing this many many many times seemed highly inefficient, but I honestly didn't know any way to do it better. Maybe I could just make a reference copy and then modify the original based on the reference...but that doesn't seem to be any different.

Whatever the case, I gave Go's garbage collection a real run for its money.

So that was all that was necessary for part 1. Just do what the instructions said. I came up with the copy-and-throw-away method as explained above. Part 2 was exactly what I expected, do it a billion times. Literally a billion.

Just for yuks I ran it a billion times in hopes that it would finish after maybe 10 minutes. No dice. Oh well. I knew how to tackle this one, I just needed to figure out how to look for patterns. There really wasn't any reliable number I could use to check for repetition, so I figured that the best way to look for repetition is to simply use a number that just represents the exact state of the map at that current time. Hash to the rescue!

I started exploring some built-in system commands like [Gob](https://pkg.go.dev/encoding/gob), but ultimately just wrestling with the fact that I created a fairly complex datatype that had multiple layers, I opted instead to just...compress the runemap and then SHA-256 hash it. Store the hash in a map. Done.

I was so happy with the elegance of this I have to print it here:

```go
func hashRunes(m [][]rune) [32]byte {
	var buf bytes.Buffer
	for _, row := range m {
		buf.WriteString(string(row))
	}
	return sha256.Sum256(buf.Bytes())
}
```

Stupid simple. Elegant. I'm pretty proud of that one.

Anyway, so I created a helper function that simply hashed the state of the map at every tick, then I let it rip for 1000 iterations. Sure enough it found a very distinct pattern: at tick `408`, the map would be in the exact same state at tick `436`, at `409` it would be in the exact same state at tick `437`, `410 == 438`, `411 == 439`...there is a definite loop going on indefinitely. Just like pseudorandom number generators! We had a definitive period.

So since it would cycle forever I can now extrapolate. The pattern begins at tick `408`, so I'll subtract that from the total, making it `999,999,592`. I need to find the offset after `408`...so since the diff between ticks that are exactly the same is `28` (`436 - 408`), I'll get the modulo of `999,999,592 % 28` which is `4`.

Since the pattern repeats starting at `408`, I'll be able to know the value of pattern `436`, `464`, and `492` will all be identical. Similarly, `437 == 465 == 493` etc -- I just need to extrapolate this up to 1 billion. One of the values between tick `408` and `436` will be the same value for 1 billion.

Since I did the calculations above I know that `999,999,592 % 28` is `4`, so `4 + 408 == 412`. If I can get the resource value at tick `412`, that should be the same value as 1 billion. Cue me running the code many times and getting the wrong answer until I realized that I wasn't starting at the same place that I should have started counting for ticks...some curse words, some hair pulling, some doubting my capability and yelling at the rubber duckie on my desk trying to explain that I'm no good and not deserving of my career...and eventually I got it. Thank you rubber duckie, I'm sorry for the things I said to you.

Believe it or not, there's a lot to be said about pseudorandomization here because this is essentially a pseudorandom number generator. Unfortunately not a very good one since it hits a state and just repeats indefinitely. A proper PRNG wouldn't repeat until an extremely long amount of time, not unlike [The Mersenne Twister](https://agrohacksstuff.io/posts/pseudo-random-number-generators-and-why-you-should-tread-lightly/). Even still, because of this predictable pattern, we can exploit a fundamental flaw here.

A thinly-veiled exploit disguised as an Advent of Code challenge. I think I just got goosebumps.